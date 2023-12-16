package dto

type ReqUserRegister struct {
	Username  string `json:"username" form:"username" validate:"required,min=1"`
	Email     string `json:"email" form:"email" validate:"required,min=1"`
	Password  string `json:"password" form:"password" validate:"required,min=8"`
}

type ReqUserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,min=1"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type ReqUserUpdatePass struct {
	Password string `json:"password" form:"password" validate:"required,min=8"`
}
