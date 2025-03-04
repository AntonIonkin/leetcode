package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElementsOther(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	root := head

	for head != nil {
		if head.Val == val {
			if prev == nil {
				root = head.Next
			} else {
				prev.Next = head.Next
			}
		} else {
			prev = head
		}
		head = head.Next
	}

	return root
}

func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}

func main() {
	a := CreateList(5)
	fmt.Println(a)
	a = removeElementsOther(a, 3)
	fmt.Println(a)
	a = removeElements(a, 1)
	fmt.Println(a)
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
