// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/repository (interfaces: PublisherRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	repository "github.com/qilin/crm-api/internal/domain/repository"
	reflect "reflect"
)

// MockPublisherRepository is a mock of PublisherRepository interface
type MockPublisherRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherRepositoryMockRecorder
}

// MockPublisherRepositoryMockRecorder is the mock recorder for MockPublisherRepository
type MockPublisherRepositoryMockRecorder struct {
	mock *MockPublisherRepository
}

// NewMockPublisherRepository creates a new mock instance
func NewMockPublisherRepository(ctrl *gomock.Controller) *MockPublisherRepository {
	mock := &MockPublisherRepository{ctrl: ctrl}
	mock.recorder = &MockPublisherRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisherRepository) EXPECT() *MockPublisherRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPublisherRepository) Create(arg0 context.Context, arg1 *entity.Publisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockPublisherRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPublisherRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockPublisherRepository) Delete(arg0 context.Context, arg1 *entity.Publisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPublisherRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPublisherRepository)(nil).Delete), arg0, arg1)
}

// FindByFilter mocks base method
func (m *MockPublisherRepository) FindByFilter(arg0 context.Context, arg1 *repository.FindByFilterPublisherData) ([]entity.Publisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByFilter", arg0, arg1)
	ret0, _ := ret[0].([]entity.Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByFilter indicates an expected call of FindByFilter
func (mr *MockPublisherRepositoryMockRecorder) FindByFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByFilter", reflect.TypeOf((*MockPublisherRepository)(nil).FindByFilter), arg0, arg1)
}

// FindByID mocks base method
func (m *MockPublisherRepository) FindByID(arg0 context.Context, arg1 uint) (*entity.Publisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockPublisherRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPublisherRepository)(nil).FindByID), arg0, arg1)
}

// FindByIDs mocks base method
func (m *MockPublisherRepository) FindByIDs(arg0 context.Context, arg1 []uint) ([]entity.Publisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDs", arg0, arg1)
	ret0, _ := ret[0].([]entity.Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs
func (mr *MockPublisherRepositoryMockRecorder) FindByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockPublisherRepository)(nil).FindByIDs), arg0, arg1)
}

// Update mocks base method
func (m *MockPublisherRepository) Update(arg0 context.Context, arg1 *entity.Publisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPublisherRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPublisherRepository)(nil).Update), arg0, arg1)
}
