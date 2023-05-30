package algorithm

// FindNum 有序数组中找到num
func FindNum(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	l := 0
	r := len(arr) - 1
	mid := 0
	for l <= r {
		mid = l + (r-l)>>1
		if arr[mid] == num { //中点正好命中
			return true
		} else if arr[mid] > num { //中点比num大，说明num在左边
			r = mid - 1
		} else { // 中点比num小，说明num在右边
			l = mid + 1
		}
	}
	return false
}
