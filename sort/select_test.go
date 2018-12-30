package sort

import (
	"fmt"
	"testing"
)

type binarySearchResult struct {
	unsotred []int
	result   []int
}

var testCases = []binarySearchResult{
	{[]int{1, 8, 3, 4, 5, 6}, []int{1, 3, 4, 5, 6, 8}},
	{[]int{1, 2, 2, 3, 8, 6}, []int{1, 2, 2, 3, 6, 8}},
}

func TestBinarySearch(t *testing.T) {
	for _, testCase := range testCases {
		find_result := SelectSort(testCase.unsotred)
		fmt.Println(find_result)

	}

}
