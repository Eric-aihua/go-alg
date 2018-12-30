package sort

import (
	"fmt"
	"testing"
)

type quickSortResult struct {
	unsotred []int
	result   []int
}

var quickTestCases = []quickSortResult{
	{[]int{1, 8, 3, 4, 5, 6}, []int{1, 3, 4, 5, 6, 8}},
	{[]int{1, 2, 2, 3, 8, 6}, []int{1, 2, 2, 3, 6, 8}},
}

func TestQuickSort(t *testing.T) {
	for _, testCase := range quickTestCases {
		sortResult := QuickSort(testCase.unsotred)
		fmt.Println(sortResult)

	}

}
