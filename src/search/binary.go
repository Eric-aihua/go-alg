package search

// 使用二分查找法，在有序列表intList中查找target,返回对应的索引
func BinarySearch(intList []int, target int) int {
	start := 0
	end := len(intList) - 1
	for {
		if end <= start {
			break
		}
		mid := int((start + end) / 2)
		if intList[mid] == target {
			return mid
		}
		// 如果mid处的值大于目标值，就在mid之前的数据中搜索
		if intList[mid] > target {
			end = mid - 1
		}
		// 如果mid处的值大于目标值，就在mid之后的数据中搜索
		if intList[mid] < target {
			start = mid + 1
		}
		//fmt.Println(start, end)
	}
	return -1
}
