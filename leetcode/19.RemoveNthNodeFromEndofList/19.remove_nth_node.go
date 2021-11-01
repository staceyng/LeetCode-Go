package leetcode

import (
	"github.com/staceyng/LeetCode-Go/structures"
)

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	left := dummy
	right := head

	// move right pointer n steps from head
	for i := 0; i < n; i++ {
		right = right.Next
	}

	// find node to remove
	for right != nil {
		left = left.Next
		right = right.Next
	}

	// change node at left pointer to point to next.next
	left.Next = left.Next.Next

	return dummy.Next

}
