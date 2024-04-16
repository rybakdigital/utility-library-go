package adapter

import (
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewAdapter(t *testing.T) {
	conf := &Config{}
	l := log.Default()
	l.SetPrefix("mysql-adapter")
	a := NewAdapter(conf, l)
	assert.Equal(t, a.IsConnected, false)
	assert.Equal(t, a.Config, conf)
	assert.Equal(t, a.Logger, l)
}

func TestDefaultAdapter(t *testing.T) {
	a := DefaultAdapter()
	assert.Equal(t, "mysql-adapter", a.Logger.Prefix())
}

func TestIsDbConnectionAvailableError(t *testing.T) {
	a := DefaultAdapter()
	_, err := a.IsDbConnectionAvailable()
	assert.NotEqual(t, err, nil)
}
