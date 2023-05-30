package List

import (
	"fmt"
	"testing"
)

func TestArrayListAppend(t *testing.T) {
	list := NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Print()
	if list.Size() == 5 {
		t.Log("append ok")
	} else {
		t.Error("append error")
	}
}

func TestArrayList_Size(t *testing.T) {
	l1 := NewArrayList()
	l1.Print()
	l2 := NewArrayList()
	l2.Append(1)
	l2.Append(2)
	l2.Append(3)
	l2.Print()
	if l1.Size() != 0 || l2.Size() != 3 {
		t.Error("size error")
	} else {
		t.Log("ok")
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	l1 := NewArrayList()
	l2 := NewArrayList()
	l2.Append(1)
	if l1.IsEmpty() != true || l2.IsEmpty() != false {
		t.Error("isEmpty error")
	} else {
		t.Log("ok")
	}
}

func TestArrayList_Clear(t *testing.T) {
	l1 := NewArrayList()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Print()
	l1.Clear()
	l1.Print()
	if l1.Size() != 0 {
		t.Error("clear error")
	} else {
		t.Log("ok")
	}
}

func TestArrayList_Get(t *testing.T) {
	l1 := NewArrayList()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	res1, err1 := l1.Get(0)
	l2 := NewArrayList()
	res2, err2 := l2.Get(1)
	if res1 != 1 || err1 != nil || res2 != nil || err2 == nil {
		t.Error("get error")
	} else {
		t.Log("ok")
	}
}

func TestLinkedList_Size(t *testing.T) {
	l1 := NewLinkedList()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Print()
	l2 := NewLinkedList()
	l2.Print()
	if l1.Size() != 3 || l2.Size() != 0 {
		t.Error("size error")
	} else {
		t.Log("ok")
	}
}

func TestLinkedList_Append(t *testing.T) {
	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Print()
	if list.Size() != 4 {
		t.Error("append error")
	} else {
		t.Log("ok")
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	l1 := NewArrayList()
	if l1.IsEmpty() != true {
		t.Error("isEmpty error")
	} else {
		t.Log("ok")
	}
}

func TestLinkedList_Clear(t *testing.T) {
	l1 := NewLinkedList()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Print()
	l1.Clear()
	l1.Print()
	if l1.Size() != 0 {
		t.Error("clear error")
	} else {
		t.Log("ok")
	}
}

func TestLinkedList_Get(t *testing.T) {
	l1 := NewLinkedList()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	res1, err1 := l1.Get(0)
	l2 := NewLinkedList()
	res2, err2 := l2.Get(1)
	fmt.Println(res1, err1, res2, err2)
	if res1 != 1 || err1 != nil || res2 != nil || err2 == nil {
		t.Error("get error")
	} else {
		t.Log("ok")
	}
}
