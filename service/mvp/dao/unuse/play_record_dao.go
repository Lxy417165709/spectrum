package unuse

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"time"
)

var PlayRecordDao playRecordDao

type playRecordDao struct{}

func (playRecordDao) Create(deskID int, beginPlayTime time.Time) error {
	dao.createTableWhenNotExist(&model.PlayRecord{})

	db := dao.mainDB.Create(&model.PlayRecord{
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
	dao.createTableWhenNotExist(&model.PlayRecord{})

	var record model.PlayRecord
	db := dao.mainDB.First(&record, "billiard_desk_id = ? and begin_play_time = ?", deskID, beginPlayTime)
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
	dao.createTableWhenNotExist(&model.PlayRecord{})
	db := dao.mainDB.Table(model.PlayRecord{}.TableName())
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
	dao.createTableWhenNotExist(&model.PlayRecord{})

	db := dao.mainDB.Table(model.PlayRecord{}.TableName())
	db = db.Where("billiard_desk_id = ? and begin_play_time = ?", deskID, beginPlayTime).Update(&model.PlayRecord{
		StopPlayTime: stopPlayTime,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
