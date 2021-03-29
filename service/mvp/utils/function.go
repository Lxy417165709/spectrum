package utils

import (
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"strconv"
	"strings"
	"time"
)

func GetPbPrice(price float64) string {
	return fmt.Sprintf("%0.2f", price)
}

func GetDbPrice(priceString string) float64 {
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		logger.Error("Fail to finish strconv.ParseFloat", zap.String("priceString", priceString))
	}
	return price
}

func ToTime(timestamp int64) time.Time {
	if timestamp != 0 {
		return time.Unix(timestamp, 0)
	}
	return NilTime
}

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

func Int64sToInterfaces(elements []int64) []interface{} {
	var interfaces []interface{}
	for _, element := range elements {
		interfaces = append(interfaces, element)
	}
	return interfaces
}

func NewBlankGoods() []*pb.Good {
	return []*pb.Good{}
}

func NewBlankGood() *pb.Good {
	return &pb.Good{
		Id:             0,
		MainElement:    NewBlankElement(),
		AttachElements: NewBlankElements(),
		Favors:         NewBlankFavors(),
		ExpenseInfo:    NewBlankExpenseInfo(),
	}
}

func NewBlankElements() []*pb.Element {
	return []*pb.Element{}
}
func NewBlankElement() *pb.Element {
	return &pb.Element{
		Id:            0,
		Name:          "",
		Type:          0,
		SizeInfos:     NewSizeInfos(),
		SelectedIndex: 0,
	}
}

func NewSizeInfos() []*pb.SizeInfo {
	return []*pb.SizeInfo{}
}

func NewBlankDesk() *pb.Desk {
	return &pb.Desk{
		Id:          0,
		Space:       NewBlankSpace(),
		StartAt:     0,
		EndAt:       0,
		Favors:      NewBlankFavors(),
		ExpenseInfo: NewBlankExpenseInfo(),
		OrderID:     0,
	}
}

func NewBlankFavors() []*pb.Favor {
	return []*pb.Favor{}
}

func NewBlankExpenseInfo() *pb.ExpenseInfo {
	return &pb.ExpenseInfo{
		NonFavorExpense: 0,
		CheckOutAt:      NilTime.Unix(),
		Expense:         0,
	}
}

func NewBlankSpace() *pb.Space {
	return &pb.Space{
		Id:               0,
		Name:             "",
		ClassName:        "",
		Price:            "",
		BillingType:      0,
		PictureStorePath: "",
	}
}
