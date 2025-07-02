package utils

import "strings"

func IsValidField(fieldValue string) bool {
	return strings.TrimSpace(fieldValue) != ""
}
