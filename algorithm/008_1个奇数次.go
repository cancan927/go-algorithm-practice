package algorithm

// 数组中只有一个数出现了奇数次，找到它

func OddOne(arr []int) int {
	result := 0
	for _, v := range arr {
		result = result ^ v
	}
	return result
}
