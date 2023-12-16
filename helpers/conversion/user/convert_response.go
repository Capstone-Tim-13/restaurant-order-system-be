package conversion

import (
	"capstone/features/user"
	"capstone/features/user/dto"
)

func UserRegisterResponse(res *user.User) dto.ResUserRegister {
	return dto.ResUserRegister{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		NoHp:      res.NoHp,
		BirthDate: res.BirthDate,
	}
}

func UserLoginResponse(res *user.User) dto.ResUserLogin {
	return dto.ResUserLogin{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
	}
}
