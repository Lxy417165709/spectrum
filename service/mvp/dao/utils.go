package dao

import (
	"reflect"
	"strings"
)

func getPlaceholderClause(count int) string {
	var placeholders []string
	for i := 0; i < count; i++ {
		placeholders = append(placeholders, "?")
	}
	return strings.Join(placeholders, ",")
}

func createTableWhenNotExist(table interface{}) {
	if !mainDB.HasTable(table) {
		trueType := reflect.TypeOf(table)
		emptyTable := reflect.New(trueType).Interface() // 之所以要获得 emptyTable，是因为如果 table有数据，create 语句会报错..
		mainDB.CreateTable(emptyTable)
	}
}
