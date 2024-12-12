package tag

import (
	"project/Database"
	"project/Database/models"
)

func storeTagRow(row models.Tag) (models.Tag, error) {
	db := Database.DB //need to call db var here
	err := db.Save(&row).Error
	return row, err
}

func AllTags() ([]models.Tag, error) {
	db := Database.DB
	var brands []models.Tag
	err := db.Find(&brands).Error
	return brands, err
}

func DeleteTagRow(request DeleteRequest) error {
	db := Database.DB
	var tag models.Tag
	err := db.First(&tag, request.Id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&tag).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTagRow(row models.Tag) (models.Tag, error) {
	db := Database.DB
	err := db.Save(&row).Error
	return row, err
}
