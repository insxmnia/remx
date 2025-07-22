package termc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTermc(t *testing.T) {
	t.Run("Test that termc can register custom colours", func(t *testing.T) {
		// Assert
		err := RegisterCustom("test", "\033[42m\033[97m")
		assert.NoError(t, err)
	})
	t.Run("Test that termc can't register custom colours with same name (overwrite)", func(t *testing.T) {
		//Assert
		err := RegisterCustom("test", "\033[42m\033[97m")
		assert.Error(t, err)
	})
	t.Run("Test that termc outputs colours", func(t *testing.T) {
		// Arrange
		expected := "\033[97m\033[41m"

		//Assert
		assert.Equal(t, expected, Red)
	})
}
