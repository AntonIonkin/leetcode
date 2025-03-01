package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	a := antonListNode(5)

	a.removeElements(3)

	fmt.Println(a.Print())

}

func antonListNode(n int) *ListNode {
	if n == 0 {
		return nil
	}

	var head, current *ListNode

	head = &ListNode{
		Val: 1,
	}

	current = head

	for i := 1; i < n; i++ {
		current.Next = &ListNode{
			Val: current.Val + 1,
		}

		current = current.Next
	}

	return head

}

func (l *ListNode) Print() string {
	if l == nil {
		return ""
	}
	var result string

	for l != nil {
		result += strconv.Itoa(l.Val) + " "
		l = l.Next
	}

	result += "nil"

	return result
}

func (l *ListNode) removeElements(num int) *ListNode {
	if l == nil {
		return nil // Если список пуст
	}

	// Общий случай: если нужно удалять узлы
	var prev *ListNode = nil
	head := l // Сохраняем начальную ссылку на голову списка

	for l != nil {
		if l.Val == num {
			if prev == nil {
				// Если это голова, перемещаем голову списка
				head = l.Next
			} else {
				// Если не голова, просто переходим к следующему узлу
				prev.Next = l.Next
			}
		} else {
			// Если значение не совпадает, обновляем prev
			prev = l
		}
		// Переходим к следующему узлу
		l = l.Next
	}

	return head // Возвращаем новую голову списка
}
