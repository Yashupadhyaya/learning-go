// ********RoostGPT********
/*
Test generated by RoostGPT for test go-unit-testing1 using AI Type  and AI Model 

ROOST_METHOD_HASH=isPowerOfTwo_fe7a80abf8
ROOST_METHOD_SIG_HASH=isPowerOfTwo_a909b954a6

Scenario 1: Testing with a number that is a power of two

Details:
  Description: This test is meant to check if the function is able to correctly determine if a number is a power of two. The target scenario is when the input number is a power of two (e.g., 2, 4, 8, 16, etc.).

Execution:
  Arrange: No setup is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a number that is a power of two.
  Assert: Use Go's testing facilities to verify that the function returns true.

Validation:
  The assertion checks if the function returns true, which is the expected outcome for numbers that are a power of two. This test is important to ensure that the function correctly identifies numbers that are a power of two.

Scenario 2: Testing with a number that is not a power of two

Details:
  Description: This test is meant to check if the function can correctly identify numbers that are not a power of two. The target scenario is when the input number is not a power of two (e.g., 3, 5, 6, 7, etc.).

Execution:
  Arrange: No setup is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a number that is not a power of two.
  Assert: Use Go's testing facilities to verify that the function returns false.

Validation:
  The assertion checks if the function returns false, which is the expected outcome for numbers that are not a power of two. This test is important to verify that the function correctly identifies numbers that are not a power of two.

Scenario 3: Testing with a negative number

Details:
  Description: This test is meant to check if the function can correctly handle negative numbers. The target scenario is when the input number is negative.

Execution:
  Arrange: No setup is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a negative number.
  Assert: Use Go's testing facilities to verify that the function returns false.

Validation:
  The assertion checks if the function returns false, which is the expected outcome for negative numbers. This test is important to ensure that the function correctly handles negative numbers, which are not a power of two.

Scenario 4: Testing with zero

Details:
  Description: This test is meant to check if the function can correctly handle zero. The target scenario is when the input number is zero.

Execution:
  Arrange: No setup is necessary for this test.
  Act: Invoke the isPowerOfTwo function with zero.
  Assert: Use Go's testing facilities to verify that the function returns false.

Validation:
  The assertion checks if the function returns false, which is the expected outcome for zero. This test is important to ensure that the function correctly handles zero, which is not a power of two.
*/

// ********RoostGPT********
package IsPowerOfTwo

import (
	"testing"
)

func TestIsPowerOfTwo92(t *testing.T) {
	// define the table driven tests
	tests := []struct {
		name   string
		input  int
		expect bool
	}{
		{name: "Testing with a number that is a power of two", input: 16, expect: true},
		{name: "Testing with a number that is not a power of two", input: 7, expect: false},
		{name: "Testing with a negative number", input: -2, expect: false},
		{name: "Testing with zero", input: 0, expect: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPowerOfTwo(tt.input)

			// log the test case details
			t.Logf("Running test case: %s with input: %d", tt.name, tt.input)

			if result != tt.expect {
				t.Errorf("Expected output %t but got %t", tt.expect, result)
			} else {
				t.Logf("Success! Expected output %t and got %t", tt.expect, result)
			}
		})
	}
}
