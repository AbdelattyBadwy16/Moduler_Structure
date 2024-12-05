package brand

type Brand struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"foreignKey:BrandID"`
}
