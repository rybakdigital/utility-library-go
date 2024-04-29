package limiter

import (
	"time"
)

const (
	RATE_LIMITER_LIMIT     = 120
	RATE_LIMITER_TIMEFRAME = time.Minute
)

type Config struct {
	Limit     int
	Interval  time.Duration
	StoreType string
}

func DefaultConfig() Config {
	return NewConfig(RATE_LIMITER_LIMIT, time.Minute, RATE_LIMITER_STORE_MEMORY)
}

func NewConfig(limit int, interval time.Duration, storeType string) Config {
	return Config{Limit: limit, Interval: interval, StoreType: storeType}
}
