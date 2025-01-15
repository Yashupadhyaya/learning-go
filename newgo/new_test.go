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
			var got *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()

			got = OldMockMatcher(tt.ctrl)

			if (got == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", got == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if got.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", got.ctrl, tt.ctrl)
				}

				if got.recorder == nil {
					t.Error("OldMockMatcher().recorder is nil")
				} else {
					if _, ok := got.recorder.(*MockMatcherMockRecorder); !ok {
						t.Errorf("OldMockMatcher().recorder is not of type *MockMatcherMockRecorder")
					}
					if got.recorder.(*MockMatcherMockRecorder).mock != got {
						t.Errorf("OldMockMatcher().recorder.mock does not point back to MockMatcher")
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
			t.Error("Multiple calls to OldMockMatcher returned the same instance")
		}
	})

	t.Run("Verify MockMatcher's recorder initialization", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := OldMockMatcher(ctrl)

		if mock.recorder == nil {
			t.Error("MockMatcher's recorder is nil")
		} else {
			if _, ok := mock.recorder.(*MockMatcherMockRecorder); !ok {
				t.Errorf("MockMatcher's recorder is not of type *MockMatcherMockRecorder")
			}
		}
	})

	t.Run("Verify MockMatcher's ctrl field assignment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := OldMockMatcher(ctrl)

		if !reflect.DeepEqual(mock.ctrl, ctrl) {
			t.Errorf("MockMatcher's ctrl field is not correctly assigned")
		}
	})

	t.Run("Verify recorder's mock field assignment", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock := OldMockMatcher(ctrl)

		if mock.recorder.(*MockMatcherMockRecorder).mock != mock {
			t.Errorf("Recorder's mock field does not point back to the MockMatcher")
		}
	})
}

