package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"test/auth/model"
	"time"
)

type UserDao struct{}

func (UserDao) Get(userID int) (*model.User, error) {
	createTableWhenNotExist(&model.User{})

	var user model.User
	db := mainDB.First(&user, "id = ?", userID)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &user, nil
}

func (UserDao) GetByEmail(email string) (*model.User, error) {
	createTableWhenNotExist(&model.User{})

	var user model.User
	db := mainDB.First(&user, "email = ?", email)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &user, nil
}

func (UserDao) Create(email, hashSaltyPassword, salt string) error {
	createTableWhenNotExist(&model.User{})
	db := mainDB.Create(&model.User{
		Email:             email,
		HashSaltyPassword: hashSaltyPassword,
		Salt:              salt,
		LastLoginTime:     time.Now(),
		Birthday:          time.Now(),
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func createTableWhenNotExist(table interface{}) {
	if !mainDB.HasTable(table) {
		mainDB.CreateTable(table)
	}
}

