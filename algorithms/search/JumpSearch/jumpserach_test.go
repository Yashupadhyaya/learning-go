package BinarySearch

import (
	"testing"
)

import (
	"math"
	"testing"
)


func SearchLinerFunc(arr []int, query int) int {
	for i, val := range arr {
		if val == query {
			return i
		}
	}
	return -1
}

func TestSearchLinerFunc(t *testing.T) {
	testCases := []struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}{{name: "Search for an existing number in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 7, expected: 6}, {name: "Search for a non-existing number in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 11, expected: -1}, {name: "Search in an empty array", arr: []int{}, query: 0, expected: -1}, {name: "Search for the first number in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 1, expected: 0}, {name: "Search for the last number in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 10, expected: 9}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SearchLinerFunc(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("failed %s: got %v, expected %v", tc.name, result, tc.expected)
			} else {
				t.Logf("success %s", tc.name)
			}
		})
	}
}

func Testjumpsearch(t *testing.T) {
	tests := []struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}{{name: "Query integer is present in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 7, expected: 6}, {name: "Query integer is not present in the array", arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, query: 11, expected: -1}, {name: "Empty array", arr: []int{}, query: 1, expected: -1}, {name: "Array with duplicate integers", arr: []int{1, 2, 3, 3, 4, 4, 5, 5, 6, 6}, query: 4, expected: 3}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpSearch(tt.arr, tt.query); got != tt.expected {
				t.Errorf("jumpSearch() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func jumpSearch(arr []int, query int) int {
	size := len(arr)
	step := int(math.Sqrt(float64(size)))
	prev := 0
	if size == 0 {
		return -1
	}
	for arr[int(math.Min(float64(step), float64(size)))-1] < query {
		prev = step
		step += int(math.Sqrt(float64(size)))
		if prev >= size {
			return -1
		}
	}
	for arr[prev] < query {
		prev++
		if prev == int(math.Min(float64(step), float64(size))) {
			return -1
		}
	}
	if arr[prev] == query {
		return prev
	}
	return -1
}

