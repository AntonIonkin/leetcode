package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	sortList := &ListNode{
		Val: 0,
	}
	curSortList := sortList
	currL1, currL2 := list1, list2
	for currL1 != nil || currL2 != nil {
		if currL1.Val < currL2.Val {
			curSortList.Next = &ListNode{
				Val: currL1.Val,
			}
			curSortList = curSortList.Next
			curSortList.Next = &ListNode{
				Val: currL2.Val,
			}
			curSortList = curSortList.Next
		} else {
			curSortList.Next = &ListNode{
				Val: currL2.Val,
			}
			curSortList = curSortList.Next

			curSortList.Next = &ListNode{
				Val: currL1.Val,
			}
			curSortList = curSortList.Next

		}
		currL1 = currL1.Next
		currL2 = currL2.Next

	}

	return sortList.Next
}

func main() {
	l1 := []int{3, 5, 9, 8}
	l2 := []int{1, 6, 9}
	list1 := CreateList(l1)
	list2 := CreateList(l2)
	fmt.Println(mergeTwoLists(list1, list2))
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
