package adapter

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_MYSQL_MAX_ATTEMPTS     = 5               // Default number of attempts to connect to mysql
	DEFAULT_MYSQL_ATTEMPT_INTERVAL = 5 * time.Second // Time between attempts
)

type Config struct {
	Dsn         *Dsn
	maxAttempts int
	interval    time.Duration
	Logger      *log.Logger
}

func NewConfig(dsn *Dsn, maxAttempts int, interval time.Duration) *Config {
	return &Config{Dsn: dsn, maxAttempts: maxAttempts, interval: interval}
}

func DefaultConfig() *Config {
	// Create logger
	logger := log.Default()

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
		logger.Printf("Adapter configured to attempt connection every %d seconds to mysql", interval/time.Second)
	}

	// Create new config
	c := &Config{}

	// Load DSN from env
	c.Dsn = c.dsnFromDefaults()
	c.SetInterval(interval).
		SetLogger(logger).
		SetMaxAttempts(maxAttempts)

	return c
}

func (c *Config) GetDsn() *Dsn {
	return c.Dsn
}

func (c *Config) GetMaxAttempts() int {
	return c.maxAttempts
}

func (c *Config) SetMaxAttempts(maxAttempts int) *Config {
	c.maxAttempts = maxAttempts

	return c
}

func (c *Config) GetInterval() time.Duration {
	return c.interval
}

func (c *Config) SetInterval(interval time.Duration) *Config {
	c.interval = interval

	return c
}

func (c *Config) SetLogger(logger *log.Logger) *Config {
	c.Logger = logger

	return c
}

func (c *Config) dsnFromDefaults() *Dsn {
	// Get MySQL envs
	user := os.Getenv("MYSQL_USER")
	if user != "" {
		if c.Logger != nil {
			c.Logger.Printf("Read MySQL user from MYSQL_USER env: %s", user)
		}
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password != "" {
		if c.Logger != nil {
			c.Logger.Printf("Read MySQL password from MYSQL_PASSWORD env")
		}
	}

	host := os.Getenv("MYSQL_HOST")
	if host != "" {
		if c.Logger != nil {
			c.Logger.Printf("Read MySQL host from MYSQL_HOST env %s", host)
		}
	}

	port := os.Getenv("MYSQL_PORT")
	if port != "" {
		if c.Logger != nil {
			c.Logger.Printf("Read MySQL port from MYSQL_PORT env %s", port)
		}
	}

	database := os.Getenv("MYSQL_DATABASE")
	if database != "" {
		if c.Logger != nil {
			c.Logger.Printf("Read MySQL database from MYSQL_DATABASE env %s", database)
		}
	}

	return NewDsn(user, password, database).SetHost(host).SetPort(port)
}
