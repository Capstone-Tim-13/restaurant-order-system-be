package mocks

import (
	feedback "capstone/features/feedback"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeedbackRepository is a mock of Repository interface.
type MockFeedbackRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFeedbackRepositoryMockRecorder
}

// MockFeedbackRepositoryMockRecorder is the mock recorder for MockFeedbackRepository.
type MockFeedbackRepositoryMockRecorder struct {
	mock *MockFeedbackRepository
}

// NewMockRepository creates a new mock instance.
func NewMockFeedbackRepository(ctrl *gomock.Controller) *MockFeedbackRepository {
	mock := &MockFeedbackRepository{ctrl: ctrl}
	mock.recorder = &MockFeedbackRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedbackRepository) EXPECT() *MockFeedbackRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockFeedbackRepository) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFeedbackRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeedbackRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockFeedbackRepository) FindAll() ([]feedback.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]feedback.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockFeedbackRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockFeedbackRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockFeedbackRepository) FindById(arg0 int) (*feedback.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*feedback.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockFeedbackRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockFeedbackRepository)(nil).FindById), arg0)
}

// Save mocks base method.
func (m *MockFeedbackRepository) Save(arg0 *feedback.Feedback) (*feedback.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*feedback.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockFeedbackRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFeedbackRepository)(nil).Save), arg0)
}

// Update mocks base method.
func (m *MockFeedbackRepository) Update(arg0 *feedback.Feedback, arg1 int) (*feedback.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*feedback.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockFeedbackRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFeedbackRepository)(nil).Update), arg0, arg1)
}