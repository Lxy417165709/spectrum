package unuse

// model
//type DeskGoodRecord struct {
//	gorm.Model
//	DeskID          int    `json:"desk_id"`
//	GoodID          int    `json:"good_id"`
//	MainElementName string `json:"main_element_name"`
//}
//
//func (*DeskGoodRecord) TableName() string {
//	return "desk_good_record"
//}

// dao
//var DeskGoodRecordDao deskGoodRecordDao
//
//type deskGoodRecordDao struct{}
//
//func (deskGoodRecordDao) Create(obj *model.DeskGoodRecord) error {
//	var table model.DeskGoodRecord
//	createTableWhenNotExist(&table)
//
//	if err := mainDB.Create(&obj).Error; err != nil {
//		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
//		return err
//	}
//	return nil
//}
//
//func (deskGoodRecordDao) GetByDeskID(deskID int) ([]*model.DeskGoodRecord, error) {
//	var table model.DeskGoodRecord
//	createTableWhenNotExist(&table)
//
//	var result []*model.DeskGoodRecord
//	if err := mainDB.Find(&result, "desk_id = ?", deskID).Error; err != nil {
//		logger.Error("Fail to finish mainDB.First", zap.Int("deskID", deskID), zap.Error(err))
//		return nil, err
//	}
//	return result, nil
//}
