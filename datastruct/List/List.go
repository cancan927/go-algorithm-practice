package List

type List[T any] interface {
	// Size 获取list长度
	Size() int

	// IsEmpty 判空
	IsEmpty() bool

	// Clear 清空
	Clear()

	// Get 获取指定下标元素
	Get(index int) (T, error)

	// Contains 判断list是否包含指定元素
	Contains(val T) bool

	// Set 修改指定下标元素
	Set(index int, newVal T) error

	// Insert 在指定位置插入元素
	Insert(index int, newVal T) error

	// Append 在list末尾追加元素
	Append(newVal T)

	// Delete 删除指定下标元素
	Delete(index int) error
}
