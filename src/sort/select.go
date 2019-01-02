package sort

// 对int数组的选择排序
func SelectSort(items []int) []int {
	result := make([]int, len(items))
	for i := 0; i < len(result); i++ {
		tmpIndex := getMinItemIndex(&items)
		result[i] = items[tmpIndex]
		//从原始数组中删除掉被选择的元素
		items = append(items[:tmpIndex], items[tmpIndex+1:]...)
	}
	return result
}

//获取一个数组中，最小元素的位置
func getMinItemIndex(ints *[]int) int {
	//默认假设第一个元素作为临时最小值
	min := (*ints)[0]
	smallestIndex := 0
	for index, item := range *ints {
		if item < min {
			min = item
			smallestIndex = index
		}
	}
	return smallestIndex
}
