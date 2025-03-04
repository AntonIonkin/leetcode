package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var length int

	curr := head
	for curr != nil {
		length++
		curr = curr.Next

	}

	k = k % length

	var startTail *ListNode
	curr = head
	for i := 0; i < k; i++ {
		curr = curr.Next
	}

	startTail = curr.Next
	curr.Next = nil

	root := startTail
	for root.Next != nil {
		root = root.Next
	}

	root.Next = head

	return root
}

func main() {
	a := CreateList(5)
	fmt.Println(a)
	fmt.Println(rotateRight(a, 2))

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
