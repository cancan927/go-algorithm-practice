package algorithm

// 切片中的两个元素交换位置
func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
