package repository

import (
	"capstone/features/user"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Save(Newuser *user.User) (*user.User, error){
	result := r.db.Create(&Newuser)
	if result.Error != nil{
		return nil, result.Error
	}
	return Newuser, nil
}

func (r *UserRepositoryImpl) Update(Newuser *user.User, id int) (*user.User, error) {
	result := r.db.Table("users").Save(Newuser)
	if result.Error != nil {
		return nil, result.Error
	}
	return Newuser, nil
}

func (r *UserRepositoryImpl) UpdatePassword(Newuser *user.User, id int) (*user.User, error){
	result := r.db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"password": Newuser.Password})
	if result.Error != nil {
		return nil, result.Error
	}
	return Newuser, nil
}

func (r *UserRepositoryImpl) FindAll() ([]user.User, error){
	user := []user.User{}

	result := r.db.Where("Delete_at IS NULL").Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*user.User, error){
	user := user.User{}

	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
	
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*user.User, error){
	user := user.User{}

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindById(id int) (*user.User, error){
	user := user.User{}

	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Delete(id int) error{
	result := r.db.Delete(&user.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
