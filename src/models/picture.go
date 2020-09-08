package models

const picNotExist = "record not found"

type Picture struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	CreatedAt  int64  `json:"created_at"`
	OriginName string `json:"origin_name"`
	UploadName string `json:"upload_name"`
}

func PictureInsertOne(msg Picture) error {
	// 查询记录
	res := DB.Where("upload_name = ?", msg.UploadName).First(&Picture{})
	if res.Error == nil {
		return nil
	}
	if res.Error.Error() != picNotExist {
		return res.Error
	}
	// 创建记录
	create := DB.Create(&msg)
	return create.Error
}

func PictureFindAll() ([]Picture, error) {
	var res []Picture
	find := DB.Find(&res)
	return res, find.Error
}

func PictureDeleteOne(msg Picture) error {
	// 查询记录
	res := DB.Where("upload_name = ?", msg.UploadName).First(&msg)
	if res.Error != nil {
		if res.Error.Error() == picNotExist {
			return nil
		}
		return res.Error
	}
	// 删除记录
	del := DB.Delete(&msg)
	return del.Error
}
