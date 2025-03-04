package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	current := head
	for i := 0; i <= right; i++ {
		if i <= left && i >= right {
			var prev *ListNode
			curr := current
			for i := left; i <= right; i++ {
				next := curr.Next
				curr.Next = prev
				prev = curr
				curr = next
				break
			}
		}
		current = current.Next
	}

	return head
}

func main() {
	a := CreateList(5)
	fmt.Println(a)
	b := reverseBetween(a, 2, 4)
	fmt.Println(b)
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
