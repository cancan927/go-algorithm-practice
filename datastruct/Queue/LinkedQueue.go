package Queue

import (
	"errors"
	"fmt"
)

type LinkedQueue struct {
	root Node
	size int
}

type Node struct {
	next, prev *Node
	value      any
}

func (l *LinkedQueue) Init() *LinkedQueue {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.size = 0
	return l
}

func NewLinkedQueue() *LinkedQueue {
	return new(LinkedQueue).Init()
}

//返回队列元素个数
func (l *LinkedQueue) Size() int {
	return l.size
}

//判断队列是否为空
func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}

//将一个元素加入队列尾部
func (l *LinkedQueue) EnQueue(a any) {
	newVal := &Node{value: a}
	//新元素后继指向根节点
	newVal.next = &l.root
	//新元素前驱指向原来的尾节点
	newVal.prev = l.root.prev
	//原来尾节点的后继指向新节点
	l.root.prev.next = newVal
	//根节点的前驱指向新节点
	l.root.prev = newVal
	l.size++
}

//将队头元素出队列
func (l *LinkedQueue) DeQueue() (any, error) {
	if l.size == 0 {
		return nil, errors.New("queue is empty")
	}
	front := l.root.next
	//将第二个元素的前驱指向front的前驱
	front.next.prev = &l.root
	//将根节点的后继指向第二个元素
	l.root.next = front.next
	l.size--
	return front.value, nil
}

//查看队头元素
func (l *LinkedQueue) Front() (any, error) {
	if l.size == 0 {
		return nil, errors.New("Queue is empty")
	} else {
		return l.root.next.value, nil
	}

}

//清空队列
func (l *LinkedQueue) Clear() {
	l.Init()
}

func (l *LinkedQueue) Print() {
	fmt.Print("data:[")
	p := l.root.next
	for i := 0; i < l.size; i++ {
		if i == l.size-1 {
			fmt.Printf("%v", p.value)
		} else {
			fmt.Printf("%v,", p.value)
		}
		p = p.next
	}
	fmt.Printf("],size:%v\n", l.size)
}
