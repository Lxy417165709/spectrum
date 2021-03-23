package model

import (
	"spectrum/common/pb"
	"strings"
)

var ValidElementTypes = []pb.ElementType{
	pb.ElementType_Main,
	pb.ElementType_Option,
	pb.ElementType_Ingredient,
}

func GetValidElementTypesString() string {
	var elementTypeStrings []string
	for _, elementType := range ValidElementTypes {
		elementTypeStrings = append(elementTypeStrings, elementType.String())
	}
	return strings.Join(elementTypeStrings, "、")
}

var ValidFavorTypes = []pb.FavorType{
	pb.FavorType_NONE,
	pb.FavorType_REBATE,
	pb.FavorType_FULL_REDUCTION,
	pb.FavorType_FREE,
}

func GetValidFavorTypesString() string {
	var favorTypeStrings []string
	for _, favorType := range ValidFavorTypes {
		favorTypeStrings = append(favorTypeStrings, favorType.String())
	}
	return strings.Join(favorTypeStrings, "、")
}
