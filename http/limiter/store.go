package limiter

import (
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/nxs23/utility-library-go/cache"
)

const (
	RATE_LIMITER_STORE_MEMORY = "memory"
	RATE_LIMITER_STORE_REDIS  = "redis"
)

func NewStore(config Config) ratelimit.Store {
	var store ratelimit.Store

	// Check store type
	if config.StoreType == RATE_LIMITER_STORE_REDIS {
		adapter := cache.Default()
		// Ping cache to confirm that is working
		adapter.Ping()

		// Redis store, this works with shared pool
		store = NewRedisStore(config, adapter)
	} else {
		// In memory store, this works only with single instance
		store = NewMemoryStore(config)
	}

	return store
}

func NewMemoryStore(config Config) ratelimit.Store {
	// In memory store, this works only with single instance
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  config.Interval,
		Limit: uint(config.Limit),
	})

	return store
}

func NewRedisStore(config Config, adapter *cache.RedisAdapter) ratelimit.Store {
	// Redis store, this works with shared pool
	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: adapter.GetClient(),
		Rate:        config.Interval,
		Limit:       uint(config.Limit),
	})

	return store
}
