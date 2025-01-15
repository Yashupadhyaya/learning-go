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
			var mock *MockMatcher
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("OldMockMatcher() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()

			mock = OldMockMatcher(tt.ctrl)

			if (mock == nil) != tt.wantNil {
				t.Errorf("OldMockMatcher() returned nil: %v, want nil: %v", mock == nil, tt.wantNil)
			}

			if !tt.wantNil {
				if mock.ctrl != tt.ctrl {
					t.Errorf("OldMockMatcher().ctrl = %v, want %v", mock.ctrl, tt.ctrl)
				}

				if mock.recorder == nil {
					t.Error("OldMockMatcher().recorder is nil")
				} else {
					if reflect.TypeOf(mock.recorder) != reflect.TypeOf(&MockMatcherMockRecorder{}) {
						t.Errorf("OldMockMatcher().recorder type = %v, want *MockMatcherMockRecorder", reflect.TypeOf(mock.recorder))
					}

					if mock.recorder.mock != mock {
						t.Error("OldMockMatcher().recorder.mock does not point back to the MockMatcher")
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
}

