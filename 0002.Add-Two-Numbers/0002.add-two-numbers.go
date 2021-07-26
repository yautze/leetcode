package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode -
type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Val: 0,
	}
	// carry - 進位
	n1, n2, carry := 0, 0, 0
	// 操作用
	current := res

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: (n1 + n2 + carry) % 10,
		}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return res.Next
}
