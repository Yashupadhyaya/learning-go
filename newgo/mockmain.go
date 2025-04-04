package main

import (
	go_mock "github.com/golang/mock/gomock"
)

// MockMatcher is a mock of Matcher interface.
type MockMatcher struct {
	ctrl     *go_mock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher.
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance.
func NewMockMatcher(ctrl *go_mock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// Matches mocks base method.
func (m *MockMatcher) Matches(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Matches", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Matches indicates an expected call of Matches

// String mocks base method.
func (m *MockMatcher) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}
