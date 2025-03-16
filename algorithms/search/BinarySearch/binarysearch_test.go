package BinarySearch

import (
	"testing"
)


func SearchLinear(arr []int, query int) int {
	for i, val := range arr {
		if val == query {
			return i
		}
	}
	return -1
}

func TestSearchLinear(t *testing.T) {
	tests := []struct {
		name	string
		arr	[]int
		query	int
		expect	int
	}{{name: "Positive Test Case - Query Item Present in Array", arr: []int{1, 2, 3, 4, 5}, query: 3, expect: 2}, {name: "Negative Test Case - Query Item Not Present in Array", arr: []int{1, 2, 3, 4, 5}, query: 10, expect: -1}, {name: "Edge Case - Empty Array", arr: []int{}, query: 1, expect: -1}, {name: "Edge Case - Array with Duplicate Items", arr: []int{1, 2, 3, 3, 5}, query: 3, expect: 2}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SearchLinear(test.arr, test.query)
			if got != test.expect {
				t.Errorf("SearchLinear(%v, %v) = %v; want %v", test.arr, test.query, got, test.expect)
			}
		})
	}
}

func Testbinarysearchfunc(t *testing.T) {
	testCases := []struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}{{name: "Successful binary search with query present in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 6, expected: 5}, {name: "Binary search with query not present in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 11, expected: -1}, {name: "Binary search with an empty array", arr: []int{}, query: 1, expected: -1}, {name: "Binary search with a single-element array, query matches the element", arr: []int{1}, query: 1, expected: 0}, {name: "Binary search with a single-element array, query does not match the element", arr: []int{1}, query: 2, expected: -1}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearchFunc(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("binarySearchFunc(%v, %v): expected %v, got %v", tc.arr, tc.query, tc.expected, result)
			}
		})
	}
}

func binarySearchFunc(arr []int, query int) int {
	minIndex := 0
	maxIndex := len(arr) - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := arr[midIndex]
		if query == midItem {
			return midIndex
		}
		if midItem < query {
			minIndex = midIndex + 1
		} else if midItem > query {
			maxIndex = midIndex - 1
		}
	}
	return -1
}

