// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/service (interfaces: TagService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	service "github.com/qilin/crm-api/internal/domain/service"
	reflect "reflect"
)

// MockTagService is a mock of TagService interface
type MockTagService struct {
	ctrl     *gomock.Controller
	recorder *MockTagServiceMockRecorder
}

// MockTagServiceMockRecorder is the mock recorder for MockTagService
type MockTagServiceMockRecorder struct {
	mock *MockTagService
}

// NewMockTagService creates a new mock instance
func NewMockTagService(ctrl *gomock.Controller) *MockTagService {
	mock := &MockTagService{ctrl: ctrl}
	mock.recorder = &MockTagServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagService) EXPECT() *MockTagServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTagService) Create(arg0 context.Context, arg1 *service.CreateTagData) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTagServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTagService)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockTagService) Delete(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTagServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTagService)(nil).Delete), arg0, arg1)
}

// GetByGameRevisionID mocks base method
func (m *MockTagService) GetByGameRevisionID(arg0 context.Context, arg1 uint) ([]entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGameRevisionID", arg0, arg1)
	ret0, _ := ret[0].([]entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGameRevisionID indicates an expected call of GetByGameRevisionID
func (mr *MockTagServiceMockRecorder) GetByGameRevisionID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGameRevisionID", reflect.TypeOf((*MockTagService)(nil).GetByGameRevisionID), arg0, arg1)
}

// GetByID mocks base method
func (m *MockTagService) GetByID(arg0 context.Context, arg1 uint) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTagServiceMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTagService)(nil).GetByID), arg0, arg1)
}

// GetByIDs mocks base method
func (m *MockTagService) GetByIDs(arg0 context.Context, arg1 []uint) ([]entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", arg0, arg1)
	ret0, _ := ret[0].([]entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockTagServiceMockRecorder) GetByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockTagService)(nil).GetByIDs), arg0, arg1)
}

// GetExistByID mocks base method
func (m *MockTagService) GetExistByID(arg0 context.Context, arg1 uint) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExistByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExistByID indicates an expected call of GetExistByID
func (mr *MockTagServiceMockRecorder) GetExistByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExistByID", reflect.TypeOf((*MockTagService)(nil).GetExistByID), arg0, arg1)
}

// Update mocks base method
func (m *MockTagService) Update(arg0 context.Context, arg1 *service.UpdateTagData) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockTagServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTagService)(nil).Update), arg0, arg1)
}

// UpdateTagsForGameRevision mocks base method
func (m *MockTagService) UpdateTagsForGameRevision(arg0 context.Context, arg1 *entity.GameRevision, arg2 []uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTagsForGameRevision", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTagsForGameRevision indicates an expected call of UpdateTagsForGameRevision
func (mr *MockTagServiceMockRecorder) UpdateTagsForGameRevision(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTagsForGameRevision", reflect.TypeOf((*MockTagService)(nil).UpdateTagsForGameRevision), arg0, arg1, arg2)
}
