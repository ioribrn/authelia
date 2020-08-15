// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handlers/totp.go

// Package handlers is a generated GoMock package.
package handlers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTOTPVerifier is a mock of TOTPVerifier interface
type MockTOTPVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockTOTPVerifierMockRecorder
}

// MockTOTPVerifierMockRecorder is the mock recorder for MockTOTPVerifier
type MockTOTPVerifierMockRecorder struct {
	mock *MockTOTPVerifier
}

// NewMockTOTPVerifier creates a new mock instance
func NewMockTOTPVerifier(ctrl *gomock.Controller) *MockTOTPVerifier {
	mock := &MockTOTPVerifier{ctrl: ctrl}
	mock.recorder = &MockTOTPVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTOTPVerifier) EXPECT() *MockTOTPVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method
func (m *MockTOTPVerifier) Verify(token, secret, algorithm string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", token, secret, algorithm)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockTOTPVerifierMockRecorder) Verify(token, secret, algorithm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTOTPVerifier)(nil).Verify), token, secret, algorithm)
}
