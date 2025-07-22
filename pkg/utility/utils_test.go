package utility

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	t.Run("Test that the GenerateString function generates a string of the correct length", func(t *testing.T) {
		// Arrange
		length := 10
		special := false
		// Act
		result := GenerateString(length, special)
		// Assert
		assert.Equal(t, length, len(result))
	})
	t.Run("Test that the GenerateString function generates a string of the correct length with special characters", func(t *testing.T) {
		// Arrange
		length := 10
		special := true
		characters := "#!£$&*@?="
		hasSpecial := false
		// Act
		result := GenerateString(length, special)
		for _, char := range result {
			if strings.ContainsRune(characters, char) {
				hasSpecial = true
				break
			}
		}
		// Assert
		assert.Equal(t, length, len(result))
		assert.True(t, hasSpecial)
	})
	t.Run("Test if isEmpty return true on an empty string", func(t *testing.T) {
		// Arrange
		value := ""
		// Act
		result := IsEmpty(value)
		// Assert
		assert.True(t, result)
	})
	t.Run("Test if isEmpty return false on a non-empty string", func(t *testing.T) {
		// Arrange
		value := "test"
		// Act
		result := IsEmpty(value)
		// Assert
		assert.False(t, result)
	})
	t.Run("Test if isEmpty return true on a string with only spaces", func(t *testing.T) {
		// Arrange
		value := "   "
		// Act
		result := IsEmpty(value)
		// Assert
		assert.True(t, result)
	})
	t.Run("Test if JSONtoBytes will return an empty byte array on an empty JSON object", func(t *testing.T) {
		// Arrange
		json := JSON{}
		// Act
		result := JSONToBytes(json)
		// Assert
		assert.Equal(t, []byte{0x7b, 0x7d}, result)
	})
	t.Run("Test if JSONtoBytes will return a byte array on a JSON object", func(t *testing.T) {
		// Arrange
		json := JSON{
			"test": "test",
		}
		// Act
		result := JSONToBytes(json)
		// Assert
		assert.NotEqual(t, []byte{}, result)
	})
	t.Run("Test if GenerateUUID will return a valid UUID", func(t *testing.T) {
		// Arrange
		var parts []int = []int{8, 4, 8}
		length := parts[0] + parts[1] + parts[2] + 2
		// Act
		result := GenerateUUID()
		// Assert
		assert.NotEqual(t, result, "")
		assert.Contains(t, result, "-")
		assert.Equal(t, length, len(result))
		assert.Equal(t, len(strings.Split(result, "-")), 3)
	})
	t.Run("Test GenerateString with zero length", func(t *testing.T) {
		// Arrange
		length := 0
		special := false
		// Act
		result := GenerateString(length, special)
		// Assert
		assert.Equal(t, "", result)
	})

	t.Run("Test GenerateString with negative length", func(t *testing.T) {
		// Arrange
		length := -5
		special := false
		// Act
		result := GenerateString(length, special)
		// Assert
		assert.Equal(t, "", result)
	})

	t.Run("Test GenerateString with length 2 and special characters", func(t *testing.T) {
		// Arrange
		length := 2
		special := true
		letters := "#!£$&*@?="
		hasSpecial := false
		// Act
		result := GenerateString(length, special)
		for _, char := range result {
			if strings.ContainsRune(letters, char) {
				hasSpecial = true
				break
			}
		}
		// Assert
		assert.Equal(t, length, len(result))
		assert.True(t, hasSpecial)
	})

	t.Run("Test GenerateString with length 1 without special characters", func(t *testing.T) {
		// Arrange
		length := 1
		special := false
		letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		// Act
		result := GenerateString(length, special)
		// Assert
		assert.Equal(t, length, len(result))
		assert.True(t, strings.ContainsRune(letters, rune(result[0])))
	})

	t.Run("Test GenerateString generates different strings on multiple (500) calls", func(t *testing.T) {
		// Arrange
		length := 20
		special := true

		// Assert
		for range 500 {
			// Act
			result := GenerateString(length, special)
			// Assert
			assert.NotEqual(t, result, GenerateString(length, special))
		}
	})

	// IsEmpty additional cases
	t.Run("Test IsEmpty with mixed spaces and tabs", func(t *testing.T) {
		// Arrange
		value := "  \t  \n  "
		// Act
		result := IsEmpty(value)
		// Assert
		assert.False(t, result)
	})

	t.Run("Test IsEmpty with single space", func(t *testing.T) {
		// Arrange
		value := " "
		// Act
		result := IsEmpty(value)
		// Assert
		assert.True(t, result)
	})

	t.Run("Test IsEmpty with string containing spaces in middle", func(t *testing.T) {
		// Arrange
		value := "hello world"
		// Act
		result := IsEmpty(value)
		// Assert
		assert.False(t, result)
	})

	// JSONToBytes additional cases
	t.Run("Test JSONToBytes with nil JSON object", func(t *testing.T) {
		// Arrange
		var json JSON = nil
		// Act
		result := JSONToBytes(json)
		// Assert
		assert.Equal(t, []byte{}, result)
	})

	t.Run("Test JSONToBytes with complex nested JSON", func(t *testing.T) {
		// Arrange
		json := JSON{
			"user": JSON{
				"name": "John",
				"age":  30,
			},
			"active": true,
			"scores": []int{1, 2, 3},
		}
		// Act
		result := JSONToBytes(json)
		// Assert
		assert.NotEqual(t, []byte{}, result)
		assert.Greater(t, len(result), 0)
	})

	t.Run("Test JSONToBytes with JSON containing special characters", func(t *testing.T) {
		// Arrange
		json := JSON{
			"message": "Hello \"World\" with special chars: #!£$&*@?=",
			"unicode": "éïéé⚡",
		}
		// Act
		result := JSONToBytes(json)
		// Assert
		assert.NotEqual(t, []byte{}, result)
		assert.Greater(t, len(result), 0)
	})

	t.Run("Test GenerateUUID format consistency", func(t *testing.T) {
		// Act
		result := GenerateUUID()
		parts := strings.Split(result, "-")
		// Assert
		assert.Equal(t, 3, len(parts))
		assert.Equal(t, 8, len(parts[0]))
		assert.Equal(t, 4, len(parts[1]))
		assert.Equal(t, 8, len(parts[2]))
	})

	t.Run("Test GenerateUUID contains only valid characters", func(t *testing.T) {
		// Arrange
		validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"
		// Act
		result := GenerateUUID()
		// Assert
		for _, char := range result {
			assert.True(t, strings.ContainsRune(validChars, char))
		}
	})

	t.Run("Test GenerateUUID generates unique IDs", func(t *testing.T) {
		// Act
		uuid1 := GenerateUUID()
		uuid2 := GenerateUUID()
		uuid3 := GenerateUUID()
		// Assert
		assert.NotEqual(t, uuid1, uuid2)
		assert.NotEqual(t, uuid2, uuid3)
		assert.NotEqual(t, uuid1, uuid3)
	})

	t.Run("Test GenerateString character set validation without special", func(t *testing.T) {
		// Arrange
		length := 50
		special := false
		letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		containsLetters := false
		// Act
		result := GenerateString(length, special)
		for _, char := range result {
			if strings.ContainsRune(letters, char) {
				containsLetters = true
				break
			}
		}
		// Assert
		assert.True(t, containsLetters)
	})

	t.Run("Test GenerateString character set validation with special", func(t *testing.T) {
		// Arrange
		length := 50
		special := true
		letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#!£$&*@?="
		containsLetters := false
		// Act
		result := GenerateString(length, special)
		for _, char := range result {
			if strings.ContainsRune(letters, char) {
				containsLetters = true
				break
			}
		}
		// Assert
		assert.True(t, containsLetters)
	})
}
