package adapter

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestNewConfig(t *testing.T) {
	dsn := "foo"
	interval := 3 * time.Second
	attempts := 10
	config := NewConfig(dsn, attempts, interval)

	assert.Equal(t, config.GetDsn(), dsn)
	assert.Equal(t, config.GetInterval(), interval)
	assert.Equal(t, config.GetMaxAttempts(), attempts)
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	assert.Equal(t, config.GetDsn(), getDefaultDsn())
	assert.Equal(t, config.GetInterval(), DEFAULT_MYSQL_ATTEMPT_INTERVAL)
	assert.Equal(t, config.GetMaxAttempts(), DEFAULT_MYSQL_MAX_ATTEMPTS)
}

func TestGetDefaultDsn(t *testing.T) {
	defaultDsn := ":@tcp(:)/?charset=utf8mb4&parseTime=True&loc=Local"
	assert.Equal(t, defaultDsn, getDefaultDsn())
}
