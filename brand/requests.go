package brand

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type DeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type UpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
