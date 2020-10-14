package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"test/service/mvp/model"
	"time"
)

var PlayRecordDao playRecordDao

type playRecordDao struct{}

func (playRecordDao) Create(deskID int, beginPlayTime time.Time) error {
	createTableWhenNotExist(&model.PlayRecord{})

	db := mainDB.Create(&model.PlayRecord{
		BilliardDeskID: deskID,
		BeginPlayTime:  beginPlayTime,
		StopPlayTime:   beginPlayTime,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (playRecordDao) Get(deskID int, beginPlayTime time.Time) (*model.PlayRecord, error) {
	createTableWhenNotExist(&model.PlayRecord{})

	var record model.PlayRecord
	db := mainDB.First(&record, "billiard_desk_id = ? and begin_play_time = ?", deskID, beginPlayTime)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &record, nil
}

func (playRecordDao) UpdateBeginPlayTime(deskID int, beginPlayTime time.Time) error {
	createTableWhenNotExist(&model.PlayRecord{})
	db := mainDB.Table(model.PlayRecord{}.TableName())
	db = db.Where("billiard_desk_id = ? and begin_play_time = ?", deskID, beginPlayTime).Update(&model.PlayRecord{
		BeginPlayTime: beginPlayTime,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (playRecordDao) UpdateStopPlayTime(deskID int, beginPlayTime time.Time, stopPlayTime time.Time) error {
	createTableWhenNotExist(&model.PlayRecord{})

	db := mainDB.Table(model.PlayRecord{}.TableName())
	db = db.Where("billiard_desk_id = ? and begin_play_time = ?", deskID, beginPlayTime).Update(&model.PlayRecord{
		StopPlayTime: stopPlayTime,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
