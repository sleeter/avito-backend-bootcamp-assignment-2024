// Code generated by MockGen. DO NOT EDIT.
// Source: ./flat_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	request "backend-bootcamp-assignment-2024/internal/model/dto/request"
	entity "backend-bootcamp-assignment-2024/internal/model/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockFlatRepository is a mock of FlatRepository interface.
type MockFlatRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFlatRepositoryMockRecorder
}

// MockFlatRepositoryMockRecorder is the mock recorder for MockFlatRepository.
type MockFlatRepositoryMockRecorder struct {
	mock *MockFlatRepository
}

// NewMockFlatRepository creates a new mock instance.
func NewMockFlatRepository(ctrl *gomock.Controller) *MockFlatRepository {
	mock := &MockFlatRepository{ctrl: ctrl}
	mock.recorder = &MockFlatRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlatRepository) EXPECT() *MockFlatRepositoryMockRecorder {
	return m.recorder
}

// CreateFlat mocks base method.
func (m *MockFlatRepository) CreateFlat(ctx context.Context, flat request.CreateFlat) (*entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlat", ctx, flat)
	ret0, _ := ret[0].(*entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlat indicates an expected call of CreateFlat.
func (mr *MockFlatRepositoryMockRecorder) CreateFlat(ctx, flat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlat", reflect.TypeOf((*MockFlatRepository)(nil).CreateFlat), ctx, flat)
}

// GetFlatById mocks base method.
func (m *MockFlatRepository) GetFlatById(ctx context.Context, id int32) (*entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlatById", ctx, id)
	ret0, _ := ret[0].(*entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlatById indicates an expected call of GetFlatById.
func (mr *MockFlatRepositoryMockRecorder) GetFlatById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlatById", reflect.TypeOf((*MockFlatRepository)(nil).GetFlatById), ctx, id)
}

// GetFlatsByHouseId mocks base method.
func (m *MockFlatRepository) GetFlatsByHouseId(ctx context.Context, houseId int32, isModerator bool) ([]entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlatsByHouseId", ctx, houseId, isModerator)
	ret0, _ := ret[0].([]entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlatsByHouseId indicates an expected call of GetFlatsByHouseId.
func (mr *MockFlatRepositoryMockRecorder) GetFlatsByHouseId(ctx, houseId, isModerator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlatsByHouseId", reflect.TypeOf((*MockFlatRepository)(nil).GetFlatsByHouseId), ctx, houseId, isModerator)
}

// UpdateFlatStatus mocks base method.
func (m *MockFlatRepository) UpdateFlatStatus(ctx context.Context, flat request.UpdateFlat) (*entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlatStatus", ctx, flat)
	ret0, _ := ret[0].(*entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFlatStatus indicates an expected call of UpdateFlatStatus.
func (mr *MockFlatRepositoryMockRecorder) UpdateFlatStatus(ctx, flat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlatStatus", reflect.TypeOf((*MockFlatRepository)(nil).UpdateFlatStatus), ctx, flat)
}