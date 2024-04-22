package logger

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	os.Setenv("LOGGER_MODE", PROD)
	log := NewLogger("tester")

	if log.Mode != PROD {
		t.Errorf("NewLogger() failed. The logger mode was not set correctly. Expected mode %s, got %s", log.Mode, PROD)
	}
}
