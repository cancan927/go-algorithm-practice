package algorithm

func SelectionSort(arr []int) {
	// 先处理边界条件
	if arr == nil || len(arr) < 2 {
		return
	}

	size := len(arr)
	for i := 0; i < size-1; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIndex] {
				swap(arr, j, minIndex)
			}
		}
	}
}
