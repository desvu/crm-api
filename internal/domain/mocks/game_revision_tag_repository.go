// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/repository (interfaces: GameRevisionTagRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	reflect "reflect"
)

// MockGameRevisionTagRepository is a mock of GameRevisionTagRepository interface
type MockGameRevisionTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameRevisionTagRepositoryMockRecorder
}

// MockGameRevisionTagRepositoryMockRecorder is the mock recorder for MockGameRevisionTagRepository
type MockGameRevisionTagRepositoryMockRecorder struct {
	mock *MockGameRevisionTagRepository
}

// NewMockGameRevisionTagRepository creates a new mock instance
func NewMockGameRevisionTagRepository(ctrl *gomock.Controller) *MockGameRevisionTagRepository {
	mock := &MockGameRevisionTagRepository{ctrl: ctrl}
	mock.recorder = &MockGameRevisionTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGameRevisionTagRepository) EXPECT() *MockGameRevisionTagRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGameRevisionTagRepository) Create(arg0 context.Context, arg1 *entity.GameRevisionTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGameRevisionTagRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).Create), arg0, arg1)
}

// CreateMultiple mocks base method
func (m *MockGameRevisionTagRepository) CreateMultiple(arg0 context.Context, arg1 []entity.GameRevisionTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMultiple", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMultiple indicates an expected call of CreateMultiple
func (mr *MockGameRevisionTagRepositoryMockRecorder) CreateMultiple(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultiple", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).CreateMultiple), arg0, arg1)
}

// Delete mocks base method
func (m *MockGameRevisionTagRepository) Delete(arg0 context.Context, arg1 *entity.GameRevisionTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGameRevisionTagRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).Delete), arg0, arg1)
}

// DeleteMultiple mocks base method
func (m *MockGameRevisionTagRepository) DeleteMultiple(arg0 context.Context, arg1 []entity.GameRevisionTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMultiple", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiple indicates an expected call of DeleteMultiple
func (mr *MockGameRevisionTagRepositoryMockRecorder) DeleteMultiple(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiple", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).DeleteMultiple), arg0, arg1)
}

// FindByGameRevisionID mocks base method
func (m *MockGameRevisionTagRepository) FindByGameRevisionID(arg0 context.Context, arg1 uint) ([]entity.GameRevisionTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGameRevisionID", arg0, arg1)
	ret0, _ := ret[0].([]entity.GameRevisionTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGameRevisionID indicates an expected call of FindByGameRevisionID
func (mr *MockGameRevisionTagRepositoryMockRecorder) FindByGameRevisionID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGameRevisionID", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).FindByGameRevisionID), arg0, arg1)
}

// FindByGameRevisionIDAndTagIDs mocks base method
func (m *MockGameRevisionTagRepository) FindByGameRevisionIDAndTagIDs(arg0 context.Context, arg1 uint, arg2 []uint) ([]entity.GameRevisionTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGameRevisionIDAndTagIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entity.GameRevisionTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGameRevisionIDAndTagIDs indicates an expected call of FindByGameRevisionIDAndTagIDs
func (mr *MockGameRevisionTagRepositoryMockRecorder) FindByGameRevisionIDAndTagIDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGameRevisionIDAndTagIDs", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).FindByGameRevisionIDAndTagIDs), arg0, arg1, arg2)
}

// FindByTagID mocks base method
func (m *MockGameRevisionTagRepository) FindByTagID(arg0 context.Context, arg1 uint) ([]entity.GameRevisionTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTagID", arg0, arg1)
	ret0, _ := ret[0].([]entity.GameRevisionTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTagID indicates an expected call of FindByTagID
func (mr *MockGameRevisionTagRepositoryMockRecorder) FindByTagID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTagID", reflect.TypeOf((*MockGameRevisionTagRepository)(nil).FindByTagID), arg0, arg1)
}
