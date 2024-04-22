package adapter

import (
	"testing"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	"github.com/go-playground/assert/v2"
)

func TestNewAdapter(t *testing.T) {
	conf := &Config{}
	l := log.NewLogger("mysql-adapter")
	a := NewAdapter(conf, l)
	assert.Equal(t, a.IsConnected, false)
	assert.Equal(t, a.Config, conf)
	assert.Equal(t, a.Logger, l)
}

func TestDefaultAdapter(t *testing.T) {
	a := DefaultAdapter()
	assert.Equal(t, "mysql-adapter", a.Logger.Module)
}

func TestIsDbConnectionAvailableError(t *testing.T) {
	a := DefaultAdapter()
	_, err := a.IsDbConnectionAvailable()
	assert.NotEqual(t, err, nil)
}
