package medium

import (
	. "project/leetcode/datastructures/list"
)

func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev, current := dummy, head

	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}

		if prev.Next == current {
			prev = prev.Next
		} else {
			prev.Next = current.Next
		}

		current = current.Next
	}

	return dummy.Next
}
