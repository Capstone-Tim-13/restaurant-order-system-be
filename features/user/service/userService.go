package service

import (
	"capstone/features/user"
	"capstone/features/user/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/user"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserServiceImpl struct {
	UserRepository user.Repository
	Validate       *validator.Validate
}

func NewUserservice(userRepository user.Repository, validate *validator.Validate) user.Service {
	return &UserServiceImpl{UserRepository: userRepository, Validate: validate}
}

func (s *UserServiceImpl) Register(ctx echo.Context, req dto.ReqUserRegister) (*user.User, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check Username already exists
	username, _ := s.UserRepository.FindByUsername(req.Username)
	if username != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Check Email already exists
	email, _ := s.UserRepository.FindByEmail(req.Email)
	if email != nil {
		return nil, fmt.Errorf("email already exists")
	}

	user := conversion.UserRegisterRequest(req)

	// Convert Password to Hash
	user.Password = helpers.HashPassword(user.Password)

	_, err = s.UserRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("error When to Register: %s", err.Error())
	}

	results, _ := s.UserRepository.FindByUsername(req.Username)

	return results, nil
}

func (s *UserServiceImpl) Login(ctx echo.Context, req dto.ReqUserLogin) (*user.User, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check Email already exists
	email, _ := s.UserRepository.FindByEmail(req.Email)
	if email == nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Convert Request to Models
	user := conversion.UserLoginRequest(req)

	// Compare Password
	if email != nil {
		err = helpers.ComparePassword(email.Password, user.Password)
		if err != nil {
			return nil, fmt.Errorf("invalid Email and Password")
		}
	} else {
		return nil, fmt.Errorf("email not found")
	}

	return email, nil
}

func (s *UserServiceImpl) FindAll(ctx echo.Context) ([]user.User, error) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return users, nil
}

func (s *UserServiceImpl) FindById(ctx echo.Context, id int) (*user.User, error) {
	user, _ := s.UserRepository.FindById(id)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *UserServiceImpl) UpdatePassword(ctx echo.Context, req dto.ReqUserUpdatePass, id int) (*user.User, error) {
	// Check if the request is valid
	err := s.Validate.Struct(req)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	// Check if the user exists
	existingUser, err := s.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check if there's a new password provided
	if req.Password != "" {
		if len(req.Password) < 8 {
			return nil, fmt.Errorf("password must be at least 8 characters long")
		}
		// Convert request to models
		userToUpdate := conversion.UserUpdateRequestPass(req)
		userToUpdate.Password = helpers.HashPassword(userToUpdate.Password)

		// Update the password in the repository
		updatedUser, err := s.UserRepository.UpdatePassword(userToUpdate, id)
		if err != nil {
			return nil, err
		}

		return updatedUser, nil
	}

	// If no new password provided, return an error or handle as needed
	return nil, fmt.Errorf("no new password provided")
}

func (s *UserServiceImpl) Delete(ctx echo.Context, id int) error {
	// Check Admin already exists
	users, _ := s.UserRepository.FindById(id)
	if users == nil {
		return fmt.Errorf("user Not Found")
	}

	err := s.UserRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}
