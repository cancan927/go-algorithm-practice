package Stack

import "errors"

type LinkedStack struct {
	root Node
	size int
}

type Node struct {
	next, prev *Node
	value      any
}

func (l *LinkedStack) Init() *LinkedStack {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.size = 0
	return l
}

func NewLikedStack() *LinkedStack {
	return new(LinkedStack).Init()
}

func (l *LinkedStack) Size() int {
	return l.size
}

func (l *LinkedStack) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedStack) Peek() (any, error) {
	if l.size == 0 {
		return nil, errors.New("stack is empty")
	}
	value := l.root.prev.value
	return value, nil
}

func (l *LinkedStack) Push(a any) {
	//压栈的元素
	newVal := &Node{value: a}
	newVal.next = &l.root
	newVal.prev = l.root.prev
	l.root.prev.next = newVal
	l.root.prev = newVal
	l.size++
}

func (l *LinkedStack) Pop() (any, error) {
	if l.size == 0 {
		return nil, errors.New("stack is empty")
	}
	top := l.root.prev
	top.prev.next = top.next
	top.next.prev = top.prev
	l.size--
	return top.value, nil
}
