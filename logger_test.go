package logger

import "testing"

func TestNewLogger(t *testing.T) {
	NewLogger("debug", "logs", "test.log", true, true)
}
