// Code generated by MockGen. DO NOT EDIT.
// Source: minio_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/gabrielopesantos/myDrive-api/internal/models"
	gomock "github.com/golang/mock/gomock"
	v7 "github.com/minio/minio-go/v7"
)

// MockMinioRepository is a mock of MinioRepository interface.
type MockMinioRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMinioRepositoryMockRecorder
}

// MockMinioRepositoryMockRecorder is the mock recorder for MockMinioRepository.
type MockMinioRepositoryMockRecorder struct {
	mock *MockMinioRepository
}

// NewMockMinioRepository creates a new mock instance.
func NewMockMinioRepository(ctrl *gomock.Controller) *MockMinioRepository {
	mock := &MockMinioRepository{ctrl: ctrl}
	mock.recorder = &MockMinioRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinioRepository) EXPECT() *MockMinioRepositoryMockRecorder {
	return m.recorder
}

// PutObject mocks base method.
func (m *MockMinioRepository) PutObject(ctx context.Context, file models.File) (*v7.UploadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", ctx, file)
	ret0, _ := ret[0].(*v7.UploadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockMinioRepositoryMockRecorder) PutObject(ctx, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockMinioRepository)(nil).PutObject), ctx, file)
}
