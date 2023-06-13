package algorithm

/**
给定一个数组nums，求和为K的最长子数组长度,如不存在返回-1
例子：
K = 3，arr={5，-2，1，1，3，2}
最长子数组长度为3

思路：利用前缀和求解
*/

func SumKLength(arr []int, k int) int {
	//先判断边界条件
	if arr == nil || (len(arr) == 0 && k != 0) {
		return -1
	}

	// 最大长度
	ans := 0
	// 整体累加和
	sum := 0
	// 用map存储前缀和，key为前缀和，value为下标
	sumMap := make(map[int]int)
	sumMap[0] = -1 //一个数也没有的时候，就已经存在前缀和0了

	for i, v := range arr {
		// 0-i 整体前缀和
		sum += v

		pre, ok := sumMap[sum-k]

		if ok && ans < i-pre {
			ans = i - pre
		}

		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}

	return ans
}
