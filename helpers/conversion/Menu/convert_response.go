package conversion

import (
	"capstone/features/menu"
	"capstone/features/menu/dto"
)

func MenuCreateResponse(res *menu.Menu) dto.ResMenuCreate {
	return dto.ResMenuCreate{
		ID:          res.ID,
		Image:       res.Image,
		Name:        res.Name,
		CategoryID:  res.CategoryID,
		Description: res.Description,
		Price:       res.Price,
		Status:      res.Status,
	}
}

func MenuUpdateResponse(res *menu.Menu) dto.ResMenuUpdate {
	return dto.ResMenuUpdate{
		ID:          res.ID,
		Image:       res.Image,
		Name:        res.Name,
		CategoryID:  res.CategoryID,
		Description: res.Description,
		Price:       res.Price,
		Status:      res.Status,
	}
}
