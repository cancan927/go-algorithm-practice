package algorithm

/**
给定a,b两个下标，求[a,b]范围的值的和,假定a,b不会越界
例如：[]int{1,2,3,4,5}, 给定2,4	和为3+4+5=12
*/

func PrefixSum(arr []int, l, r int) int {
	size := len(arr)
	sum := make([]int, size)

	// 0-0
	sum[0] = arr[0]
	for i := 1; i < size; i++ {
		sum[i] = sum[i-1] + arr[i]
	}

	if l == 0 {
		return sum[0]
	} else {
		return sum[r] - sum[l-1]
	}
}
