package product

type CreateRequest struct {
	Name    string `json:"name" validate:"required"`
	BrandID uint   `json:"brand_id" validate:"required"`
}

type UpdateRequest struct {
	Id      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	BrandID uint   `json:"brand_id" validate:"required"`
}

type DeleteRequest struct {
	Id uint `json:"id" validate:"required"`
}
