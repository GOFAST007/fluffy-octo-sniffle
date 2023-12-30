// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/addon/package.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockuploader is a mock of uploader interface.
type Mockuploader struct {
	ctrl     *gomock.Controller
	recorder *MockuploaderMockRecorder
}

// MockuploaderMockRecorder is the mock recorder for Mockuploader.
type MockuploaderMockRecorder struct {
	mock *Mockuploader
}

// NewMockuploader creates a new mock instance.
func NewMockuploader(ctrl *gomock.Controller) *Mockuploader {
	mock := &Mockuploader{ctrl: ctrl}
	mock.recorder = &MockuploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockuploader) EXPECT() *MockuploaderMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *Mockuploader) Upload(bucket, key string, data io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", bucket, key, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockuploaderMockRecorder) Upload(bucket, key, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*Mockuploader)(nil).Upload), bucket, key, data)
}