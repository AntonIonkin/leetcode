package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 != nil {
		return l2
	}

	var remainder int
	res := &ListNode{}
	root := res

	for l1 != nil && l2 != nil {
		res.Next, remainder = &ListNode{
			Val: (l1.Val + l2.Val + remainder) % 10,
		}, (l1.Val+l2.Val+remainder)/10
		res = res.Next
		l1, l2 = l1.Next, l2.Next
	}

	if l2 != nil {
		l1 = l2
	}

	for l1 != nil {
		res.Next, remainder = &ListNode{
			Val: (l1.Val + remainder) % 10,
		}, (l1.Val+remainder)/10
		res = res.Next
		l1 = l1.Next
	}

	if remainder != 0 {
		res.Next = &ListNode{
			Val: remainder,
		}
	}

	return root.Next
}

func main() {
	l1Slice := []int{9, 4, 3}
	l2Slice := []int{5, 6, 4}
	l1 := CreateList(l1Slice)
	l2 := CreateList(l2Slice)

	fmt.Println(addTwoNumbers(l1, l2))

}

func CreateList(n []int) *ListNode {
	if len(n) == 0 {
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
