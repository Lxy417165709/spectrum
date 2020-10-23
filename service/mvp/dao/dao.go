package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"spectrum/common/env"
	"spectrum/common/logger"
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
		mainDB.CreateTable(table)
	}
}
func universalCreate(value interface{}) error {
	createTableWhenNotExist(value)
	db := mainDB.Create(value)
	if err := db.Error; err != nil {
		logger.Info("Fail to finish mainDB.Create",
			zap.Any("value", value),
			zap.Error(err))
		return err
	}
	return nil
}
