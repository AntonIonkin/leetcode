package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEndOher(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	length -= n
	curr := dummy
	for i := 0; i <= length; i++ {
		if i == length {
			curr.Next = curr.Next.Next
			break
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fisrt, second := dummy, dummy

	for i := 0; i <= n; i++ {
		fisrt = fisrt.Next
	}

	for fisrt != nil {
		fisrt = fisrt.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	a := CreateList(5)
	fmt.Println(a)
	//b := removeNthFromEnd(a, 3)
	c := removeNthFromEndOher(a, 3)
	//fmt.Println(b)
	fmt.Println(c)

}

func CreateList(n int) *ListNode {
	if n == 0 {
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
