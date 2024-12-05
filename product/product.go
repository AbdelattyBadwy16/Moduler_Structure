package product

type Product struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	BrandID uint   `json:"brand_id"`
	Brand   Brand
	Tags    []Tag `gorm:"many2many:product_tags;"`
}
