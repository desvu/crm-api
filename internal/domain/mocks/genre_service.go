// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/service (interfaces: GenreService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	service "github.com/qilin/crm-api/internal/domain/service"
	reflect "reflect"
)

// MockGenreService is a mock of GenreService interface
type MockGenreService struct {
	ctrl     *gomock.Controller
	recorder *MockGenreServiceMockRecorder
}

// MockGenreServiceMockRecorder is the mock recorder for MockGenreService
type MockGenreServiceMockRecorder struct {
	mock *MockGenreService
}

// NewMockGenreService creates a new mock instance
func NewMockGenreService(ctrl *gomock.Controller) *MockGenreService {
	mock := &MockGenreService{ctrl: ctrl}
	mock.recorder = &MockGenreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenreService) EXPECT() *MockGenreServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGenreService) Create(arg0 context.Context, arg1 *service.CreateGenreData) (*entity.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entity.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockGenreServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGenreService)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockGenreService) Delete(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGenreServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGenreService)(nil).Delete), arg0, arg1)
}

// GetByGameID mocks base method
func (m *MockGenreService) GetByGameID(arg0 context.Context, arg1 uint) ([]entity.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGameID", arg0, arg1)
	ret0, _ := ret[0].([]entity.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGameID indicates an expected call of GetByGameID
func (mr *MockGenreServiceMockRecorder) GetByGameID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGameID", reflect.TypeOf((*MockGenreService)(nil).GetByGameID), arg0, arg1)
}

// GetByID mocks base method
func (m *MockGenreService) GetByID(arg0 context.Context, arg1 uint) (*entity.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockGenreServiceMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockGenreService)(nil).GetByID), arg0, arg1)
}

// GetByIDs mocks base method
func (m *MockGenreService) GetByIDs(arg0 context.Context, arg1 []uint) ([]entity.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", arg0, arg1)
	ret0, _ := ret[0].([]entity.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockGenreServiceMockRecorder) GetByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockGenreService)(nil).GetByIDs), arg0, arg1)
}

// GetExistByID mocks base method
func (m *MockGenreService) GetExistByID(arg0 context.Context, arg1 uint) (*entity.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExistByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExistByID indicates an expected call of GetExistByID
func (mr *MockGenreServiceMockRecorder) GetExistByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExistByID", reflect.TypeOf((*MockGenreService)(nil).GetExistByID), arg0, arg1)
}

// Update mocks base method
func (m *MockGenreService) Update(arg0 context.Context, arg1 *service.UpdateGenreData) (*entity.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*entity.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockGenreServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGenreService)(nil).Update), arg0, arg1)
}

// UpdateGenreForGame mocks base method
func (m *MockGenreService) UpdateGenreForGame(arg0 context.Context, arg1 *entity.Game, arg2 []uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGenreForGame", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGenreForGame indicates an expected call of UpdateGenreForGame
func (mr *MockGenreServiceMockRecorder) UpdateGenreForGame(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGenreForGame", reflect.TypeOf((*MockGenreService)(nil).UpdateGenreForGame), arg0, arg1, arg2)
}
