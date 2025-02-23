package cache

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const (
	// Compression types
	CompressNone = 0 // No compression
	CompressGzip = 1 // Gzip compression
	CompressLz4  = 2 // LZ4 compression
	CompressZstd = 3 // Zstd compression

	// Cache types
	CacheDisabled  = 0 // Cache disabled
	CacheEmbedded  = 1 // Use embedded Redis server
	CacheExternal  = 2 // Use external Redis server
	CacheAutomatic = 3 // Automatic mode, prefer external server
)

var (
	// Rc is the Redis client instance
	Rc *redis.Client
	// Rs is the embedded Redis server instance
	Rs *miniredis.Miniredis
	// Ra is the Redis server address, defaults to localhost:6379
	Ra = "localhost:6379"
	// Mu is a mutex for thread-safe operations
	Mu              sync.Mutex
	Type            = CacheEmbedded // Cache type: 0=disabled, 1=embedded, 2=external, 3=automatic
	Debug           = 0             // Debug switch: 1=enabled, 0=disabled
	CompressType    = CompressNone  // Compression type: no compression by default
	CompressMinSize = 1024          // Minimum compression size: 1KB

	ErrDisabled = errors.New("cache is disabled")
)

func init() {
	// if err := InitCache(); err != nil {
	// 	logrus.Errorf("Failed to initialize cache: %v", err)
	// }
}

// InitCache initializes the Redis cache based on the configured Type.
// It supports embedded, external, and automatic modes.
// Returns an error if initialization fails or times out after 10 seconds.
func InitCache() error {

	if Type == CacheDisabled {
		logrus.Info("Cache is disabled")
		return nil
	}

	done := make(chan bool)
	errChan := make(chan error)

	go func() {
		switch Type {
		case CacheEmbedded:
			logrus.Info("Starting embedded Redis server")
			Rs = start()
			if err := connect(Rs.Addr()); err != nil {
				errChan <- fmt.Errorf("failed to connect to embedded Redis: %v", err)
				return
			}
		case CacheExternal:
			logrus.Info("Connecting to external Redis server")
			if err := connect(Ra); err != nil {
				errChan <- fmt.Errorf("failed to connect to external Redis: %v", err)
				return
			}
		case CacheAutomatic:
			logrus.Info("Connecting to  Redis server ")
			if err := connect(Ra); err != nil {
				logrus.Info("Failed to connect to external Redis, falling back to embedded server")
				Rs = start()
				if err := connect(Rs.Addr()); err != nil {
					errChan <- fmt.Errorf("failed to connect to embedded Redis: %v", err)
					return
				}
			}
		}
		errChan <- nil
		done <- true
	}()

	select {
	case err := <-errChan:
		return err
	case <-done:
		return nil
	case <-time.After(10 * time.Second):
		return fmt.Errorf("timeout while initializing cache")
	}
}

// start starts an embedded Redis server
// Returns a new miniredis instance or panics if startup fails
func start() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		logrus.Fatalf("Failed to start embedded Redis server: %v", err)
	}
	return s
}

// connect establishes a connection to the Redis server at the specified address
// Returns an error if the connection or ping fails
func connect(addr string) error {
	Rc = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Rc.Ping(ctx).Result()
	if err != nil {
		return err
	}

	logrus.Println("Connected to Redis at", addr)
	return nil
}

// Shutdown gracefully closes the embedded Redis server if it exists
func Shutdown() {
	if Rs != nil {
		Rs.Close()
	}
}

// compress compresses the input data using the configured compression method
// Returns the compressed data and any error that occurred during compression
// If the data size is smaller than CompressMinSize, returns the original data
func compress(data []byte) ([]byte, error) {
	// Skip compression if data is smaller than threshold
	if len(data) < CompressMinSize {
		return data, nil
	}

	switch CompressType {
	case CompressNone:
		return data, nil

	case CompressGzip:
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		if _, err := zw.Write(data); err != nil {
			return nil, err
		}
		if err := zw.Close(); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil

	case CompressLz4:
		buf := make([]byte, lz4.CompressBlockBound(len(data)))
		n, err := lz4.CompressBlock(data, buf, nil)
		if err != nil {
			return nil, err
		}
		return buf[:n], nil

	case CompressZstd:
		encoder, err := zstd.NewWriter(nil)
		if err != nil {
			return nil, err
		}
		return encoder.EncodeAll(data, make([]byte, 0, len(data))), nil

	default:
		return nil, fmt.Errorf("unsupported compression type: %v", CompressType)
	}
}

// decompress decompresses the input data using the configured compression method
// Returns the decompressed data and any error that occurred during decompression
// If the data size is smaller than CompressMinSize, returns the original data
func decompress(data []byte) ([]byte, error) {
	if len(data) < CompressMinSize {
		return data, nil
	}

	switch CompressType {
	case CompressNone:
		return data, nil

	case CompressGzip:
		zr, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		defer zr.Close()
		return ioutil.ReadAll(zr)

	case CompressLz4:
		buf := make([]byte, len(data)*3)
		n, err := lz4.UncompressBlock(data, buf)
		if err != nil {
			return nil, err
		}
		return buf[:n], nil

	case CompressZstd:
		decoder, err := zstd.NewReader(nil)
		if err != nil {
			return nil, err
		}
		return decoder.DecodeAll(data, nil)

	default:
		return nil, fmt.Errorf("unsupported compression type: %v", CompressType)
	}
}

// Set stores a value in the cache with optional expiration time in minutes
// The value is JSON marshaled and compressed before storage
// Returns an error if the cache is disabled or if marshaling/compression fails
func Set(key string, value interface{}, expiration ...int) error {
	if Type == CacheDisabled {
		return ErrDisabled
	}

	data, err := json.Marshal(value)
	if err != nil {
		logrus.Errorf("Error Cache marshalling value: %v", err)
		return err
	}

	compressed, err := compress(data)
	if err != nil {
		logrus.Errorf("Error compressing data: %v", err)
		return err
	}

	if Debug == 1 && len(data) >= CompressMinSize {
		compressionRatio := float64(len(compressed)) / float64(len(data)) * 100
		logrus.Infof("Cache: Compression for key %s: %.2f%% (%.2fKB -> %.2fKB) using type %v",
			key, compressionRatio, float64(len(data))/1024, float64(len(compressed))/1024,
			CompressType)
	}

	// Default expiration is 10 minutes
	exp := 10 * time.Minute
	if len(expiration) > 0 {
		exp = time.Duration(expiration[0]) * time.Minute
	}

	return Rc.Set(context.Background(), key, compressed, exp).Err()
}

// Get retrieves a value from the cache and unmarshals it into the target interface
// The value is decompressed and JSON unmarshaled after retrieval
// Returns an error if the cache is disabled or if the key doesn't exist
func Get(key string, target interface{}) error {
	if Type == CacheDisabled {
		return ErrDisabled
	}

	compressed, err := Rc.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}

	data, err := decompress(compressed)
	if err != nil {
		logrus.Errorf("Error decompressing data: %v", err)
		return err
	}

	if Debug == 1 && len(compressed) >= CompressMinSize {
		compressionRatio := float64(len(compressed)) / float64(len(data)) * 100
		logrus.Infof("Cache: Decompression for key %s: %.2f%% (%.2fKB -> %.2fKB) using type %v",
			key, compressionRatio, float64(len(compressed))/1024, float64(len(data))/1024,
			CompressType)
	}

	return json.Unmarshal(data, target)
}

// Del removes a value from the cache by its key
// Returns an error if the cache is disabled or if deletion fails
func Del(key string) error {
	if Type == CacheDisabled {
		return ErrDisabled
	}
	return Rc.Del(context.Background(), key).Err()
}
