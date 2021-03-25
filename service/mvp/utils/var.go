package utils

import (
	"strings"
)

func GetValidElementTypesString() string {
	var elementTypeStrings []string
	for _, elementType := range ValidElementTypes {
		elementTypeStrings = append(elementTypeStrings, elementType.String())
	}
	return strings.Join(elementTypeStrings, "、")
}

func GetValidFavorTypesString() string {
	var favorTypeStrings []string
	for _, favorType := range ValidFavorTypes {
		favorTypeStrings = append(favorTypeStrings, favorType.String())
	}
	return strings.Join(favorTypeStrings, "、")
}
