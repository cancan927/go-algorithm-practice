package Queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue_Size(t *testing.T) {
	queue := NewArrayQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()
	if queue.size != 3 {
		t.Error("enqueue error")
	} else {
		t.Log("ok")
	}
}

func TestArrayQueue_DeQueue(t *testing.T) {
	queue := NewArrayQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()
	result, _ := queue.DeQueue()
	fmt.Println(result)
	queue.Print()
	if result != 1 || queue.size != 2 {
		t.Error("dequeue error")
	} else {
		t.Log("ok")
	}
}

func TestArrayQueue_Front(t *testing.T) {
	q1 := NewArrayQueue()
	q1.EnQueue(1)
	q1.EnQueue(2)
	q1.EnQueue(3)
	q1.Print()
	r1, _ := q1.Front()
	q2 := NewArrayQueue()
	r2, e2 := q2.Front()
	if r1 != 1 || r2 != nil || e2 == nil {
		t.Error("front error")
	} else {
		t.Log("ok")
	}
}

func TestLinkedQueue_Size(t *testing.T) {
	queue := NewLinkedQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	if queue.size != 4 {
		t.Error("size method error")
	} else {
		t.Log("size method ok")
	}
}

func TestLinkedQueue_Clear(t *testing.T) {
	queue := NewLinkedQueue()
	queue.Print()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.Print()
	queue.Clear()
	queue.Print()
	if queue.size != 0 {
		t.Error("clear method error")
	} else {
		t.Log("clear method ok")
	}
}

func TestLinkedQueue_DeQueue(t *testing.T) {
	queue := NewLinkedQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.Print()
	result, _ := queue.DeQueue()
	queue.Print()
	if result != 1 || queue.size != 3 {
		t.Error("dequeue method error")
	} else {
		t.Log("dequeue method ok")
	}
}
