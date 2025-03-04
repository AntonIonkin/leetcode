package medium

import (
	. "project/leetcode/datastructures/list"
)

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {

		l := head
		r := head.Next
		l.Next = r.Next
		r.Next = l
		prev.Next = r
		prev = l
		head = l.Next
	}

	return dummy.Next 
}
