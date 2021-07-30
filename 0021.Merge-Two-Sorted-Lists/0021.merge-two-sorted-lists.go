package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// ListNode -
type ListNode = structures.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	// 操作用
	current := res

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 補足剩下的那一邊
	if l1 == nil {
		current.Next = l2
	}
	if l2 == nil {
		current.Next = l1
	}

	return res.Next
}

// 遞迴的做法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		// 小的那個往前一個位置,跑下一次比較
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
