package adapter

import (
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_MYSQL_MAX_ATTEMPTS     = 5               // Default number of attempts to connect to mysql
	DEFAULT_MYSQL_ATTEMPT_INTERVAL = 5 * time.Second // Time between attempts
)

type Config struct {
	dsn         string
	maxAttempts int
	interval    time.Duration
}

func NewConfig(dsn string, maxAttempts int, interval time.Duration) *Config {
	return &Config{dsn: dsn, maxAttempts: maxAttempts, interval: interval}
}

func DefaultConfig() *Config {
	// Check if MySQL attempts have been defined in env
	max := os.Getenv("MYSQL_MAX_ATTEMPTS")
	maxAttempts := DEFAULT_MYSQL_MAX_ATTEMPTS

	if max != "" {
		if logger != nil {
			logger.Printf("Read max attempts from MYSQL_MAX_ATTEMPTS env: %s", max)
		}
		maxAttempts, _ = strconv.Atoi(max)
	}

	if logger != nil {
		logger.Printf("Adapter configured to attempt to connect %d times to mysql", maxAttempts)
	}

	// Check if MySQL attempts interval have been defined in env
	i := os.Getenv("MYSQL_ATTEMPT_INTERVAL")
	interval := DEFAULT_MYSQL_ATTEMPT_INTERVAL

	if i != "" {
		if logger != nil {
			logger.Printf("Read attempts interval from MYSQL_ATTEMPT_INTERVAL env: %s", i)
		}

		intSeconds, _ := strconv.Atoi(i)
		interval = time.Duration(float64(intSeconds) * float64(time.Second))
	}

	if logger != nil {
		logger.Printf("Adapter configured to attempt connection every %d to mysql", interval)
	}

	return NewConfig(
		getDefaultDsn(),
		maxAttempts,
		interval)
}

func (c *Config) GetDsn() string {
	return c.dsn
}

func (c *Config) GetMaxAttempts() int {
	return c.maxAttempts
}

func (c *Config) GetInterval() time.Duration {
	return c.interval
}

func getDefaultDsn() string {
	// Get MySQL envs
	user := os.Getenv("MYSQL_USER")
	if user != "" {
		if logger != nil {
			logger.Printf("Read MySQL user from MYSQL_USER env: %s", user)
		}
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password != "" {
		if logger != nil {
			logger.Printf("Read MySQL password from MYSQL_PASSWORD env")
		}
	}

	host := os.Getenv("MYSQL_HOST")
	if host != "" {
		if logger != nil {
			logger.Printf("Read MySQL host from MYSQL_HOST env %s", host)
		}
	}

	port := os.Getenv("MYSQL_PORT")
	if port != "" {
		if logger != nil {
			logger.Printf("Read MySQL port from MYSQL_PORT env %s", port)
		}
	}

	database := os.Getenv("MYSQL_DATABASE")
	if database != "" {
		if logger != nil {
			logger.Printf("Read MySQL database from MYSQL_DATABASE env %s", database)
		}
	}

	return user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
}
