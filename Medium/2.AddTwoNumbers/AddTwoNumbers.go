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
	var lengthL1, lengthL2 int
	var prevL1, prevL2 *ListNode
	currL1, currL2 := l1, l2

	for currL1 != nil {
		lengthL1++
		next := currL1.Next
		currL1.Next = prevL1
		prevL1 = currL1
		currL1 = next
	}

	for currL2 != nil {
		lengthL2++
		next := currL2.Next
		currL2.Next = prevL2
		prevL2 = currL2
		currL2 = next
	}

	maxL, minL := prevL1, prevL2
	maxLength, minLength := lengthL1, lengthL2

	if lengthL1 < lengthL2 {
		maxL, minL = minL, maxL
		maxLength, minLength = minLength, maxLength
	}

	remainder := 0
	for i := 0; i <= maxLength; i++ {
		sum := maxL.Val
		if minLength >= 0 {
			sum += minL.Val
			minL = minL.Next
			minLength--
		}

		if sum > 9 {
			sum, remainder = sum%10, 1
		}

		maxL.Val = sum
		if maxL.Next == nil && remainder == 1 {
			maxL.Next = &ListNode{
				Val: 1,
			}
			remainder = 0
		}

		if maxL.Next != nil && remainder == 1 {
			maxL.Next = &ListNode{
				Val: maxL.Val + 1,
			}
			remainder = 0
		}
		maxL = maxL.Next
	}

	return maxL
}

func main() {
	l1Slice := []int{2, 4, 3}
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
