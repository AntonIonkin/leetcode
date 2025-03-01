package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	a := CreateList(6)
	fmt.Println(a)
	b := middleNode(a)
	fmt.Println(b)
}

func CreateList(n int) *ListNode {
	if n == 0 {
		return nil
	}

	head := &ListNode{
		Val: 1,
	}

	cur := head

	for i := 1; i < n; i++ {
		cur.Next = &ListNode{
			Val: cur.Val + 1,
		}
		cur = cur.Next
	}

	return head
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}

	var result string

	for l != nil {
		result += strconv.Itoa(l.Val) + "->"
		l = l.Next
	}

	result += "nil"

	return result
}
