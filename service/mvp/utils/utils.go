package utils

import (
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"strconv"
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
