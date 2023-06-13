package algorithm

func getMax(arr []int) int {
	return reverseGetMax(arr, 0, len(arr)-1)
}

func reverseGetMax(arr []int, l int, r int) int {
	if l == r {
		return arr[l]
	}

	m := (l + r) / 2 //中点
	leftMax := reverseGetMax(arr, l, m)
	rightMax := reverseGetMax(arr, m+1, r)

	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}

}
