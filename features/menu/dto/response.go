package dto

type ResMenuCreate struct {
	ID          uint    `json:"id"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	CategoryID  uint    `json:"categoryid"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Status      bool    `json:"status"`
}

type ResMenuUpdate struct {
	ID          uint    `json:"id"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	CategoryID  uint    `json:"categoryid"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Status      bool    `json:"status"`
}
