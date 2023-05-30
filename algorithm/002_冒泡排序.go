package algorithm

func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	size := len(arr)
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

}
