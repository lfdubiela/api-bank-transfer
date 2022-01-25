package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebug(t *testing.T) {
	assert.NotPanics(t, func() {
		Debug("Test messages", "test", nil)
	})
}

func TestInfo(t *testing.T) {
	assert.NotPanics(t, func() {
		Info("Test messages", "test", nil)
	})
}

func TestWarn(t *testing.T) {
	assert.NotPanics(t, func() {
		Warn("Test messages", "test", nil)
	})
}

func TestError(t *testing.T) {
	assert.NotPanics(t, func() {
		Error("Test messages", "test", nil)
	})
}
