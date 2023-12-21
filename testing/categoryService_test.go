package mocks_test

import (
	"testing"

	"capstone/features/category"
	"capstone/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCategoryRepository(ctrl)
	repo.EXPECT().Delete(1).Return(nil)

	err := repo.Delete(1)
	assert.NoError(t, err)
}

func TestFindAllCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCategoryRepository(ctrl)
	expectedCategories := []category.Category{
		// ... populate with sample categories for testing
	}
	repo.EXPECT().FindAll().Return(expectedCategories, nil)

	categories, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, categories, len(expectedCategories))
}

func TestFindByIdCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCategoryRepository(ctrl)
	expectedCategory := &category.Category{
		// ... populate with sample category for testing
	}
	repo.EXPECT().FindById(1).Return(expectedCategory, nil)

	category, err := repo.FindById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestFindByNameCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCategoryRepository(ctrl)
	expectedCategory := &category.Category{
		// ... populate with sample category for testing
	}
	repo.EXPECT().FindByName("SampleName").Return(expectedCategory, nil)

	category, err := repo.FindByName("SampleName")
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestSaveCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCategoryRepository(ctrl)
	expectedCategory := &category.Category{
		// ... populate with sample category for testing
	}
	repo.EXPECT().Save(expectedCategory).Return(expectedCategory, nil)

	category, err := repo.Save(expectedCategory)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}
