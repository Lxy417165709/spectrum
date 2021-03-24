package dao

import (
	"strings"
)

func getPlaceholderClause(count int) string {
	var placeholders []string
	for i := 0; i < count; i++ {
		placeholders = append(placeholders, "?")
	}
	return strings.Join(placeholders, ",")
}
