package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode = nil
	current := slow
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	left := head
	right := prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func main() {
	a := CreateList(5)
	fmt.Println(a)
	fmt.Println(isPalindrome(a))
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
