package repository

import (
	"capstone/features/admin"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) admin.Repository {
	return &AdminRepositoryImpl{db: db}
}

func (r *AdminRepositoryImpl) Save(Newadmin *admin.Admin) (*admin.Admin, error) {
	result := r.db.Create(&Newadmin)
	if result.Error != nil {
		return nil, result.Error
	}
	return Newadmin, nil
}

func (r *AdminRepositoryImpl) UpdatePassword(Newadmin *admin.Admin, id int) (*admin.Admin, error) {
	result := r.db.Table("admins").Where("id= ?", id).Updates(admin.Admin{Password: Newadmin.Password})
	if result.Error != nil {
		return nil, result.Error
	}
	return Newadmin, nil
}

func (r *AdminRepositoryImpl) FindAll() ([]admin.Admin, error) {
	admin := []admin.Admin{}

	result := r.db.Where("delete_at IS NULL").Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return admin, nil
}

func (r *AdminRepositoryImpl) FindByUsername(username string) (*admin.Admin, error) {
	admin := admin.Admin{}

	result := r.db.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AdminRepositoryImpl) FindByEmail(email string) (*admin.Admin, error) {
	admin := admin.Admin{}

	result := r.db.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AdminRepositoryImpl) FindById(id int) (*admin.Admin, error) {
	admin := admin.Admin{}

	result := r.db.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AdminRepositoryImpl) Delete(id int) error {
	result := r.db.Delete(&admin.Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
