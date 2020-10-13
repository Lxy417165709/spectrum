package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mainDB *gorm.DB

func init() {
	mainDB = NewDB(GenDSN("localhost:3306", "mysql"), 5)
}

func GenDSN(host, db string) string {
	return fmt.Sprintf("%v/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", host, db)
}

func NewDB(dsn string, maxCon int) *gorm.DB {
	con, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Got error when connecting database, the error is '%s'", err))
	}
	idle := maxCon
	if maxCon/3 > 10 {
		idle = maxCon / 3
	}
	con.DB().SetMaxOpenConns(maxCon)
	con.DB().SetMaxIdleConns(idle)
	con.LogMode(true) // 开启sql日志
	con.BlockGlobalUpdate(true)
	return con
}
