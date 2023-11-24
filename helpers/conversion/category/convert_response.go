package conversion

import (
	"capstone/features/category"
	"capstone/features/category/dto"
)

func CategoryCreateResponse(res *category.Category) dto.ResCategoryCreate {
	return dto.ResCategoryCreate{
		ID: res.ID,
		Name: res.Name,
	}
}