package algorithm

func InsertSort(arr []int) {

	// 先处理边界条件
	if arr == nil || len(arr) < 2 {
		return
	}

	size := len(arr)
	for i := 1; i < size; i++ { // 0-0位置天然有序，从0-1位置开始判断
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}
