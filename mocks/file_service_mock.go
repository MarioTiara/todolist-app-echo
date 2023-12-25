// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/services/file_service.go

// Package mock_services is a generated GoMock package.
package mocks

import (
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dtos "github.com/marioTiara/todolistapp/internal/api/dtos"
	models "github.com/marioTiara/todolistapp/internal/api/models"
)

// MockFileService is a mock of FileService interface.
type MockFileService struct {
	ctrl     *gomock.Controller
	recorder *MockFileServiceMockRecorder
}

// MockFileServiceMockRecorder is the mock recorder for MockFileService.
type MockFileServiceMockRecorder struct {
	mock *MockFileService
}

// NewMockFileService creates a new mock instance.
func NewMockFileService(ctrl *gomock.Controller) *MockFileService {
	mock := &MockFileService{ctrl: ctrl}
	mock.recorder = &MockFileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileService) EXPECT() *MockFileServiceMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockFileService) DeleteByID(fileID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockFileServiceMockRecorder) DeleteByID(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockFileService)(nil).DeleteByID), fileID)
}

// DeleteByTaskID mocks base method.
func (m *MockFileService) DeleteByTaskID(taskID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByTaskID", taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByTaskID indicates an expected call of DeleteByTaskID.
func (mr *MockFileServiceMockRecorder) DeleteByTaskID(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByTaskID", reflect.TypeOf((*MockFileService)(nil).DeleteByTaskID), taskID)
}

// Download mocks base method.
func (m *MockFileService) Download(fileName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", fileName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockFileServiceMockRecorder) Download(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockFileService)(nil).Download), fileName)
}

// GetByID mocks base method.
func (m *MockFileService) GetByID(fileID uint) (models.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", fileID)
	ret0, _ := ret[0].(models.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockFileServiceMockRecorder) GetByID(fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockFileService)(nil).GetByID), fileID)
}

// GetByTaskID mocks base method.
func (m *MockFileService) GetByTaskID(taskID uint) ([]models.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTaskID", taskID)
	ret0, _ := ret[0].([]models.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTaskID indicates an expected call of GetByTaskID.
func (mr *MockFileServiceMockRecorder) GetByTaskID(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTaskID", reflect.TypeOf((*MockFileService)(nil).GetByTaskID), taskID)
}

// SaveFile mocks base method.
func (m *MockFileService) SaveFile(taskID uint, file *multipart.FileHeader) (dtos.FileQueryModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", taskID, file)
	ret0, _ := ret[0].(dtos.FileQueryModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileServiceMockRecorder) SaveFile(taskID, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileService)(nil).SaveFile), taskID, file)
}
