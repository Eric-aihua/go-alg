package sort

import (
	"fmt"
)

// 对int数组的快速排序
func QuickSort(items []int) []int {
	//fmt.Println(items)
	if len(items) < 2 {
		// 对于长度小于2的数组，不需要排序
		return items
	} else {
		//1. 选取第一个数为基准值
		randomIndex := 0
		//fmt.Println("index:",randomIndex)
		pivot := items[randomIndex]
		//利用pivot,将items分成两部分
		leftArray := make([]int, 0)
		rightArray := make([]int, 0)
		for _, item := range items[1:] {
			if item >= pivot {
				rightArray = append(rightArray, item)
			}
			if item < pivot {
				leftArray = append(leftArray, item)
			}
		}
		// 将leftArray,pivot,rightArray进行拼接
		return merge(QuickSort(leftArray), []int{pivot}, QuickSort(rightArray))
	}
}

func merge(s1 []int, s2 []int, s3 []int) []int {
	result := make([]int, 0)
	for _, e := range s1 {
		result = append(result, e)
	}
	for _, e := range s2 {
		result = append(result, e)
	}
	for _, e := range s3 {
		result = append(result, e)
	}
	return result
}

func main() {
	fmt.Println(QuickSort([]int{3, 5, 6, 2, 6, 2}))
}
