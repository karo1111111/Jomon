// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/traPtitech/Jomon/model"
)

// MockCommentRepository is a mock of CommentRepository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockCommentRepository) CreateComment(ctx context.Context, comment string, requestID, userID uuid.UUID) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, comment, requestID, userID)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentRepositoryMockRecorder) CreateComment(ctx, comment, requestID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentRepository)(nil).CreateComment), ctx, comment, requestID, userID)
}

// DeleteComment mocks base method.
func (m *MockCommentRepository) DeleteComment(ctx context.Context, requestID, commentID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, requestID, commentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockCommentRepositoryMockRecorder) DeleteComment(ctx, requestID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentRepository)(nil).DeleteComment), ctx, requestID, commentID)
}

// UpdateComment mocks base method.
func (m *MockCommentRepository) UpdateComment(ctx context.Context, comment string, requestID, commentID uuid.UUID) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", ctx, comment, requestID, commentID)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockCommentRepositoryMockRecorder) UpdateComment(ctx, comment, requestID, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCommentRepository)(nil).UpdateComment), ctx, comment, requestID, commentID)
}
