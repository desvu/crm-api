// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/repository (interfaces: GameRevisionPublisherRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	reflect "reflect"
)

// MockGameRevisionPublisherRepository is a mock of GameRevisionPublisherRepository interface
type MockGameRevisionPublisherRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameRevisionPublisherRepositoryMockRecorder
}

// MockGameRevisionPublisherRepositoryMockRecorder is the mock recorder for MockGameRevisionPublisherRepository
type MockGameRevisionPublisherRepositoryMockRecorder struct {
	mock *MockGameRevisionPublisherRepository
}

// NewMockGameRevisionPublisherRepository creates a new mock instance
func NewMockGameRevisionPublisherRepository(ctrl *gomock.Controller) *MockGameRevisionPublisherRepository {
	mock := &MockGameRevisionPublisherRepository{ctrl: ctrl}
	mock.recorder = &MockGameRevisionPublisherRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGameRevisionPublisherRepository) EXPECT() *MockGameRevisionPublisherRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGameRevisionPublisherRepository) Create(arg0 context.Context, arg1 *entity.GameRevisionPublisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).Create), arg0, arg1)
}

// CreateMultiple mocks base method
func (m *MockGameRevisionPublisherRepository) CreateMultiple(arg0 context.Context, arg1 []entity.GameRevisionPublisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMultiple", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMultiple indicates an expected call of CreateMultiple
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) CreateMultiple(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultiple", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).CreateMultiple), arg0, arg1)
}

// Delete mocks base method
func (m *MockGameRevisionPublisherRepository) Delete(arg0 context.Context, arg1 *entity.GameRevisionPublisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).Delete), arg0, arg1)
}

// DeleteMultiple mocks base method
func (m *MockGameRevisionPublisherRepository) DeleteMultiple(arg0 context.Context, arg1 []entity.GameRevisionPublisher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMultiple", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiple indicates an expected call of DeleteMultiple
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) DeleteMultiple(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiple", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).DeleteMultiple), arg0, arg1)
}

// FindByGameRevisionID mocks base method
func (m *MockGameRevisionPublisherRepository) FindByGameRevisionID(arg0 context.Context, arg1 uint) ([]entity.GameRevisionPublisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGameRevisionID", arg0, arg1)
	ret0, _ := ret[0].([]entity.GameRevisionPublisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGameRevisionID indicates an expected call of FindByGameRevisionID
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) FindByGameRevisionID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGameRevisionID", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).FindByGameRevisionID), arg0, arg1)
}

// FindByGameRevisionIDAndPublisherIDs mocks base method
func (m *MockGameRevisionPublisherRepository) FindByGameRevisionIDAndPublisherIDs(arg0 context.Context, arg1 uint, arg2 []uint) ([]entity.GameRevisionPublisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGameRevisionIDAndPublisherIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entity.GameRevisionPublisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGameRevisionIDAndPublisherIDs indicates an expected call of FindByGameRevisionIDAndPublisherIDs
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) FindByGameRevisionIDAndPublisherIDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGameRevisionIDAndPublisherIDs", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).FindByGameRevisionIDAndPublisherIDs), arg0, arg1, arg2)
}

// FindByGameRevisionIDs mocks base method
func (m *MockGameRevisionPublisherRepository) FindByGameRevisionIDs(arg0 context.Context, arg1 []uint) ([]entity.GameRevisionPublisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGameRevisionIDs", arg0, arg1)
	ret0, _ := ret[0].([]entity.GameRevisionPublisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGameRevisionIDs indicates an expected call of FindByGameRevisionIDs
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) FindByGameRevisionIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGameRevisionIDs", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).FindByGameRevisionIDs), arg0, arg1)
}

// FindByPublisherID mocks base method
func (m *MockGameRevisionPublisherRepository) FindByPublisherID(arg0 context.Context, arg1 uint) ([]entity.GameRevisionPublisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPublisherID", arg0, arg1)
	ret0, _ := ret[0].([]entity.GameRevisionPublisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPublisherID indicates an expected call of FindByPublisherID
func (mr *MockGameRevisionPublisherRepositoryMockRecorder) FindByPublisherID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPublisherID", reflect.TypeOf((*MockGameRevisionPublisherRepository)(nil).FindByPublisherID), arg0, arg1)
}