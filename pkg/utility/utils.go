package utility

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type JSON map[string]any

func GenerateString(length int, special bool) string {
	if length <= 0 {
		return ""
	}

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialCharacters := "#!£$&*@?="
	allCharacters := letters

	var result []rune

	if special {
		allCharacters = letters + specialCharacters
		for range length {
			result = append(result, rune(allCharacters[rand.Intn(len(allCharacters))]))
		}
		result[0] = rune(specialCharacters[rand.Intn(len(specialCharacters))])
	} else {
		for range length {
			result = append(result, rune(letters[rand.Intn(len(letters))]))
		}
	}

	if len(string(result)) != length {
		return string(result)[0:length]
	}
	return string(result)
}

func IsEmpty(value string) bool {
	return value == "" || len(value) == 0 || strings.ReplaceAll(value, " ", "") == ""
}

func JSONToBytes(obj JSON) []byte {
	if obj == nil {
		return []byte{}
	}

	data, err := json.Marshal(obj)
	if err != nil {
		return []byte{}
	}
	return data
}

func GenerateUUID() string {
	return fmt.Sprintf("%s-%s-%s", GenerateString(8, false), GenerateString(4, false), GenerateString(8, false))
}
