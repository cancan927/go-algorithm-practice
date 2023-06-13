package algorithm

func mergeSortRecursive(arr []int) []int {

	//处理边界条件
	if arr == nil || len(arr) < 2 {
		return arr
	}

	n := len(arr)
	help := make([]int, n)
	mergeSort(arr, help, 0, n-1)
	return arr

}

func mergeSort(arr []int, help []int, l, r int) {
	if l >= r {
		return
	}

	m := (l + r) / 2
	//左边有序！
	mergeSort(arr, help, l, m)
	//右边有序！
	mergeSort(arr, help, m+1, r)
	//整合
	merge(arr, help, l, m, r)

}

func merge(arr []int, help []int, l, m, r int) {
	p1 := l     //p1指向左边
	p2 := m + 1 //p2指向右边

	i := l
	for ; p1 <= m && p2 <= r; i++ { //p1和p2都不越界
		if arr[p1] <= arr[p2] { //谁小拷贝谁
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}

	for ; p1 <= m; i++ { // p2已经走完，只剩p1
		help[i] = arr[p1]
		p1++
	}
	for ; p2 <= r; i++ { // p1已经走完，只剩p2
		help[i] = arr[p2]
		p2++
	}

	for i := l; i <= r; i++ {
		arr[i] = help[i]
	}
}
