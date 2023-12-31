package conversion

import (
	"capstone/features/user"
	"capstone/features/user/dto"
)

func UserRegisterRequest(req dto.ReqUserRegister) *user.User {
	return &user.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func UserLoginRequest(req dto.ReqUserLogin) *user.User {
	return &user.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func UserUpdateRequestPass(req dto.ReqUserUpdatePass) *user.User {
	return &user.User{
		Password: req.Password,
	}
}
