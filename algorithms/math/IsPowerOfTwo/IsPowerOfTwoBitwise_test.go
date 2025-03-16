// ********RoostGPT********
/*
Test generated by RoostGPT for test go-unit-testing1 using AI Type  and AI Model 

ROOST_METHOD_HASH=isPowerOfTwoBitwise_1ca9da92f9
ROOST_METHOD_SIG_HASH=isPowerOfTwoBitwise_e8104160a5

Scenario 1: Test with Power of Two Number

Details:
Description: This test is meant to check if the function correctly identifies a number that is a power of two.
Execution:
Arrange: No setup is required as we're directly passing the numbers.
Act: Invoke the target function with a number that is a power of two.
Assert: Use Go testing facilities to verify that the function returns true.
Validation:
The function should return true because the given number is a power of two. This test is important as it verifies the basic functionality of the function.

Scenario 2: Test with Non-Power of Two Number

Details:
Description: This test is meant to check if the function correctly identifies a number that is not a power of two.
Execution:
Arrange: No setup is required as we're directly passing the numbers.
Act: Invoke the target function with a number that is not a power of two.
Assert: Use Go testing facilities to verify that the function returns false.
Validation:
The function should return false because the given number is not a power of two. This test is crucial to ensure that the function correctly identifies numbers that are not powers of two.

Scenario 3: Test with Negative Number

Details:
Description: This test is meant to check if the function correctly handles a negative number.
Execution:
Arrange: No setup is required as we're directly passing the numbers.
Act: Invoke the target function with a negative number.
Assert: Use Go testing facilities to verify that the function returns false.
Validation:
The function should return false because negative numbers can't be powers of two. This test is important to ensure that the function handles edge cases correctly.

Scenario 4: Test with Zero

Details:
Description: This test is meant to check if the function correctly handles zero.
Execution:
Arrange: No setup is required as we're directly passing the numbers.
Act: Invoke the target function with zero.
Assert: Use Go testing facilities to verify that the function returns true.
Validation:
The function should return true because zero is considered as a power of two. This test is important to ensure that the function handles edge cases correctly.
*/

// ********RoostGPT********
package IsPowerOfTwo

import (
	"testing"
)

// Renamed the test function to TestIsPowerOfTwo to avoid redeclaration error
func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "Test with Power of Two Number",
			num:  16,
			want: true,
		},
		{
			name: "Test with Non-Power of Two Number",
			num:  15,
			want: false,
		},
		{
			name: "Test with Negative Number",
			num:  -4,
			want: false,
		},
		{
			name: "Test with Zero",
			num:  0,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwoBitwise(tt.num); got != tt.want {
				t.Errorf("isPowerOfTwoBitwise() = %v, want %v", got, tt.want)
			}
		})
	}
}
