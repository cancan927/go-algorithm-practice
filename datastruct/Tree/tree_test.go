package Tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Add_Size(t *testing.T) {
	b := new(BinarySearchTree)
	s := []int{6, 3, 4, 8, 2, 1}
	for i := 0; i < len(s); i++ {
		b.Add(s[i])
	}
	fmt.Println("size:", b.Size())

}

func TestBinarySearchTree_PreOrderTraversal(t *testing.T) {
	b := new(BinarySearchTree)
	s := []int{6, 3, 4, 8, 2, 1}
	for i := 0; i < len(s); i++ {
		b.Add(s[i])
	}
	traversal := b.PreOrderTraversal()
	for _, v := range traversal {
		fmt.Println(v)
	}
}

func TestBinarySearchTree_InOrderTraversal(t *testing.T) {
	b := new(BinarySearchTree)
	s := []int{6, 3, 4, 8, 2, 1}
	for i := 0; i < len(s); i++ {
		b.Add(s[i])
	}
	traversal := b.InOrderTraversal()
	for _, v := range traversal {
		fmt.Println(v)
	}
}

func TestBinarySearchTree_PostOrderTraversal(t *testing.T) {
	b := new(BinarySearchTree)
	s := []int{6, 3, 4, 8, 2, 1}
	for i := 0; i < len(s); i++ {
		b.Add(s[i])
	}
	traversal := b.PostOrderTraversal()
	for _, v := range traversal {
		fmt.Println(v)
	}
}
