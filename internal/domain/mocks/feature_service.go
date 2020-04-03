// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qilin/crm-api/internal/domain/service (interfaces: FeatureService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/qilin/crm-api/internal/domain/entity"
	service "github.com/qilin/crm-api/internal/domain/service"
	reflect "reflect"
)

// MockFeatureService is a mock of FeatureService interface
type MockFeatureService struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureServiceMockRecorder
}

// MockFeatureServiceMockRecorder is the mock recorder for MockFeatureService
type MockFeatureServiceMockRecorder struct {
	mock *MockFeatureService
}

// NewMockFeatureService creates a new mock instance
func NewMockFeatureService(ctrl *gomock.Controller) *MockFeatureService {
	mock := &MockFeatureService{ctrl: ctrl}
	mock.recorder = &MockFeatureServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFeatureService) EXPECT() *MockFeatureServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFeatureService) Create(arg0 context.Context, arg1 *service.CreateFeatureData) (*entity.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*entity.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFeatureServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFeatureService)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockFeatureService) Delete(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockFeatureServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeatureService)(nil).Delete), arg0, arg1)
}

// GetByGameID mocks base method
func (m *MockFeatureService) GetByGameID(arg0 context.Context, arg1 uint) ([]entity.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGameID", arg0, arg1)
	ret0, _ := ret[0].([]entity.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGameID indicates an expected call of GetByGameID
func (mr *MockFeatureServiceMockRecorder) GetByGameID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGameID", reflect.TypeOf((*MockFeatureService)(nil).GetByGameID), arg0, arg1)
}

// GetByID mocks base method
func (m *MockFeatureService) GetByID(arg0 context.Context, arg1 uint) (*entity.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockFeatureServiceMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockFeatureService)(nil).GetByID), arg0, arg1)
}

// GetByIDs mocks base method
func (m *MockFeatureService) GetByIDs(arg0 context.Context, arg1 []uint) ([]entity.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", arg0, arg1)
	ret0, _ := ret[0].([]entity.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (mr *MockFeatureServiceMockRecorder) GetByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockFeatureService)(nil).GetByIDs), arg0, arg1)
}

// GetExistByID mocks base method
func (m *MockFeatureService) GetExistByID(arg0 context.Context, arg1 uint) (*entity.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExistByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExistByID indicates an expected call of GetExistByID
func (mr *MockFeatureServiceMockRecorder) GetExistByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExistByID", reflect.TypeOf((*MockFeatureService)(nil).GetExistByID), arg0, arg1)
}

// Update mocks base method
func (m *MockFeatureService) Update(arg0 context.Context, arg1 *service.UpdateFeatureData) (*entity.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*entity.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockFeatureServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFeatureService)(nil).Update), arg0, arg1)
}

// UpdateFeaturesForGame mocks base method
func (m *MockFeatureService) UpdateFeaturesForGame(arg0 context.Context, arg1 *entity.Game, arg2 []uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeaturesForGame", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFeaturesForGame indicates an expected call of UpdateFeaturesForGame
func (mr *MockFeatureServiceMockRecorder) UpdateFeaturesForGame(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeaturesForGame", reflect.TypeOf((*MockFeatureService)(nil).UpdateFeaturesForGame), arg0, arg1, arg2)
}
