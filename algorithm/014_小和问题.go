package algorithm

//在一个数组中，一个数的左边比它小的数的和称为小和
//数组中所有数的小和加起来叫做数组的小和

func littleSum(arr []int) int {
	//先处理边界条件
	if arr == nil || len(arr) < 2 {
		return 0
	}

	n := len(arr)
	help := make([]int, n)

	return process(arr, help, 0, n-1)

}

//arr[l,,,r]范围上，排好序，并且求出这个范围在变有序的过程中，产生的小和
func process(arr, help []int, l, r int) int {
	if l >= r {
		return 0
	}

	m := (l + r) / 2
	return process(arr, help, l, m) + process(arr, help, m+1, r) + merge2(arr, help, l, m, r)

}

func merge2(arr []int, help []int, l int, m int, r int) int {
	ans := 0
	p1 := l
	p2 := m + 1
	i := l

	for ; p1 <= m && p2 <= r; i++ {
		if arr[p1] < arr[p2] {
			ans += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
	}

	for ; p1 <= m; i++ {
		help[i] = arr[p1]
		p1++
	}

	for ; p2 <= r; i++ {
		help[i] = arr[p2]
		p2++
	}

	for i := l; i <= r; i++ {
		arr[i] = help[i]
	}
	return ans
}
