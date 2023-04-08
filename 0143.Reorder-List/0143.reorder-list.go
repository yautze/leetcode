package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode -
type ListNode = structures.ListNode

func reorderList(head *ListNode){
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    firstHalf := head
    secondHalf := reverseList(slow)

    for firstHalf != nil && secondHalf != nil {
        tmp := firstHalf.Next
        firstHalf.Next = secondHalf
        firstHalf = tmp
        tmp = secondHalf.Next
        secondHalf.Next = firstHalf
        secondHalf = tmp
    }
}

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
