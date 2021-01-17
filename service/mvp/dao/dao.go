package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
	"spectrum/common/env"
	"spectrum/common/utils"
)

var mainDB *gorm.DB

func InitMainDB() {
	mainDB = utils.NewDB(
		utils.GenDSN(env.Conf.MainDB.Link, env.Conf.MainDB.Name),
		env.Conf.MainDB.MaxConn,
	)
}

func createTableWhenNotExist(table interface{}) {
	if !mainDB.HasTable(table) {
		trueType := reflect.TypeOf(table)
		emptyTable := reflect.New(trueType).Interface() // 之所以要获得 emptyTable，是因为如果 table有数据，create 语句会报错..
		mainDB.CreateTable(emptyTable)
	}
}