package List

import (
	"errors"
	"fmt"
	"reflect"
)

type ArrayList[T any] struct {
	dataStore []T
	size      int
}

func NewArrayList() *ArrayList[any] {
	list := &ArrayList[any]{
		dataStore: make([]any, 0, 10),
		size:      0,
	}
	return list
}

func (a *ArrayList[T]) Size() int {
	return a.size
}

func (a *ArrayList[T]) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList[T]) Clear() {
	a.dataStore = make([]T, 0, 10)
	a.size = 0
}

func (a *ArrayList[T]) Get(index int) (any, error) {
	//下标范围是0 ~ size-1
	if index < 0 || index >= a.size {
		return nil, errors.New("index out of bounds")
	}
	return a.dataStore[index], nil
}

func (a *ArrayList[T]) Contains(val T) bool {
	for _, v := range a.dataStore {
		if reflect.DeepEqual(v, val) {
			return true
		}
	}
	return false
}

//将下标为index的值设置为newVal
func (a *ArrayList[T]) Set(index int, newVal any) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	a.dataStore[index] = newVal
	return nil
}

func (a *ArrayList[T]) Insert(index int, newVal any) error {
	if index < 0 || index > a.size {
		return errors.New("index out of bounds")
	}
	//考虑两种特殊情况，插入最前面和最后面
	if index == 0 { //插入最前面
		newStore := make([]T, 0, 10)
		newStore = append(newStore, newVal)
		a.dataStore = append(newStore, a.dataStore[:]...)
		a.size++
		return nil
	}
	if index == a.size { //插入最后面
		a.Append(newVal)
		return nil
	}
	//一般情况
	newStore := make([]T, 0, 10)
	newStore = append(newStore, a.dataStore[:index]...) //前半部分
	newStore = append(newStore, newVal)
	newStore = append(newStore, a.dataStore[index:]...) //后半部分
	a.dataStore = newStore
	a.size++
	return nil
}

func (a *ArrayList[T]) Append(newVal any) {
	a.dataStore = append(a.dataStore, newVal)
	a.size++
}

//删除指定位置的元素
func (a *ArrayList[T]) Delete(index int) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	a.dataStore = append(a.dataStore[:index], a.dataStore[index+1:]...)
	a.size--
	return nil
}

//将ArrayList以字符串形式打印在控制台上方便查看
func (a *ArrayList[T]) Print() {
	fmt.Printf("data:%v,size:%v\n", a.dataStore, a.size)
}
