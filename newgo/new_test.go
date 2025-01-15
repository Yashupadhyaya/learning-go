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
			name:    "Test with nil Controller",
			ctrl:    nil,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			result = OldMockMatcher(tt.ctrl)

			if (result == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", result == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if result.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher() ctrl = %v, want %v", result.ctrl, tt.ctrl)
				}

				if result.recorder == nil {
					t.Error("OldMockMatcher() recorder is nil")
				} else {
					if _, ok := result.recorder.(*MockMatcherMockRecorder); !ok {
						t.Errorf("OldMockMatcher() recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(result.recorder))
					}
					if result.recorder.(*MockMatcherMockRecorder).mock != result {
						t.Error("OldMockMatcher() recorder.mock does not point back to the MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Multiple calls to OldMockMatcher", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock1 := OldMockMatcher(ctrl)
		mock2 := OldMockMatcher(ctrl)

		if mock1 == mock2 {
			t.Error("OldMockMatcher() returned the same instance for multiple calls")
		}

		if mock1.ctrl != mock2.ctrl {
			t.Error("OldMockMatcher() returned mocks with different controllers")
		}
	})
}

