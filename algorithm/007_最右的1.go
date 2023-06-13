package algorithm

// 一串二进制数，只保留最右的1
// 1001 0100  => 0000 0100

func RightOne(num int) int {
	return num & (-num)
}
