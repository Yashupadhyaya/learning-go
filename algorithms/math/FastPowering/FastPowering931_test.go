// ********RoostGPT********
/*
Test generated by RoostGPT for test go-unit-testing1 using AI Type  and AI Model 

ROOST_METHOD_HASH=fastPowering_860ba92ce3
ROOST_METHOD_SIG_HASH=fastPowering_be2e642e07

Scenario 1: Testing power of 0
Details:
    Description: This test is meant to check if the function correctly returns 1 when the power is 0. This is because any number raised to the power of 0 is 1.
Execution:
    Arrange: No arrangement is necessary for this test.
    Act: Invoke the fastPowering function with any base and 0 as the power.
    Assert: Assert that the result of the function is 1.
Validation:
    The choice of assertion is straightforward because we expect the result to always be 1 in this case. This is important to validate the function's compliance with mathematical rules.

Scenario 2: Testing with power as an even number
Details:
    Description: This test is meant to check if the function correctly calculates the power when the power is an even number.
Execution:
    Arrange: No arrangement is necessary for this test.
    Act: Invoke the fastPowering function with any base and an even number as the power.
    Assert: Assert that the result of the function is equal to the base raised to the given power using the math.Pow function.
Validation:
    The choice of assertion is to ensure that the function behaves as expected when the power is an even number. This is important to confirm the function's correctness in handling different power values.

Scenario 3: Testing with power as an odd number
Details:
    Description: This test is meant to check if the function correctly calculates the power when the power is an odd number.
Execution:
    Arrange: No arrangement is necessary for this test.
    Act: Invoke the fastPowering function with any base and an odd number as the power.
    Assert: Assert that the result of the function is equal to the base raised to the given power using the math.Pow function.
Validation:
    The choice of assertion is to ensure that the function behaves as expected when the power is an odd number. This is important to confirm the function's correctness in handling different power values.

Scenario 4: Testing with negative power
Details:
    Description: This test is meant to check if the function correctly calculates the power when the power is a negative number.
Execution:
    Arrange: No arrangement is necessary for this test.
    Act: Invoke the fastPowering function with any base and a negative number as the power.
    Assert: Assert that the result of the function is equal to 1 divided by the base raised to the absolute value of the given power using the math.Pow function.
Validation:
    The choice of assertion is to ensure that the function behaves as expected when the power is a negative number. This test is important to confirm the function's ability to handle all valid input values.
*/

// ********RoostGPT********
package FastPowering

import (
	"math"
	"testing"
)

func TestFastPowering931(t *testing.T) {
	type args struct {
		base  float64
		power int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test power of 0",
			args: args{
				base:  2,
				power: 0,
			},
			want: 1,
		},
		{
			name: "Test with power as an even number",
			args: args{
				base:  2,
				power: 4,
			},
			want: math.Pow(2, 4),
		},
		{
			name: "Test with power as an odd number",
			args: args{
				base:  2,
				power: 3,
			},
			want: math.Pow(2, 3),
		},
		{
			name: "Test with negative power",
			args: args{
				base:  2,
				power: -2,
			},
			want: 1 / math.Pow(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fastPowering(tt.args.base, tt.args.power); got != tt.want {
				t.Errorf("fastPowering() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Success: expected %v, got %v", tt.want, got)
			}
		})
	}
}
