package dto

type ResMenuCreate struct {
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	CategoryID  uint    `json:"categoryid"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ResMenuUpdate struct {
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	CategoryID  uint    `json:"categoryid"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
