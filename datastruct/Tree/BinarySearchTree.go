package Tree

type BinarySearchTree struct {
	root *Node
	size int
}

/**
注意：
二叉搜索树的元素必须是可比较的，否则就失去了意义,这里只使用int来避免这种问题
*/
type Node struct {
	left, right, parent *Node
	value               int
}

func NewBinarySearchTree() *BinarySearchTree {
	tree := &BinarySearchTree{size: 0}
	return tree
}

func (b *BinarySearchTree) Size() int {
	return b.size
}

func (b *BinarySearchTree) IsEmpty() bool {
	return b.size == 0
}

func (b *BinarySearchTree) Clear() {
	b = NewBinarySearchTree()
}

func (b *BinarySearchTree) Add(n int) {
	//最开始添加根节点
	if b.root == nil {
		b.root = &Node{
			value: n,
		}
	} else {
		//添加的不是第一个节点,找到父节点
		cur := b.root //记录当前节点
		ori := 0      //记录方向
		parent := &Node{}
		for cur != nil {
			parent = cur
			if n < cur.value {
				ori = -1
				cur = cur.left
			} else if n > cur.value {
				ori = 1
				cur = cur.right
			} else { //相等什么也不做
				return
			}
		}
		if ori > 0 {
			parent.right = &Node{
				value:  n,
				parent: parent,
			}
		} else {
			parent.left = &Node{
				value:  n,
				parent: parent,
			}
		}
	}
	b.size++
	return
}

func (b *BinarySearchTree) Remove() error {
	return nil
}

func (b *BinarySearchTree) Contains(c any) bool {
	return false
}

//中序遍历
func (b *BinarySearchTree) InOrderTraversal() (res []int) {
	return inOrder(b.root)

}

func inOrder(node *Node) (res []int) {
	if node == nil {
		return nil
	}
	res = make([]int, 0)
	res = append(res, inOrder(node.left)...)
	res = append(res, node.value)
	res = append(res, inOrder(node.right)...)
	return res
}

//先序遍历
func (b *BinarySearchTree) PreOrderTraversal() (res []int) {
	return preOrder(b.root)
}

func preOrder(node *Node) (res []int) {
	if node == nil {
		return nil
	}
	res = make([]int, 0)
	res = append(res, node.value)
	res = append(res, preOrder(node.left)...)
	res = append(res, preOrder(node.right)...)
	return res
}

//后序遍历
func (b *BinarySearchTree) PostOrderTraversal() (res []int) {
	return postOrder(b.root)
}

func postOrder(node *Node) (res []int) {
	if node == nil {
		return nil
	}
	res = make([]int, 0)
	res = append(res, postOrder(node.left)...)
	res = append(res, postOrder(node.right)...)
	res = append(res, node.value)
	return res
}
