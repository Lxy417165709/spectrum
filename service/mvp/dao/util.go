package dao

import (
	"strings"
)

func GetPlaceholderClause(count int) string {
	var placeholders []string
	for i := 0; i < count; i++ {
		placeholders = append(placeholders, "?")
	}
	return  strings.Join(placeholders, ",")
}
