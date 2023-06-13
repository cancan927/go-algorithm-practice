package parctice

func SelectSort(arr []int) []int {

	//1.处理边界条件
	if arr == nil || len(arr) < 2 {
		return arr
	}

	//2.不断选出最小值并移动到前面
	n := len(arr)
	// [0,n-1] [1,n-1] [2,n-1]...[n-2,n-1]
	for i := 0; i < n-1; i++ {
		minPos := i //记录当前最小位置
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minPos] { //如果后面的数比当前最小值小，就更新最小值位置
				minPos = j
			}
		}
		// 一轮下来，找到当前循环最小值位置，跟i交换
		arr[i], arr[minPos] = arr[minPos], arr[i]
	}
	return arr
}
