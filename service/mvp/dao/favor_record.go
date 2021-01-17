package dao

//var FavorRecordDao favorRecordDao

//type favorRecordDao struct{}

//func (favorRecordDao) Create(obj *model.FavorRecord) error {
//	var table model.FavorRecord
//	createTableWhenNotExist(&table)
//
//	if err := mainDB.Create(obj).Error; err != nil {
//		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
//		return err
//	}
//	return nil
//}

//func (favorRecordDao) Get(favorableStructName string, favorableStructID int64) ([]*model.FavorRecord, error) {
//	var table model.FavorRecord
//	createTableWhenNotExist(&table)
//
//	var result []*model.FavorRecord
//	if err := mainDB.Find(
//		&result,
//		"favorable_struct_name = ? and favorable_struct_id = ？",
//		favorableStructName, favorableStructID,
//	).Error; err != nil {
//		logger.Error(
//			"Fail to finish mainDB.Find",
//			zap.Int64("favorableStructID", favorableStructID),
//			zap.String("favorableStructName", favorableStructName),
//			zap.Error(err))
//		return nil, err
//	}
//	return result, nil
//}
