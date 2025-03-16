package BinarySearch

import (
	"testing"
)


func Testinterpolationsearchfunc(t *testing.T) {
	testCases := []struct {
		arr	[]int
		query	int
		want	int
		message	string
	}{{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 5, want: 4, message: "Valid Query in the Middle of Array"}, {arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 1, want: 0, message: "Query at the Beginning of Array"}, {arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 10, want: 9, message: "Query at the End of Array"}, {arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 11, want: -1, message: "Query Not in Array"}}
	for _, tc := range testCases {
		t.Run(tc.message, func(t *testing.T) {
			got := interpolationSearchFunc(tc.arr, tc.query)
			if got != tc.want {
				t.Errorf("Failed %s: interpolationSearchFunc(%v, %d): want %d but got %d", tc.message, tc.arr, tc.query, tc.want, got)
			} else {
				t.Logf("Success %s: interpolationSearchFunc(%v, %d): got expected %d", tc.message, tc.arr, tc.query, got)
			}
		})
	}
}

func interpolationSearchFunc(arr []int, query int) int {
	lo := 0
	hi := len(arr) - 1
	for lo <= hi && query >= arr[lo] && query <= arr[hi] {
		midIndex := lo + (query-arr[lo])*(hi-lo)/(arr[hi]-arr[lo])
		midItem := arr[midIndex]
		if midItem == query {
			return midIndex
		} else if midItem < query {
			lo = midIndex + 1
		} else if midItem > query {
			hi = midIndex - 1
		}
	}
	return -1
}

