package conversion

import (
	"capstone/features/admin"
	"capstone/features/admin/dto"
)

func AdminRegisterResponse(res *admin.Admin) dto.ResAdminRegister {
	return dto.ResAdminRegister{
		Username: res.Username,
		Email:    res.Email,
	}
}

func AdminLoginResponse(res *admin.Admin) dto.ResAdminLogin {
	return dto.ResAdminLogin{
		Username: res.Username,
		Email:    res.Email,
	}
}
