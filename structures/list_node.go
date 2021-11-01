package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntArrToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}

	return l.Next
}

func LinkedListToIntArr(head *ListNode) []int {
	limit := 100
	times := 0

	res := []int{}

	for head != nil {
		times++
		if times > limit {
			panic("items in linked list exceed limit, unable to convert")
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
