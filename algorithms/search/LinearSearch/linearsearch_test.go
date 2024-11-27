package BinarySearch

import (
	"testing"
)


func Testlinearsearch(t *testing.T) {
	testCases := []struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}{{name: "Positive Test Case - Query Item Present in Array", arr: []int{1, 2, 3, 4, 5, 6}, query: 4, expected: 3}, {name: "Negative Test Case - Query Item Not Present in Array", arr: []int{1, 2, 3, 4, 5, 6}, query: 7, expected: -1}, {name: "Edge Case - Empty Array", arr: []int{}, query: 1, expected: -1}, {name: "Edge Case - Query Item Present Multiple Times in Array", arr: []int{1, 2, 3, 4, 4, 4, 5, 6}, query: 4, expected: 3}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := linearSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("linearSearch(%v, %v) = %v; want %v", tc.arr, tc.query, result, tc.expected)
			}
		})
	}
}

func linearSearch(arr []int, query int) int {
	for i, val := range arr {
		if val == query {
			return i
		}
	}
	return -1
}

