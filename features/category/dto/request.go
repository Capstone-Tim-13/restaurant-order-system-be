package dto

type ReqCategoryCreate struct {
	Name string `json:"name" form:"name" validate:"required,min=1"`
}