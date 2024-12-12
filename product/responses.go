package product

import "project/Database/models"

type ProductResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	BrandID uint   `json:"brand_id"`
	Brand   models.Brand
}
