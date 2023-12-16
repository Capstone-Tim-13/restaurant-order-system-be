package dto

type ResUserRegister struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	NoHp      string `json:"no_hp"`
	BirthDate string `json:"birth_date"`
}

type ResUserLogin struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
