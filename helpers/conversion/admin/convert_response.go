package conversion

import (
	"capstone/features/admin"
	"capstone/features/admin/dto"
)

func AdminRegisterResponse(res *admin.Admin) dto.ResAdminRegister {
	return dto.ResAdminRegister{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
	}
}

func AdminLoginResponse(res *admin.Admin) dto.ResAdminLogin {
	return dto.ResAdminLogin{
		ID:       res.ID,
		Username: res.Username,
		Email:    res.Email,
	}
}
