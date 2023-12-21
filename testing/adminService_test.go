package mocks_test

import (
	"testing"

	"capstone/features/admin"
	"capstone/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	repo.EXPECT().Delete(1).Return(nil)

	err := repo.Delete(1)
	assert.NoError(t, err)
}

func TestFindAllAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	expectedAdmins := []admin.Admin{
		// ... populate with sample admins for testing
	}
	repo.EXPECT().FindAll().Return(expectedAdmins, nil)

	admins, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, admins, len(expectedAdmins))
}

func TestFindByEmailAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	expectedAdmin := &admin.Admin{
		// ... populate with a sample admin for testing
	}
	repo.EXPECT().FindByEmail("sample@email.com").Return(expectedAdmin, nil)

	admin, err := repo.FindByEmail("sample@email.com")
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestFindByIdAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	expectedAdmin := &admin.Admin{
		// ... populate with a sample admin for testing
	}
	repo.EXPECT().FindById(1).Return(expectedAdmin, nil)

	admin, err := repo.FindById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestFindByUsernameAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	expectedAdmin := &admin.Admin{
		// ... populate with a sample admin for testing
	}
	repo.EXPECT().FindByUsername("sampleUsername").Return(expectedAdmin, nil)

	admin, err := repo.FindByUsername("sampleUsername")
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestSaveAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	expectedAdmin := &admin.Admin{
		// ... populate with a sample admin for testing
	}
	repo.EXPECT().Save(expectedAdmin).Return(expectedAdmin, nil)

	admin, err := repo.Save(expectedAdmin)
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestUpdatePasswordAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockAdminRepository(ctrl)
	adminToUpdate := &admin.Admin{
		// ... populate with a sample admin for testing
	}
	repo.EXPECT().UpdatePassword(adminToUpdate, 123456).Return(adminToUpdate, nil)

	admin, err := repo.UpdatePassword(adminToUpdate, 123456)
	assert.NoError(t, err)
	assert.Equal(t, adminToUpdate, admin)
}

// Additional tests can be added to handle error scenarios, etc.
