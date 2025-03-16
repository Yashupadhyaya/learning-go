// ********RoostGPT********
/*
Test generated by RoostGPT for test go-unit-testing1 using AI Type  and AI Model 

ROOST_METHOD_HASH=degreeToRadian_19ddfa9792
ROOST_METHOD_SIG_HASH=degreeToRadian_c38219c312

Scenario 1: Testing with Positive Degree Input

  Details:
    Description: This test is meant to check if the degreeToRadian function correctly converts a positive degree value to its radian equivalent.
  Execution:
    Arrange: No setup required for this test.
    Act: Invoke the degreeToRadian function with a positive degree value.
    Assert: Use Go testing facilities to verify that the actual result matches the expected radian value.
  Validation:
    The assertion checks if the calculated radian value matches the expected radian value for the given degree. The importance of this test is to verify the basic functionality of the degreeToRadian function.

Scenario 2: Testing with Negative Degree Input

  Details:
    Description: This test is meant to check if the degreeToRadian function correctly converts a negative degree value to its radian equivalent.
  Execution:
    Arrange: No setup required for this test.
    Act: Invoke the degreeToRadian function with a negative degree value.
    Assert: Use Go testing facilities to verify that the actual result matches the expected radian value.
  Validation:
    The assertion checks if the calculated radian value matches the expected radian value for the given degree. This test is important to verify that the degreeToRadian function can handle negative degree values.

Scenario 3: Testing with Zero Degree Input

  Details:
    Description: This test is meant to check if the degreeToRadian function correctly handles a zero degree input.
  Execution:
    Arrange: No setup required for this test.
    Act: Invoke the degreeToRadian function with zero degree.
    Assert: Use Go testing facilities to verify that the actual result is zero.
  Validation:
    The assertion checks if the calculated radian value is zero for zero degree. This test is important to verify that the degreeToRadian function can handle zero degree values.

Scenario 4: Testing with Large Degree Input

  Details:
    Description: This test is meant to check if the degreeToRadian function correctly handles a large degree value input.
  Execution:
    Arrange: No setup required for this test.
    Act: Invoke the degreeToRadian function with a large degree value.
    Assert: Use Go testing facilities to verify that the actual result matches the expected radian value.
  Validation:
    The assertion checks if the calculated radian value matches the expected radian value for the given large degree. This test is important to verify that the degreeToRadian function can handle large degree values. 

Scenario 5: Testing with Floating Point Degree Input

  Details:
    Description: This test is meant to check if the degreeToRadian function correctly handles a floating point degree value input.
  Execution:
    Arrange: No setup required for this test.
    Act: Invoke the degreeToRadian function with a floating point degree value.
    Assert: Use Go testing facilities to verify that the actual result matches the expected radian value.
  Validation:
    The assertion checks if the calculated radian value matches the expected radian value for the given floating point degree. This test is important to verify that the degreeToRadian function can handle floating point degree values.
*/

// ********RoostGPT********
package Radian

import (
	"math"
	"testing"
)

func TestDegreeToRadianConversion(t *testing.T) {
	type test struct {
		degree float64
		want   float64
	}

	tests := []test{
		{degree: 90, want: math.Pi / 2},
		{degree: -90, want: -math.Pi / 2},
		{degree: 0, want: 0},
		{degree: 360, want: 2 * math.Pi},
		{degree: 45.5, want: 45.5 * math.Pi / 180},
	}

	for _, tc := range tests {
		got := degreeToRadian(tc.degree)
		if got != tc.want {
			t.Errorf("degreeToRadian(%v) = %v; want %v", tc.degree, got, tc.want)
		} else {
			t.Logf("Success: degreeToRadian(%v) = %v", tc.degree, got)
		}
	}
}
