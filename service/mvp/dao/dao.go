package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"test/common/env"
	"test/common/utils"
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
		mainDB.CreateTable(table)
	}
}
