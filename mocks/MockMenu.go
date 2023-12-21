package mocks

import (
	menu "capstone/features/menu"
	context "context"
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMenuRepository is a mock of Repository interface.
type MockMenuRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMenuRepositoryMockRecorder
}

// MockMenuRepositoryMockRecorder is the mock recorder for MockMenuRepository.
type MockMenuRepositoryMockRecorder struct {
	mock *MockMenuRepository
}

// NewMockRepository creates a new mock instance.
func NewMockMenuRepository(ctrl *gomock.Controller) *MockMenuRepository {
	mock := &MockMenuRepository{ctrl: ctrl}
	mock.recorder = &MockMenuRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMenuRepository) EXPECT() *MockMenuRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMenuRepository) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMenuRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMenuRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockMenuRepository) FindAll() ([]menu.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]menu.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockMenuRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockMenuRepository)(nil).FindAll))
}

// FindByCategoryId mocks base method.
func (m *MockMenuRepository) FindByCategoryId(arg0 int) ([]menu.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCategoryId", arg0)
	ret0, _ := ret[0].([]menu.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCategoryId indicates an expected call of FindByCategoryId.
func (mr *MockMenuRepositoryMockRecorder) FindByCategoryId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCategoryId", reflect.TypeOf((*MockMenuRepository)(nil).FindByCategoryId), arg0)     
}

// FindById mocks base method.
func (m *MockMenuRepository) FindById(arg0 int) (*menu.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*menu.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockMenuRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockMenuRepository)(nil).FindById), arg0)
}

// FindByName mocks base method.
func (m *MockMenuRepository) FindByName(arg0 string) (*menu.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0)
	ret0, _ := ret[0].(*menu.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockMenuRepositoryMockRecorder) FindByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockMenuRepository)(nil).FindByName), arg0)
}

// Save mocks base method.
func (m *MockMenuRepository) Save(arg0 *menu.Menu) (*menu.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*menu.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockMenuRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMenuRepository)(nil).Save), arg0)
}

// Update mocks base method.
func (m *MockMenuRepository) Update(arg0 *menu.Menu) (*menu.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*menu.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMenuRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMenuRepository)(nil).Update), arg0)
}

// UploadImage mocks base method.
func (m *MockMenuRepository) UploadImage(arg0 context.Context, arg1 multipart.File, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockMenuRepositoryMockRecorder) UploadImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockMenuRepository)(nil).UploadImage), arg0, arg1, arg2)   
}