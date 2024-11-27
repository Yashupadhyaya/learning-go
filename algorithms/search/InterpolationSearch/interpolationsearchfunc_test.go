// ********RoostGPT********
/*
Test generated by RoostGPT for test go-single-file-test using AI Type Open AI and AI Model gpt-4

ROOST_METHOD_HASH=interpolationSearchFunc_aed10f9c49
ROOST_METHOD_SIG_HASH=interpolationSearchFunc_bd78985e43

Scenario 1: Valid Query in the Middle of Array

Details:
    Description: This test is meant to check if the function correctly identifies the index of a query located in the middle of the array.
Execution:
    Arrange: Initialize an array of integers in ascending order and a query that exists in that array.
    Act: Invoke the interpolationSearchFunc function with the array and the query as parameters.
    Assert: Use Go's testing facilities to assert that the returned index matches the known index of the query in the array.
Validation:
    The choice of assertion is based on the expected behavior of the function: if the query exists in the array, the function should return its index. This test is important as it checks the basic functionality of the function.

Scenario 2: Query at the Beginning of Array

Details:
    Description: This test is meant to check if the function correctly identifies the index of a query located at the beginning of the array.
Execution:
    Arrange: Initialize an array of integers in ascending order and a query that is the first element of that array.
    Act: Invoke the interpolationSearchFunc function with the array and the query as parameters.
    Assert: Use Go's testing facilities to assert that the returned index is 0.
Validation:
    The choice of assertion is based on the expected behavior of the function: if the query is the first element of the array, the function should return 0. This test is important as it checks the edge case of the function.

Scenario 3: Query at the End of Array

Details:
    Description: This test is meant to check if the function correctly identifies the index of a query located at the end of the array.
Execution:
    Arrange: Initialize an array of integers in ascending order and a query that is the last element of that array.
    Act: Invoke the interpolationSearchFunc function with the array and the query as parameters.
    Assert: Use Go's testing facilities to assert that the returned index is the length of the array minus 1.
Validation:
    The choice of assertion is based on the expected behavior of the function: if the query is the last element of the array, the function should return the index of the last element. This test is important as it checks another edge case of the function.

Scenario 4: Query Not in Array

Details:
    Description: This test is meant to check if the function correctly returns -1 when the query is not in the array.
Execution:
    Arrange: Initialize an array of integers in ascending order and a query that does not exist in that array.
    Act: Invoke the interpolationSearchFunc function with the array and the query as parameters.
    Assert: Use Go's testing facilities to assert that the returned index is -1.
Validation:
    The choice of assertion is based on the expected behavior of the function: if the query does not exist in the array, the function should return -1. This test is important as it checks the function's error handling.
*/

// ********RoostGPT********
package InterpolationSearch

import (
	"testing"
)

func Testinterpolationsearchfunc(t *testing.T) {
	// Define test cases
	testCases := []struct {
		arr     []int
		query   int
		want    int
		message string
	}{
		{
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:   5,
			want:    4,
			message: "Valid Query in the Middle of Array",
		},
		{
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:   1,
			want:    0,
			message: "Query at the Beginning of Array",
		},
		{
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:   10,
			want:    9,
			message: "Query at the End of Array",
		},
		{
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:   11,
			want:    -1,
			message: "Query Not in Array",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.message, func(t *testing.T) {
			got := interpolationSearchFunc(tc.arr, tc.query)

			if got != tc.want {
				t.Errorf("Failed %s: interpolationSearchFunc(%v, %d): want %d but got %d", 
					tc.message, tc.arr, tc.query, tc.want, got)
			} else {
				t.Logf("Success %s: interpolationSearchFunc(%v, %d): got expected %d", 
					tc.message, tc.arr, tc.query, got)
			}
		})
	}
}
