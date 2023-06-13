package parctice

import (
	"fmt"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func GeneralSorter[T any](arr []T, compare func(o1, o2 T) int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minPos := 0
		for j := i + 1; j < n; j++ {
			if compare(arr[j], arr[minPos]) == -1 { //如果后面的位置比当前位置小
				minPos = j
			}
		}
		arr[i], arr[minPos] = arr[minPos], arr[i]
	}

}

func main() {
	persons := []Person{
		{Id: 1, Name: "zhangsan", Age: 19},
		{2, "lisi", 31},
		{3, "wangwu", 16}}

	GeneralSorter[Person](persons, func(o1, o2 Person) int {
		if o1.Age < o1.Age {
			return 1
		} else if o1.Age == o2.Age {
			return 0
		} else {
			return -1
		}
	})

	fmt.Println(persons)

}
