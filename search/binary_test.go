package search

import (
	"testing"
)

type binarySearchResult struct {
	expected int
	result   int
}

var testCases = []binarySearchResult{
	{98, 98},
	{101, -1},
	{-1, -1},
}

func TestBinarySearch(t *testing.T) {
	intList := make([]int, 100)
	for i := 0; i < 100; i++ {
		intList[i] = i
	}

	for _, testCase := range testCases {

		find_result := BinarySearch(intList, testCase.expected)
		if testCase.result != find_result {
			t.Error("not find!")
		}
	}

}
