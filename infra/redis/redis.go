package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var RedisClient *redis.Client

type RedisConfig struct {
	Host         string
	Port         string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
	MaxRetries   int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewRedisClientWithPool(config RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,     // Maximum number of socket connections
		MinIdleConns: config.MinIdleConns, // Minimum number of idle connections
		MaxRetries:   config.MaxRetries,   // Maximum number of retries before giving up
		DialTimeout:  config.DialTimeout,  // Dial timeout for establishing new connections
		ReadTimeout:  config.ReadTimeout,  // Timeout for socket reads
		WriteTimeout: config.WriteTimeout, // Timeout for socket writes
		
		// Connection pool settings
		PoolTimeout:  30 * time.Second, // Amount of time client waits for connection		
		// Retry settings
		MaxRetryBackoff: 512 * time.Millisecond,
		MinRetryBackoff: 8 * time.Millisecond,
	})

	ctx := context.Background()
	
	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to Redis")
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	// Print pool stats
	stats := rdb.PoolStats()
	log.Printf("Redis Pool Stats - Hits: %d, Misses: %d, Timeouts: %d, TotalConns: %d, IdleConns: %d\n",
		stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns)

	return rdb, nil
}