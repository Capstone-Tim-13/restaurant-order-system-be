package repository

import (
	"capstone/features/category"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.Repository {
	return &CategoryRepositoryImpl{db: db}
}

func (r *CategoryRepositoryImpl) Save(Newcategory *category.Category) (*category.Category, error) {
	result := r.db.Create(&Newcategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return Newcategory, nil
}

func (r *CategoryRepositoryImpl) FindAll() ([]category.Category, error) {
	category := []category.Category{}

	result := r.db.Where("delete_at IS NULL").Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) FindById(id int) (*category.Category, error) {
	category := &category.Category{}

	result := r.db.Where("delete_at IS NULL").First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) FindByName(Name string) (*category.Category, error) {
	category := &category.Category{}

	result := r.db.Where("delete_at IS NULL").First(&category, Name)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) Delete(id int) error {
	result := r.db.Delete(&category.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
