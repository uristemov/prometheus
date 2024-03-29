// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/murat96k/kitaptar.kz/api"
	entity "github.com/murat96k/kitaptar.kz/internal/user/entity"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ConfirmUser mocks base method.
func (m *MockService) ConfirmUser(ctx context.Context, userID, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmUser", ctx, userID, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfirmUser indicates an expected call of ConfirmUser.
func (mr *MockServiceMockRecorder) ConfirmUser(ctx, userID, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmUser", reflect.TypeOf((*MockService)(nil).ConfirmUser), ctx, userID, code)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(ctx context.Context, u *entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, u)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, u)
}

// DeleteUser mocks base method.
func (m *MockService) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockServiceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockService)(nil).DeleteUser), ctx, id)
}

// GetUserByEmail mocks base method.
func (m *MockService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockServiceMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockService)(nil).GetUserByEmail), ctx, email)
}

// GetUserById mocks base method.
func (m *MockService) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockServiceMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockService)(nil).GetUserById), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockService) UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceMockRecorder) UpdateUser(ctx, id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockService)(nil).UpdateUser), ctx, id, req)
}

// VerifyToken mocks base method.
func (m *MockService) VerifyToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockServiceMockRecorder) VerifyToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockService)(nil).VerifyToken), token)
}
