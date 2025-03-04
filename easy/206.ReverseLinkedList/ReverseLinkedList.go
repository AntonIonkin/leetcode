package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil // Предыдущий узел
	current := head          // Текущий узел

	for current != nil { // Пока текущий узел не равен nil
		next := current.Next // Сохраняем следующий узел
		current.Next = prev  // Перенаправляем текущий узел на предыдущий
		prev = current       // Обновляем предыдущий узел на текущий
		current = next       // Переходим к следующему узлу
	}

	return prev // Вернуть новый голову перевернутого списка
}

func main() {
	a := CreateList(1)
	fmt.Println(a)
	reversed := reverseList(a)
	fmt.Println(reversed)
	fmt.Println(a)
}

func CreateList(n int) *ListNode {
	if n == 0 {
		return nil
	}

	head := &ListNode{
		Val: 1,
	}
	current := head

	for i := 0; i < n-1; i++ {
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
