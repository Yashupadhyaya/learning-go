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
		name     string
		ctrl     *go_mock.Controller
		wantMock *MockMatcher
		wantErr  bool
	}{
		{
			name: "Create a new MockMatcher with a valid Controller",
			ctrl: go_mock.NewController(t),
			wantMock: &MockMatcher{
				ctrl:     go_mock.NewController(t),
				recorder: &MockMatcherMockRecorder{},
			},
			wantErr: false,
		},
		{
			name:     "Verify behavior with a nil Controller",
			ctrl:     nil,
			wantMock: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotMock *MockMatcher
			var err interface{}

			func() {
				defer func() {
					err = recover()
				}()
				gotMock = NewMockMatcher(tt.ctrl)
			}()

			if (err != nil) != tt.wantErr {
				t.Errorf("NewMockMatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if gotMock == nil {
					t.Errorf("NewMockMatcher() returned nil, want non-nil")
					return
				}

				if gotMock.ctrl != tt.ctrl {
					t.Errorf("NewMockMatcher().ctrl = %v, want %v", gotMock.ctrl, tt.ctrl)
				}

				if gotMock.recorder == nil {
					t.Errorf("NewMockMatcher().recorder is nil, want non-nil")
				} else if gotMock.recorder.mock != gotMock {
					t.Errorf("NewMockMatcher().recorder.mock = %v, want %v", gotMock.recorder.mock, gotMock)
				}
			}
		})
	}

	t.Run("Verify MockMatcher's recorder initialization", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock := NewMockMatcher(ctrl)

		if mock.recorder == nil {
			t.Errorf("NewMockMatcher().recorder is nil, want non-nil")
		} else if mock.recorder.mock != mock {
			t.Errorf("NewMockMatcher().recorder.mock = %v, want %v", mock.recorder.mock, mock)
		}
	})

	t.Run("Create multiple MockMatcher instances with the same Controller", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock1 := NewMockMatcher(ctrl)
		mock2 := NewMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Errorf("NewMockMatcher() returned the same instance for different calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Errorf("NewMockMatcher() created instances with different Controllers")
		}
	})

	t.Run("Check Controller assignment", func(t *testing.T) {
		ctrl := go_mock.NewController(t)
		mock := NewMockMatcher(ctrl)

		if !reflect.DeepEqual(mock.ctrl, ctrl) {
			t.Errorf("NewMockMatcher().ctrl = %v, want %v", mock.ctrl, ctrl)
		}
	})
}

