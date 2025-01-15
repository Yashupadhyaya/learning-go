package main

import (
	"reflect"
	"testing"
	go_mock "github.com/golang/mock/gomock"
)









/*
ROOST_METHOD_HASH=NewMockMatcher_2a2968746f
ROOST_METHOD_SIG_HASH=NewMockMatcher_f53d5470ec

FUNCTION_DEF=func NewMockMatcher(ctrl *go_mock.Controller) *MockMatcher 

 */
func TestNewMockMatcher(t *testing.T) {
	tests := []struct {
		name      string
		ctrl      *go_mock.Controller
		wantPanic bool
	}{
		{
			name: "Create MockMatcher with valid Controller",
			ctrl: go_mock.NewController(t),
		},
		{
			name:      "Create MockMatcher with nil Controller",
			ctrl:      nil,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("NewMockMatcher() did not panic with nil Controller")
					}
				}()
			}

			got := NewMockMatcher(tt.ctrl)

			if !tt.wantPanic {
				if got == nil {
					t.Errorf("NewMockMatcher() returned nil")
					return
				}

				if got.ctrl != tt.ctrl {
					t.Errorf("NewMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Errorf("NewMockMatcher().recorder is nil")
				}

				if !reflect.TypeOf(got.recorder).AssignableTo(reflect.TypeOf((*MockMatcherMockRecorder)(nil))) {
					t.Errorf("NewMockMatcher().recorder is not of type *MockMatcherMockRecorder")
				}

				if got.recorder.mock != got {
					t.Errorf("NewMockMatcher().recorder.mock does not point back to MockMatcher")
				}
			}
		})
	}

	t.Run("Multiple calls with same Controller", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Errorf("NewMockMatcher() returned the same instance for multiple calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Errorf("NewMockMatcher() returned instances with different Controllers")
		}
	})

	t.Run("Multiple Controller instances", func(t *testing.T) {
		ctrl1 := go_mock.NewController(t)
		ctrl2 := go_mock.NewController(t)

		mock1 := NewMockMatcher(ctrl1)
		mock2 := NewMockMatcher(ctrl2)

		if mock1.ctrl == mock2.ctrl {
			t.Errorf("NewMockMatcher() returned instances with the same Controller for different inputs")
		}
	})
}

