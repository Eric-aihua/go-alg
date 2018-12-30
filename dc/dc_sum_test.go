package dc

import (
	"fmt"
	"testing"
)

type dcSumResult struct {
	expected []int
	result   int
}

var testCases = []dcSumResult{
	{[]int{1, 2, 3, 4, 5}, 15},
	{[]int{1, 2, 3, 4, 2}, 12},
}

func TestBinarySearch(t *testing.T) {
	intList := make([]int, 100)
	for i := 0; i < 100; i++ {
		intList[i] = i
	}
	for _, testCase := range testCases {

		sumResult := Sum(testCase.expected)
		if testCase.result != sumResult {
			t.Error("not find!")
		}
	}

}
