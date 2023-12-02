package conversion

import (
	"capstone/features/menu"
	"capstone/features/menu/dto"
)

func MenuCreateRequest(req dto.ReqMenuCreate) *menu.Menu {
	return &menu.Menu{
		Image:       req.Image,
		Name:        req.Name,
		CategoryID:  req.CategoryID,
		Description: req.Description,
		Price:       req.Price,
	}
}

func MenuUpdateRequest(req dto.ReqMenuUpdate) *menu.Menu {
	return &menu.Menu{
		Image:       req.Image,
		Name:        req.Name,
		CategoryID:  uint(req.CategoryID),
		Description: req.Description,
		Price:       req.Price,
	}
}
