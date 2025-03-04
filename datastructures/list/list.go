package list

import "strconv"

//import . "project/leetcode/datastructures/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(n int) *ListNode {
	if n <= 0 {
		return nil
	}

	head := &ListNode{
		Val: 1,
	}

	current := head

	for i := 1; i < n; i++ {
		current.Next = &ListNode{
			Val: current.Val + 1,
		}
		current = current.Next
	}

	return head
}

func CreateListSlice(n []int) *ListNode {
	if len(n) <= 0 {
		return nil
	}

	head := &ListNode{
		Val: n[0],
	}

	current := head

	for i := 1; i < len(n); i++ {
		current.Next = &ListNode{
			Val: n[i],
		}
		current = current.Next
	}

	return head
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}

	result := ""

	for l != nil {
		result += strconv.Itoa(l.Val) + "->"
		l = l.Next
	}

	result += "nil"

	return result
}
