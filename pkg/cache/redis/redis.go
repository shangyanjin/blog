package redis

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

var (
	rdb       *redis.Client
	rdbServer *miniredis.Miniredis
	mu        sync.Mutex
)

// InitRedis initializes the Redis client, preferring an external Redis server,
// and falls back to an embedded Redis server if the external one is unavailable.
func InitRedis(externalAddr string) {
	done := make(chan bool)

	go func() {
		if err := connect(externalAddr); err != nil {
			log.Println("Failed to connect to external Redis, starting embedded Redis:", err)
			// Start embedded Redis server
			rdbServer = startRedisServer()
			// Initialize Redis client with the embedded Redis server address
			if err := connect(rdbServer.Addr()); err != nil {
				log.Fatalf("Failed to connect to embedded Redis: %v", err)
			}
		}
		done <- true
	}()

	// Wait for initialization to complete
	<-done

	// Ensure the embedded Redis server is shut down when the program exits
	defer Shutdown()
}

// startRedisServer starts an embedded Redis server
func startRedisServer() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		log.Fatalf("Failed to start embedded Redis server: %v", err)
	}
	return s
}

// connect connects to the Redis server at the specified address
func connect(addr string) error {
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	log.Println("Connected to Redis at", addr)
	return nil
}

// GetClient returns the Redis client instance, initializing it if necessary
func GetClient() *redis.Client {
	mu.Lock()         // Lock the mutex to ensure exclusive access to the initialization logic
	defer mu.Unlock() // Defer the unlocking to ensure it happens after the function returns

	if rdb == nil {
		InitRedis("localhost:6379")
	}
	return rdb
}

// Shutdown cleans up the embedded Redis server
func Shutdown() {
	if rdbServer != nil {
		rdbServer.Close()
	}
}
