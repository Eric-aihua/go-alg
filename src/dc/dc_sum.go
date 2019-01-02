package dc

// 使用dc的思想统计一个数组元素的和
func Sum(ints []int) int {
	if len(ints) == 1 {
		//当集合只有一个元素时，返回结果
		return ints[0]
	}
	subArr := ints[1:]
	//每次递归的时候，传个Sum的集合数量都比上一次少一个元素
	return ints[0] + Sum(subArr)
}
