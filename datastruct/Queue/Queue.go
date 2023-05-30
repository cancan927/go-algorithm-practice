package Queue

type Queue interface {
	//返回队列元素个数
	Size() int

	//判断队列是否为空
	IsEmpty() bool

	//入队列
	EnQueue(any)

	//出队列
	DeQueue() (any, error)

	//返回队头元素
	Front() (any, error)

	//清空队列
	Clear()
}
