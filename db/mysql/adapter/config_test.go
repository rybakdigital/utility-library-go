package adapter

import (
	"os"
	"strconv"
	"testing"
	"time"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	"github.com/go-playground/assert/v2"
)

func TestNewConfig(t *testing.T) {
	dsn := NewDsn("john", "secret", "foo")
	interval := 3 * time.Second
	attempts := 10
	config := NewConfig(dsn, attempts, interval)
	assert.Equal(t, config.GetDsn(), dsn)
	assert.Equal(t, config.GetInterval(), interval)
	assert.Equal(t, config.GetMaxAttempts(), attempts)
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	assert.Equal(t, config.GetDsn(), config.dsnFromDefaults())
	assert.Equal(t, config.GetInterval(), DEFAULT_MYSQL_ATTEMPT_INTERVAL)
	assert.Equal(t, config.GetMaxAttempts(), DEFAULT_MYSQL_MAX_ATTEMPTS)
}

func TestGetDefaultDsn(t *testing.T) {
	user := "john"
	pass := "pass"
	host := "mysql"
	port := "3333"
	db := "foo"
	defaultDsn := NewDsn(user, pass, db).SetHost(host).SetPort(port)

	os.Setenv("MYSQL_USER", user)
	os.Setenv("MYSQL_PASSWORD", pass)
	os.Setenv("MYSQL_HOST", host)
	os.Setenv("MYSQL_PORT", port)
	os.Setenv("MYSQL_DATABASE", db)
	c := &Config{Logger: log.NewLogger("mysql-config")}
	assert.Equal(t, defaultDsn, c.dsnFromDefaults())
}

func TestDefaultConfigCustomEnv(t *testing.T) {
	attempts := 12
	os.Setenv("MYSQL_MAX_ATTEMPTS", strconv.Itoa(attempts))
	os.Setenv("MYSQL_ATTEMPT_INTERVAL", strconv.Itoa(attempts))
	config := DefaultConfig()
	assert.Equal(t, config.GetMaxAttempts(), attempts)
	assert.Equal(t, config.GetInterval(), time.Duration(attempts)*time.Second)
}
