package stringtools

import (
	"strings"
)

func Contains(str string, items []string) bool {
	for _, item := range items {
		if find := strings.Contains(str, item); find {
			return true
		}
	}
	return false
}
