package utils

import (
	"spectrum/common/pb"
	"time"
)

const INF = 1E18

const (
	ChargeableObjectNameOfGood  = "good"
	ChargeableObjectNameOfOrder = "order"
	ChargeableObjectNameOfDesk  = "desk"
)

var (
	NilTime = time.Unix(10086, 0)
)

var ValidElementTypes = []pb.ElementType{
	pb.ElementType_Main,
	pb.ElementType_Option,
	pb.ElementType_Ingredient,
}

var ValidFavorTypes = []pb.FavorType{
	pb.FavorType_NONE,
	pb.FavorType_REBATE,
	pb.FavorType_FULL_REDUCTION,
	pb.FavorType_FREE,
}
