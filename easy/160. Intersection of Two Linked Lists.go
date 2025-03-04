package easy

import . "project/leetcode/datastructures/list"

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB
	var common *ListNode
	for curA != nil {
		curA.Val -= 1_000_000
		curA = curA.Next
	}

	for curB != nil {
		if curB.Val < 0 {
			common = curB
			break
		}
		curB = curB.Next
	}

	curA = headA
	for curA != nil {
		curA.Val += 1_000_000
		curA = curA.Next
	}

	return common
}
