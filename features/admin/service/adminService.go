package service

import (
	"capstone/features/admin"
	"capstone/features/admin/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/admin"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AdminServiceImpl struct {
	AdminRepository admin.Repository
	Validate        *validator.Validate
}

func NewAdminService(adminRepository admin.Repository, validate *validator.Validate) admin.Service {
	return &AdminServiceImpl{AdminRepository: adminRepository, Validate: validate}
}

func (s *AdminServiceImpl) Register(ctx echo.Context, req dto.ReqAdminRegister) (*admin.Admin, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check Username already exists
	username, _ := s.AdminRepository.FindByUsername(req.Username)
	if username != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Check Email already exists
	email, _ := s.AdminRepository.FindByEmail(req.Email)
	if email != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Convert Request to Models
	admin := conversion.AdminRegisterRequest(req)

	// Convert Password to Hash
	admin.Password = helpers.HashPassword(admin.Password)

	_, err = s.AdminRepository.Save(admin)
	if err != nil {
		return nil, fmt.Errorf("error When to Register: %s", err.Error())
	}

	results, _ := s.AdminRepository.FindByUsername(req.Username)

	return results, nil
}

func (s *AdminServiceImpl) Login(ctx echo.Context, req dto.ReqAdminLogin) (*admin.Admin, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check Email already exists
	email, _ := s.AdminRepository.FindByEmail(req.Email)
	if email == nil {
		return nil, fmt.Errorf("email not found")
	}

	// Convert Request to Models
	admin := conversion.AdminLoginRequest(req)

	// Compare Password
	if email != nil {
		err = helpers.ComparePassword(email.Password, admin.Password)
		if err != nil {
			return nil, fmt.Errorf("invalid Email and Password")
		}
	} else {
		return nil, fmt.Errorf("email not found")
	}

	return email, nil
}

func (s *AdminServiceImpl) FindAll(ctx echo.Context) ([]admin.Admin, error) {
	admins, err := s.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("admin not found")
	}

	return admins, nil
}

func (s *AdminServiceImpl) FindById(ctx echo.Context, id int) (*admin.Admin, error) {
	admin, _ := s.AdminRepository.FindById(id)
	if admin == nil {
		return nil, fmt.Errorf("admin not found")
	}
	return admin, nil
}

func (s *AdminServiceImpl) UpdatePassword(ctx echo.Context, req dto.ReqAdminUpdate, id int) (*admin.Admin, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check if the admin exists
	existingAdmin, err := s.AdminRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	// Check if there's a new password provided
	if req.Password != "" {
		if len(req.Password) < 8 {
			return nil, fmt.Errorf("password must be at least 8 characters long")
		}
		// Convert request to models
		adminToUpdate := conversion.AdminUpdateRequest(req)
		adminToUpdate.Password = helpers.HashPassword(adminToUpdate.Password)

		// Update the password in the repository
		updatedAdmin, err := s.AdminRepository.UpdatePassword(adminToUpdate, id)
		if err != nil {
			return nil, err
		}

		return updatedAdmin, nil
	}

	// If no new password provided, return an error or handle as needed
	return nil, fmt.Errorf("no new password provided")
}

func (s *AdminServiceImpl) Delete(ctx echo.Context, id int) error {
	// Check Admin already exists
	admins, _ := s.AdminRepository.FindById(id)
	if admins == nil {
		return fmt.Errorf("admin Not Found")
	}

	err := s.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
