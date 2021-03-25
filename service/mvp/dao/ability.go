package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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