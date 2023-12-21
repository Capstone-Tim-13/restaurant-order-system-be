package mocks_test

import (
	"testing"

	"capstone/features/user"
	"capstone/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().Delete(gomock.Any()).Return(nil)

	err := repo.Delete(1)
	assert.NoError(t, err)
}

func TestFindAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	expectedUsers := []user.User{{ID: 1, Username: "user1"}, {ID: 2, Username: "user2"}}
	repo.EXPECT().FindAll().Return(expectedUsers, nil)

	users, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestFindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	expectedUser := &user.User{ID: 1, Email: "test@example.com"}
	repo.EXPECT().FindByEmail(gomock.Any()).Return(expectedUser, nil)

	userByEmail, err := repo.FindByEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, userByEmail)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	expectedUser := &user.User{ID: 1, Username: "testUser"}
	repo.EXPECT().FindById(gomock.Any()).Return(expectedUser, nil)

	userByID, err := repo.FindById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, userByID)
}

func TestFindByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	expectedUser := &user.User{ID: 1, Username: "testUser"}
	repo.EXPECT().FindByUsername(gomock.Any()).Return(expectedUser, nil)

	userByUsername, err := repo.FindByUsername("testUser")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, userByUsername)
}

func TestSaveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	userToSave := &user.User{Username: "newUser"}
	repo.EXPECT().Save(gomock.Any()).Return(userToSave, nil)

	savedUser, err := repo.Save(userToSave)
	assert.NoError(t, err)
	assert.Equal(t, userToSave, savedUser)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	updatedUser := &user.User{ID: 1, Username: "updatedUser"}
	repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(updatedUser, nil)

	userToUpdate := &user.User{ID: 1, Username: "oldUser"}
	updated, err := repo.Update(userToUpdate, 1)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, updated)
}

func TestUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	updatedUser := &user.User{ID: 1, Password: "newPassword"}
	repo.EXPECT().UpdatePassword(gomock.Any(), gomock.Any()).Return(updatedUser, nil)

	userToUpdate := &user.User{ID: 1, Password: "oldPassword"}
	updated, err := repo.UpdatePassword(userToUpdate, 1)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, updated)
}
