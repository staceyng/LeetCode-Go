package leetcode

import (
	"github.com/staceyng/LeetCode-Go/structures"
)

type ListNode = structures.ListNode

func middleNode(head *ListNode) *ListNode {
	// use dummy node as ref, fast pointer move 2 and slow pointer move 1
	dummy := new(ListNode)
	dummy.Next = head
	slow := dummy
	fast := dummy.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Next

}
