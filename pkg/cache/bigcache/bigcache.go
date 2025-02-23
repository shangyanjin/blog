package bigcache

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/allegro/bigcache"
	"github.com/sirupsen/logrus"
)

// Cache is a bigcache instance
var Cache *bigcache.BigCache
var Debug int = 1
var Enable int = 1

var ErrDisabled = errors.New("cache is disabled")

func init() {
	if Enable != 1 {
		logrus.Info("Cache is disabled")
		return
	}

	var err error
	// Max Memory Usage = 128MB, each entry up to 32KB
	cacheConfig := bigcache.Config{
		Shards:             1024,             // Increased shards for better performance
		LifeWindow:         10 * time.Minute, // Keep the same, adjust if needed
		MaxEntriesInWindow: 1500,             // Adjusted for mix of large and small entries
		MaxEntrySize:       32 * 1024,        // 32KB per entry
		HardMaxCacheSize:   128,              // 128MB hard limit
	}

	Cache, err = bigcache.NewBigCache(cacheConfig)
	if err != nil {
		logrus.Errorf("Error initializing cache: %v", err)
	} else {
		logrus.Info("Cache initialized successfully")
	}
}

// Set saves a compressed value to cache. The value must be JSON serializable.
func Set(key string, value interface{}) error {
	if Enable != 1 {
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

	if Debug == 1 {
		compressionRatio := float64(len(compressed)) / float64(len(data)) * 100
		logrus.Infof("Cache : Compression for key %s: %.2f%% (%.2fKB -> %.2fKB)",
			key, compressionRatio, float64(len(data))/1024, float64(len(compressed))/1024)
	}

	return Cache.Set(key, compressed)
}

// Get retrieves a compressed value from cache and decompresses it. The target variable must be a pointer to the type of the cached value.
func Get(key string, target interface{}) error {
	if Enable != 1 {
		return ErrDisabled
	}

	compressed, err := Cache.Get(key)
	if err != nil {
		return err
	}
	data, err := decompress(compressed)
	if err != nil {
		logrus.Errorf("Error decompressing data: %v", err)
		return err
	}
	if Debug == 1 {
		compressionRatio := float64(len(compressed)) / float64(len(data)) * 100
		logrus.Infof("Cache : Compression for key %s: %.2f%% (%.2fKB -> %.2fKB)",
			key, compressionRatio, float64(len(data))/1024, float64(len(compressed))/1024)
	}
	return json.Unmarshal(data, target)
}

// Delete removes a value from cache.
func Del(key string) error {
	if Enable != 1 {
		return ErrDisabled
	}
	return Cache.Delete(key)
}

// compress compresses the input data using gzip
func compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write(data)
	if err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// decompress decompresses the input data using gzip
func decompress(data []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	return ioutil.ReadAll(zr)
}
