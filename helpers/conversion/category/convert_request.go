package conversion

import (
	"capstone/features/category"
	"capstone/features/category/dto"
)

func CategoryCreateRequest(req dto.ReqCategoryCreate) *category.Category{
	return &category.Category{
		Name: req.Name,
	}
}