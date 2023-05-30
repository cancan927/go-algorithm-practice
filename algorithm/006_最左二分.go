package algorithm

//有序数组中找到大于等于num的最左位置
func MostLeft(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	mid := 0
	ans := -1
	for l <= r {
		mid = l + (r-l)>>1
		if arr[mid] >= num { //中点符合条件，说明右侧全部满足条件，不用再判断
			ans = mid
			r = mid - 1 //检查左边
		} else { //中点不符合条件，说明左边全部不满足条件，往右边判断
			l = mid + 1
		}
	}
	return ans
}
