package medium

import (
	. "project/leetcode/datastructures/list"
)

func ReorderList(head *ListNode) {
	dummy := &ListNode{Next: head}
	slow, fast, cur := head, head, dummy

	for fast != nil && fast.Next != nil {
		cur = cur.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur.Next = nil

	var reverseSlow *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = reverseSlow
		reverseSlow = slow
		slow = next
	}

	first, second := head, reverseSlow

	for second != nil || first != nil {

		nextFirst := first.Next
		nextSecond := second.Next

		first.Next = second
		second.Next = nextFirst
		if nextFirst == nil && nextSecond != nil {
			first.Next.Next = nextSecond
			break
		}
		first = nextFirst
		second = nextSecond
	}
}
