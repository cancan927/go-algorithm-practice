package algorithm

//非递归归并排序

func mergeSortUnrecursived(arr []int) []int {
	//先处理边界条件
	if arr == nil || len(arr) < 2 {
		return arr
	}

	n := len(arr)
	help := make([]int, n)

	//步长
	// 步长<n,才继续
	// 步长>=n,已经排好序了，可以停止了
	mergeSize := 1

	for mergeSize < n {
		// 左组从0开始
		l := 0

		// l >= n 说明不再有左右对了
		for l < n {

			// 最后的左右对，要么左组不足，要么左组足了，没有右组
			// 直接跳出当前步长的调整，去下一个步长
			if mergeSize >= n-1 {
				break
			}

			m := l + mergeSize - 1

			r := m + min(mergeSize, n-m-1)
			merge(arr, help, l, m, r)
			l = r + 1
		}

		if mergeSize > n/2 {
			break
		}

		mergeSize *= 2

	}
	return arr
}

func min(i, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}
