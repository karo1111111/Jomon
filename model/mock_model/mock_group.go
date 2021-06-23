// Code generated by MockGen. DO NOT EDIT.
// Source: model/group.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/traPtitech/Jomon/model"
)

// MockGroupRepository is a mock of GroupRepository interface.
type MockGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepositoryMockRecorder
}

// MockGroupRepositoryMockRecorder is the mock recorder for MockGroupRepository.
type MockGroupRepositoryMockRecorder struct {
	mock *MockGroupRepository
}

// NewMockGroupRepository creates a new mock instance.
func NewMockGroupRepository(ctrl *gomock.Controller) *MockGroupRepository {
	mock := &MockGroupRepository{ctrl: ctrl}
	mock.recorder = &MockGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupRepository) EXPECT() *MockGroupRepositoryMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method.
func (m *MockGroupRepository) CreateGroup(ctx context.Context, name, description string, budget *int, owners *[]model.User) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, name, description, budget, owners)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockGroupRepositoryMockRecorder) CreateGroup(ctx, name, description, budget, owners interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockGroupRepository)(nil).CreateGroup), ctx, name, description, budget, owners)
}

// CreateMember mocks base method.
func (m *MockGroupRepository) CreateMember(ctx context.Context, groupID, userID uuid.UUID) (*model.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", ctx, groupID, userID)
	ret0, _ := ret[0].(*model.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockGroupRepositoryMockRecorder) CreateMember(ctx, groupID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockGroupRepository)(nil).CreateMember), ctx, groupID, userID)
}

// DeleteMember mocks base method.
func (m *MockGroupRepository) DeleteMember(ctx context.Context, groupID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMember", ctx, groupID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMember indicates an expected call of DeleteMember.
func (mr *MockGroupRepositoryMockRecorder) DeleteMember(ctx, groupID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMember", reflect.TypeOf((*MockGroupRepository)(nil).DeleteMember), ctx, groupID, userID)
}

// GetGroups mocks base method.
func (m *MockGroupRepository) GetGroups(ctx context.Context) ([]*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", ctx)
	ret0, _ := ret[0].([]*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroups indicates an expected call of GetGroups.
func (mr *MockGroupRepositoryMockRecorder) GetGroups(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockGroupRepository)(nil).GetGroups), ctx)
}

// GetMembers mocks base method.
func (m *MockGroupRepository) GetMembers(ctx context.Context, groupID uuid.UUID) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers", ctx, groupID)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockGroupRepositoryMockRecorder) GetMembers(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockGroupRepository)(nil).GetMembers), ctx, groupID)
}