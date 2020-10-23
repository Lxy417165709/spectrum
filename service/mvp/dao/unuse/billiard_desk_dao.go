package unuse

//var BilliardDeskDao billiardDeskDao
//
//type billiardDeskDao struct{}
//
//func (billiardDeskDao) GetByName(goodName string) (*model.BilliardDesk, error) {
//	dao.createTableWhenNotExist(&model.BilliardDesk{})
//
//	var desk model.BilliardDesk
//	db := dao.mainDB.First(&desk, "name = ?", goodName)
//	if err := db.Error; err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			return nil, nil
//		}
//		logs.Error(err)
//		return nil, err
//	}
//	return &desk, nil
//}
//
//func (billiardDeskDao) Create(deskName string) error {
//	dao.createTableWhenNotExist(&model.BilliardDesk{})
//
//	db := dao.mainDB.Create(&model.BilliardDesk{
//		Name: deskName,
//	})
//	if err := db.Error; err != nil {
//		logs.Error(err)
//		return err
//	}
//	return nil
//}
