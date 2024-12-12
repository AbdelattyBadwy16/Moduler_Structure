package brand

import (
	"project/Database"
	"project/Database/models"
)

func storeBrandRow(row models.Brand) (models.Brand, error) {
	db := Database.DB //need to call db var here
	err := db.Save(&row).Error
	return row, err
}

func AllBrand() ([]models.Brand, error) {
	db := Database.DB
	var brands []models.Brand
	err := db.Find(&brands).Error
	return brands, err
}

func DeleteBrandRow(request DeleteRequest) (error) {
	db := Database.DB
	var brand models.Brand
	err := db.First(&brand,request.Id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&brand).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateBrandRow(row models.Brand) (models.Brand,error){
	db := Database.DB
	err := db.Save(&row).Error
	return row, err
}