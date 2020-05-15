// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/repository (interfaces: TagRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	reflect "reflect"
)

// MockTagRepository is a mock of TagRepository interface
type MockTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTagRepositoryMockRecorder
}

// MockTagRepositoryMockRecorder is the mock recorder for MockTagRepository
type MockTagRepositoryMockRecorder struct {
	mock *MockTagRepository
}

// NewMockTagRepository creates a new mock instance
func NewMockTagRepository(ctrl *gomock.Controller) *MockTagRepository {
	mock := &MockTagRepository{ctrl: ctrl}
	mock.recorder = &MockTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagRepository) EXPECT() *MockTagRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTagRepository) Create(arg0 context.Context, arg1 *entity.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTagRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTagRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockTagRepository) Delete(arg0 context.Context, arg1 *entity.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTagRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTagRepository)(nil).Delete), arg0, arg1)
}

// FindAll mocks base method
func (m *MockTagRepository) FindAll(arg0 context.Context) ([]entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockTagRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTagRepository)(nil).FindAll), arg0)
}

// FindByID mocks base method
func (m *MockTagRepository) FindByID(arg0 context.Context, arg1 uint) (*entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockTagRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockTagRepository)(nil).FindByID), arg0, arg1)
}

// FindByIDs mocks base method
func (m *MockTagRepository) FindByIDs(arg0 context.Context, arg1 []uint) ([]entity.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", arg0, arg1)
	ret0, _ := ret[0].([]entity.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs
func (mr *MockTagRepositoryMockRecorder) FindByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockTagRepository)(nil).FindByIDs), arg0, arg1)
}

// Update mocks base method
func (m *MockTagRepository) Update(arg0 context.Context, arg1 *entity.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTagRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTagRepository)(nil).Update), arg0, arg1)
}
