package Stack

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	elements []any
	size     int
}

func NewArrayStack() *ArrayStack {
	stack := new(ArrayStack)
	stack.elements = make([]any, 0, 10)
	stack.size = 0
	return stack
}

func (a *ArrayStack) Size() int {
	return a.size
}

func (a *ArrayStack) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayStack) Peek() (any, error) {
	if a.size == 0 {
		return nil, errors.New("stack is empty")
	}
	return a.elements[a.size-1], nil
}

func (a *ArrayStack) Push(any) {
	newStack := append(a.elements, a)
	a.elements = newStack
	a.size++
}

//弹出栈顶元素
func (a *ArrayStack) Pop() (any, error) {
	if a.size == 0 {
		return nil, errors.New("stack is empty")
	}
	top := a.elements[a.size-1]
	newStack := a.elements[:a.size-1]
	a.elements = newStack
	a.size--
	return top, nil
}

func (a *ArrayStack) ToString() {
	fmt.Printf("ele:%v,size:%v\n", a.elements, a.size)
}
