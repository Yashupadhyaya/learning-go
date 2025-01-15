package main

import (
	"reflect"
	"testing"
	"github.com/golang/mock/gomock"
)









/*
ROOST_METHOD_HASH=OldMockMatcher_b241463c7a
ROOST_METHOD_SIG_HASH=OldMockMatcher_4087741302

FUNCTION_DEF=func OldMockMatcher(ctrl *gomock.Controller) *MockMatcher 

 */
func TestOldMockMatcher(t *testing.T) {
	tests := []struct {
		name      string
		ctrl      *gomock.Controller
		wantNil   bool
		wantPanic bool
	}{
		{
			name:    "Create MockMatcher with valid Controller",
			ctrl:    gomock.NewController(t),
			wantNil: false,
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
						t.Errorf("OldMockMatcher() did not panic with nil Controller")
					}
				}()
			}

			got := OldMockMatcher(tt.ctrl)

			if tt.wantNil && got != nil {
				t.Errorf("OldMockMatcher() = %v, want nil", got)
			}

			if !tt.wantNil {
				if got == nil {
					t.Errorf("OldMockMatcher() returned nil, want non-nil")
				}

				if got.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Errorf("OldMockMatcher().recorder is nil")
				}

				if reflect.TypeOf(got.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
					t.Errorf("OldMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(got.recorder))
				}

				if got.recorder.mock != got {
					t.Errorf("OldMockMatcher().recorder.mock does not point back to MockMatcher")
				}
			}
		})
	}

	t.Run("Multiple calls return unique instances", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Errorf("Multiple calls to OldMockMatcher() returned the same instance")
		}
	})
}

