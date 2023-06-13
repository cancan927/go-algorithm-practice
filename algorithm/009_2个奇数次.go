package algorithm

// 数组中有两个数出现了奇数次，找到它们

func OddTwo(arr []int) (a, b int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
	}
	// eor = a ^ b

	// 把所有数分为两类
	rightOne := eor & (-eor) // 找到最右的1，假设在第N位上是1
	onlyOne := 0
	for i := 0; i < len(arr); i++ {
		if (arr[i] & rightOne) != 0 { // 把数分为两类，N位是1和N位是0的
			onlyOne = onlyOne ^ arr[i]
		}
	}

	return onlyOne, eor ^ onlyOne
}
