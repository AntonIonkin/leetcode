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
	current := head
	count, num := 0, 0

	for current != nil || current.Next != nil {
		if current.Val != num {
			num = current.Val
			count = 0
		}
		if current.Val == num {
			count++
		}
		if current.Next != nil && current.Next.Val != num {
			dummy.Next, dummy = current, current
		}

		current = current.Next
	}
	return dummy
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
