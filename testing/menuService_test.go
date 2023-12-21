package mocks_test

import (
	"context"
	"testing"

	"capstone/features/menu"
	"capstone/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	repo.EXPECT().Delete(1).Return(nil)

	err := repo.Delete(1)
	assert.NoError(t, err)
}

func TestFindAllMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedMenus := []menu.Menu{
		// ... populate with sample menus for testing
	}
	repo.EXPECT().FindAll().Return(expectedMenus, nil)

	menus, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, menus, len(expectedMenus))
}

func TestFindByCategoryIdMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedMenus := []menu.Menu{
		// ... populate with sample menus for testing
	}
	repo.EXPECT().FindByCategoryId(1).Return(expectedMenus, nil)

	menus, err := repo.FindByCategoryId(1)
	assert.NoError(t, err)
	assert.Len(t, menus, len(expectedMenus))
}

func TestFindByIdMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedMenu := &menu.Menu{
		// ... populate with a sample menu for testing
	}
	repo.EXPECT().FindById(1).Return(expectedMenu, nil)

	menu, err := repo.FindById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedMenu, menu)
}

func TestFindByNameMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedMenu := &menu.Menu{
		// ... populate with a sample menu for testing
	}
	repo.EXPECT().FindByName("SampleMenu").Return(expectedMenu, nil)

	menu, err := repo.FindByName("SampleMenu")
	assert.NoError(t, err)
	assert.Equal(t, expectedMenu, menu)
}

func TestSaveMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedMenu := &menu.Menu{
		// ... populate with a sample menu for testing
	}
	repo.EXPECT().Save(expectedMenu).Return(expectedMenu, nil)

	menu, err := repo.Save(expectedMenu)
	assert.NoError(t, err)
	assert.Equal(t, expectedMenu, menu)
}

func TestUpdateMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedMenu := &menu.Menu{
		// ... populate with a sample menu for testing
	}
	repo.EXPECT().Update(expectedMenu).Return(expectedMenu, nil)

	menu, err := repo.Update(expectedMenu)
	assert.NoError(t, err)
	assert.Equal(t, expectedMenu, menu)
}

func TestUploadImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockMenuRepository(ctrl)
	expectedURL := "http://example.com/image.jpg"

	// Using context.TODO() instead of nil
	repo.EXPECT().UploadImage(gomock.Any(), gomock.Any(), "sample.jpg").Return(expectedURL, nil)

	url, err := repo.UploadImage(context.TODO(), nil, "sample.jpg")
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, url)
}

// Additional tests can be added to handle error scenarios, etc.
