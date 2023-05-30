package Stack

type Stack interface {
	//返回栈的大小
	Size() int

	//判空
	IsEmpty() bool

	//返回栈顶元素
	Peek() (any, error)

	//压栈
	Push(any)

	//弹栈
	Pop() (any, error)
}
