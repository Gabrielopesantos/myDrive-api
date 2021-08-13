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

// GetObject mocks base method.
func (m *MockMinioRepository) GetObject(ctx context.Context, bucket, fileName string) (*v7.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, bucket, fileName)
	ret0, _ := ret[0].(*v7.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockMinioRepositoryMockRecorder) GetObject(ctx, bucket, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockMinioRepository)(nil).GetObject), ctx, bucket, fileName)
}

// PutObject mocks base method.
func (m *MockMinioRepository) PutObject(ctx context.Context, input models.UploadInput) (*v7.UploadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", ctx, input)
	ret0, _ := ret[0].(*v7.UploadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockMinioRepositoryMockRecorder) PutObject(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockMinioRepository)(nil).PutObject), ctx, input)
}

// RemoveObject mocks base method.
func (m *MockMinioRepository) RemoveObject(ctx context.Context, bucket, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveObject", ctx, bucket, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveObject indicates an expected call of RemoveObject.
func (mr *MockMinioRepositoryMockRecorder) RemoveObject(ctx, bucket, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObject", reflect.TypeOf((*MockMinioRepository)(nil).RemoveObject), ctx, bucket, fileName)
}
