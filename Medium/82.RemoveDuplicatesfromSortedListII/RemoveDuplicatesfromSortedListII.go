package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	curr := dummy
	duplicate := curr.Next.Val

	for curr.Next != nil {
		if duplicate != curr.Val { // duplicate != curr.Next.Val
			duplicate = curr.Val
		}

		if curr.Next.Val == duplicate {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

func main() {
	aSlice := []int{1, 1, 2, 3, 3}
	a := CreateList(aSlice)
	fmt.Println(a)
	fmt.Println(deleteDuplicates(a))

}

func CreateList(n []int) *ListNode {
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
