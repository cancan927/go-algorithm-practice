package Queue

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	elements []any
	size     int
}

func NewArrayQueue() *ArrayQueue {
	queue := new(ArrayQueue)
	queue.elements = make([]any, 0, 10)
	queue.size = 0
	return queue
}

func (a *ArrayQueue) Size() int {
	return a.size
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.size == 0
}

//入队列
func (a *ArrayQueue) EnQueue(arg any) {
	a.elements = append(a.elements, arg)
	a.size++
}

//出队列
func (a *ArrayQueue) DeQueue() (any, error) {
	if a.size == 0 {
		return nil, errors.New("queue is empty")
	} else {
		result := a.elements[0]
		a.elements = a.elements[1:]
		a.size--
		return result, nil
	}
}

//获取队头元素
func (a *ArrayQueue) Front() (any, error) {
	if a.size == 0 {
		return nil, errors.New("queue is empty")
	}
	return a.elements[0], nil
}

//清空整个队列
func (a *ArrayQueue) Clear() {
	a = NewArrayQueue()
}

//打印整个队列到控制台上
func (a *ArrayQueue) Print() {
	fmt.Printf("data:%v,size:%v\n", a.elements, a.size)
}
