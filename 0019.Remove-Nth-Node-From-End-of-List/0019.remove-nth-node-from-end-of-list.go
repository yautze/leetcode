package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// ListNode -
type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	if n <= 0 {
		return head
	}

	len := 0
	current := head
	// 計算head的總長度
	for current != nil {
		len++
		current = current.Next
	}

	if n > len {
		return head
	}

	// 如果相等則是移掉第一個數
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}

	// reset current to first node
	current = head

	i := 0
	for current != nil {
		// 移動到要移除的節點的前一個時要將當節點的next替換掉
		if len-n-1 == i {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
		}
		i++
		current = current.Next
	}

	return head
}
