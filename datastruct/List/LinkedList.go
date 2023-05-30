package List

import (
	"errors"
	"fmt"
)

//链表
type LinkedList struct {
	//根节点
	root Node
	len  int
}

//节点
type Node struct {
	//前后指针
	//这里不提供包外访问是因为防止一些危险操作，提供方法来获取前驱和后继
	next, prev *Node

	//值
	Value any

	//指向链表的指针
	list *LinkedList
}

//初始化链表，根节点前驱和后继都指向自己
func (l *LinkedList) Init() *LinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func NewLinkedList() *LinkedList {
	return new(LinkedList).Init()
}

func (n *Node) Next() *Node {
	//如果当前节点所属链表不为nil且当前节点的后继不是根节点
	if p := n.next; p.list != nil && p != &p.list.root {
		return p
	}
	return nil
}

func (n *Node) Prev() *Node {
	//如果当前节点所属链表不为空且当前节点的前驱不是根节点
	if p := n.prev; p.list != nil && p != &p.list.root {
		return p
	}
	return nil
}

func (l *LinkedList) Size() int {
	return l.len
}

//返回链表的第一个节点
func (l *LinkedList) Front() *Node {
	//如果链表为空
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

//返回链表的最后一个节点
func (l *LinkedList) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

//判空
func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

//清空
func (l *LinkedList) Clear() {
	l.Init()
}

func (l *LinkedList) Contains(val any) bool {
	if l.IsEmpty() {
		return false
	}
	p := &l.root
	for i := 0; i < l.len; i++ {
		p = p.next
		if p.Value == val {
			return true
		}
	}
	return false
}

func (l *LinkedList) GetNode(index int) (*Node, error) {
	if index < 0 || index >= l.len {
		return nil, errors.New("index out of bounds")
	}
	p := &l.root
	if index < l.len>>1 { //小于链表长度一半，从根节点向后遍历
		for i := 0; i <= index; i++ {
			p = p.next
		}
	} else {
		for i := (l.len - 1); i >= index; i-- {
			p = p.prev
		}
	}
	return p, nil
}

func (l *LinkedList) Get(index int) (any, error) {
	if index < 0 || index >= l.len {
		return nil, errors.New("index out of bounds")
	}
	p := &l.root
	if index < l.len>>1 { //小于链表长度一半，从根节点向后遍历
		for i := 0; i <= index; i++ {
			p = p.next
		}
	} else {
		for i := (l.len - 1); i >= index; i-- {
			p = p.prev
		}
	}
	return p.Value, nil
}

func (l *LinkedList) Set(index int, newVal any) error {
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	p := &l.root
	for i := 0; i <= index; i++ {
		p = p.next
	}
	p.Value = newVal
	return nil
}

func (l *LinkedList) Insert(index int, newVal any) error {
	if index < 0 || index > l.len {
		return errors.New("index out of bounds")
	}

	if index == l.len {
		l.Append(newVal)
		return nil
	}
	p := &l.root
	for i := 0; i <= index; i++ {
		p = p.next //index ==0 时，p指向首节点
	}
	//把新元素插入到index位置元素的前面
	ele := &Node{Value: newVal}
	//新节点前驱指向老节点的前驱
	ele.prev = p.prev //index == 0 时，p.prev 指向root   <--E
	//新节点后继指向老节点
	ele.next = p //  E-->
	//老前驱节点的后继指向新节点
	ele.prev.next = ele //root.next = ele    -->E
	//老节点前驱指向新节点
	ele.next.prev = ele //       E<--
	ele.list = l
	l.len++
	return nil
}

func (l *LinkedList) Append(newVal any) {
	p := &Node{Value: newVal}
	//新节点后继指向根节点
	p.next = &l.root
	//新节点的前驱指向之前的尾节点
	p.prev = l.root.prev
	//旧尾节点的后继指向新节点
	l.root.prev.next = p
	//根节点的前驱指向新节点
	l.root.prev = p
	p.list = l
	l.len++
}

func (l *LinkedList) Delete(index int) error {
	if index < 0 || index >= l.len {
		return errors.New("index out of bounds")
	}
	p := &l.root
	for i := 0; i <= index; i++ {
		p = p.next
	}
	p.prev.next = p.next
	p.next.prev = p.prev
	l.len--
	return nil
}

//将整个链表打印在控制台上方便查看
func (l *LinkedList) Print() {
	fmt.Print("data:[")
	p := l.root.next
	for i := 0; i < l.len; i++ {
		if i == l.len-1 {
			fmt.Printf("%v", p.Value)
		} else {
			fmt.Printf("%v,", p.Value)
		}
		p = p.next
	}
	fmt.Printf("],size:%v\n", l.len)

}
