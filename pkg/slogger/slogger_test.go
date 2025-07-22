package slogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlogger(t *testing.T) {
	t.Run("Test that logger initialises an instance", func(t *testing.T) {
		assert.NotNil(t, Instance)
	})
	t.Run("Test that logger can log info", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Info("testing", "test slogger message", "key", "value")
		})
	})
	t.Run("Test that logger can log warn", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Warn("testing", "test slogger message", "key", "value")
		})
	})
	t.Run("Test that logger can log error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Error("testing", "test slogger message", "key", "value")
		})
	})
	t.Run("Test that logger can log & panic on fata", func(t *testing.T) {
		assert.Panics(t, func() {
			Fatal("testing", "test slogger message", "key", "value")
		})
	})
}
