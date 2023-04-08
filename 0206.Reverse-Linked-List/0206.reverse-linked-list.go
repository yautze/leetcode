package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode -
type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
    var revNode *ListNode

    for head != nil {
        nextNode := head.Next
        // break old next
        head.Next = revNode
        revNode = head
        head = nextNode
    }

    return revNode
}
