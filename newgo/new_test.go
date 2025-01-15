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
			name:      "Create MockMatcher with valid Controller",
			ctrl:      gomock.NewController(t),
			wantNil:   false,
			wantPanic: false,
		},
		{
			name:      "Create MockMatcher with nil Controller",
			ctrl:      nil,
			wantNil:   true,
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()

			result = OldMockMatcher(tt.ctrl)

			if (result == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", result == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if result.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", result.ctrl, tt.ctrl)
				}

				if result.recorder == nil {
					t.Error("OldMockMatcher().recorder is nil")
				} else {
					if _, ok := result.recorder.(*MockMatcherMockRecorder); !ok {
						t.Errorf("OldMockMatcher().recorder is not of type *MockMatcherMockRecorder")
					}
					if result.recorder.(*MockMatcherMockRecorder).mock != result {
						t.Errorf("OldMockMatcher().recorder.mock does not point back to MockMatcher")
					}
				}
			}
		})
	}

	t.Run("Multiple calls return unique instances", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		result1 := OldMockMatcher(ctrl)
		result2 := OldMockMatcher(ctrl)

		if reflect.DeepEqual(result1, result2) {
			t.Errorf("Multiple calls to OldMockMatcher() returned the same instance")
		}
	})
}

