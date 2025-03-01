package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var intersection *ListNode
	var lengthA, lengthB int
	currA, currB := headA, headB

	for currA != nil {
		lengthA++
		currA = currA.Next
	}

	for currB != nil {
		lengthB++
		currB = currB.Next
	}

	currA, currB = headA, headB
	if lengthA > lengthB {

		for i := 0; i < lengthA-lengthB; i++ {
			currA = currA.Next
		}
	} else {
		for i := 0; i < lengthB-lengthA; i++ {
			currB = currB.Next
		}
	}

	for currA != nil || currB != nil {
		if currA.Val == currB.Val {
			intersection = currA.Next
			break
		}
		currA = headA.Next
		currB = currB.Next
	}

	return intersection
}

func main() {
	l1 := []int{4, 1, 8, 4, 5}
	l2 := []int{5, 6, 1, 8, 4, 5}
	list1 := CreateList(l1)
	list2 := CreateList(l2)
	fmt.Println(getIntersectionNode(list1, list2))
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
