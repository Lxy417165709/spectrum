package model

import (
	"spectrum/common/pb"
	"time"
)

const (
	NotCheckedOut = 1
	HasCheckedOut = 2
)

const (
	ChargeableObjectNameOfGood  = "good"
	ChargeableObjectNameOfOrder = "order"
	ChargeableObjectNameOfDesk  = "desk"
)

const INF = 1E18

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
