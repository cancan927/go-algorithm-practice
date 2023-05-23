package leecode

import "errors"

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics 递归 链表 数学 👍 9607 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	elements []int
	size     int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]int, 0),
		size:     0,
	}
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}
	result := s.elements[s.size-1]
	return result, nil
}

func (s *Stack) Push(newVal int) {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//考虑用栈
	res := &ListNode{}
	sumList1 := make([]int, 0)
	sumList2 := make([]int, 0)
	for l1.Next != nil {
		sumList1 = append(sumList1, l1.Val)
	}
	for l2.Next != nil {
		sumList2 = append(sumList2, l2.Val)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
