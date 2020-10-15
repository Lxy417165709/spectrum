package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"spectrum/service/auth/model"
	"time"
)

var UserDao userDao

type userDao struct{}

func (userDao) Get(userID int) (*model.User, error) {
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

func (userDao) GetByEmail(email string) (*model.User, error) {
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

func (userDao) Create(email, hashSaltyPassword, salt string) error {
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


