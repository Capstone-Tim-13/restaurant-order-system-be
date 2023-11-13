package conversion

import (
	"capstone/features/admin"
	"capstone/features/admin/dto"
)

func AdminRegisterRequest(req dto.ReqAdminRegister) *admin.Admin {
	return &admin.Admin{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func AdminLoginRequest(req dto.ReqAdminLogin) *admin.Admin {
	return &admin.Admin{
		Email:    req.Email,
		Password: req.Password,
	}
}

func AdminUpdateRequest(req dto.ReqAdminUpdate) *admin.Admin {
	return &admin.Admin{
		Password: req.Password,
	}
}
