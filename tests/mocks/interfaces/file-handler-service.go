// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/application/common/interfaces/file-handler-service.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIFileHandlerService is a mock of IFileHandlerService interface.
type MockIFileHandlerService struct {
	ctrl     *gomock.Controller
	recorder *MockIFileHandlerServiceMockRecorder
}

// MockIFileHandlerServiceMockRecorder is the mock recorder for MockIFileHandlerService.
type MockIFileHandlerServiceMockRecorder struct {
	mock *MockIFileHandlerService
}

// NewMockIFileHandlerService creates a new mock instance.
func NewMockIFileHandlerService(ctrl *gomock.Controller) *MockIFileHandlerService {
	mock := &MockIFileHandlerService{ctrl: ctrl}
	mock.recorder = &MockIFileHandlerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFileHandlerService) EXPECT() *MockIFileHandlerServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIFileHandlerService) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIFileHandlerServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIFileHandlerService)(nil).Close))
}

// LoadFile mocks base method.
func (m *MockIFileHandlerService) LoadFile(filename string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadFile", filename)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFile indicates an expected call of LoadFile.
func (mr *MockIFileHandlerServiceMockRecorder) LoadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFile", reflect.TypeOf((*MockIFileHandlerService)(nil).LoadFile), filename)
}

// ToJson mocks base method.
func (m *MockIFileHandlerService) ToJson() (*any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToJson")
	ret0, _ := ret[0].(*any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToJson indicates an expected call of ToJson.
func (mr *MockIFileHandlerServiceMockRecorder) ToJson() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToJson", reflect.TypeOf((*MockIFileHandlerService)(nil).ToJson))
}
