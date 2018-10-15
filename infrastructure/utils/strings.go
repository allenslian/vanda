package utils

import (
	"strings"
)

//ContainsString checks if the list contains the value.
func ContainsString(list []string, value string) bool {
	cleared := strings.ToLower(strings.TrimSpace(value))
	for _, v := range list {
		item := strings.ToLower(strings.TrimSpace(v))
		if strings.Compare(item, cleared) == 0 {
			return true
		}
	}
	return false
}
