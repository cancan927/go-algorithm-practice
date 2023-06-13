package algorithm

//给定一个数组arr和一个数num，
//小于等于这个数的都在左边
//大于这个数的都在右边
//要求额外空间复杂度为O(1)，时间复杂度为O(N)

func partition(arr []int, p int) {
	//1.先处理边界条件
	if arr == nil || len(arr) < 2 {
		return
	}
	//2.less区域
	less := -1                            //less区域一开始位于-1位置
	for cur := 0; cur < len(arr); cur++ { //[0,n-1]
		if arr[cur] <= p { //当前位置小于等于目标值
			//交换当前值与less区的下一个值
			arr[cur], arr[less+1] = arr[less+1], arr[cur]
			less++ //less区右移一位
		}
	}

}
