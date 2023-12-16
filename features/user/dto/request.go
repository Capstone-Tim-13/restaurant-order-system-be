package dto

type ReqUserRegister struct {
	Username  string `json:"username" form:"username" validate:"required,min=1"`
	Email     string `json:"email" form:"email" validate:"required,min=1"`
	Password  string `json:"password" form:"password" validate:"required,min=8"`
	NoHp      string `json:"no_hp" form:"no_hp" validate:"required,min=12,max=13"`
	BirthDate string `json:"birth_date" form:"birth_date" validate:"required"`
}

type ReqUserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,min=1"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type ReqUserUpdatePass struct {
	Password string `json:"password" form:"password" validate:"required,min=8"`
}
