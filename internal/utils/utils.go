package utils

import (
	"strings"
)

func IsString(input string) bool {
	if strings.HasPrefix(input, "\"") && strings.HasSuffix(input, "\"") {
		return true
	}

	return false
}

func IsRawString(input string) bool {
	if strings.HasPrefix(input, "'") && strings.HasSuffix(input, "'") {
		return true
	}

	return false
}
