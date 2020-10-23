package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"reflect"
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
		trueType := reflect.TypeOf(table)
		emptyTable := reflect.New(trueType).Interface() // 之所以要获得 emptyTable，是因为如果 table有数据，create 语句会报错..
		mainDB.CreateTable(emptyTable)
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

func universalGet(id int, obj interface{}) (interface{}, error) {
	createTableWhenNotExist(obj)
	db := mainDB.First(obj, "id = ?", id)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.first", zap.Int("id", id), zap.Error(err))
		return nil, err
	}
	return obj, nil
}
