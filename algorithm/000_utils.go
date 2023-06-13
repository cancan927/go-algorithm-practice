package algorithm

import (
	"math/rand"
	"time"
)

// 切片中的两个元素交换位置
func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

//生成长度为[1,n]，取值范围为[1,v]的随机数组
func generateRandomArray(length, peek int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(length) + 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(peek) + 1
	}
	return arr
}

func copyArray(src []int) []int {
	if src == nil {
		return nil
	}
	length := len(src)
	des := make([]int, length)
	for i := 0; i < length; i++ {
		des[i] = src[i]
	}
	return des
}
