// Code generated by MockGen. DO NOT EDIT.
// Source: tag.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/traPtitech/Jomon/model"
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

// GetTags mocks base method
func (m *MockTagRepository) GetTags(ctx context.Context) ([]*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", ctx)
	ret0, _ := ret[0].([]*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags
func (mr *MockTagRepositoryMockRecorder) GetTags(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockTagRepository)(nil).GetTags), ctx)
}

// GetTag mocks base method
func (m *MockTagRepository) GetTag(ctx context.Context, tagID uuid.UUID) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", ctx, tagID)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag
func (mr *MockTagRepositoryMockRecorder) GetTag(ctx, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockTagRepository)(nil).GetTag), ctx, tagID)
}

// CreateTag mocks base method
func (m *MockTagRepository) CreateTag(ctx context.Context, name, description string) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", ctx, name, description)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag
func (mr *MockTagRepositoryMockRecorder) CreateTag(ctx, name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockTagRepository)(nil).CreateTag), ctx, name, description)
}

// UpdateTag mocks base method
func (m *MockTagRepository) UpdateTag(ctx context.Context, tagID uuid.UUID, name, description string) (*model.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTag", ctx, tagID, name, description)
	ret0, _ := ret[0].(*model.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTag indicates an expected call of UpdateTag
func (mr *MockTagRepositoryMockRecorder) UpdateTag(ctx, tagID, name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTag", reflect.TypeOf((*MockTagRepository)(nil).UpdateTag), ctx, tagID, name, description)
}

// DeleteTag mocks base method
func (m *MockTagRepository) DeleteTag(ctx context.Context, tagID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", ctx, tagID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag
func (mr *MockTagRepositoryMockRecorder) DeleteTag(ctx, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockTagRepository)(nil).DeleteTag), ctx, tagID)
}
