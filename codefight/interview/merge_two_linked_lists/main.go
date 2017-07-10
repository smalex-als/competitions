package main

import (
	"fmt"
)

func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	b := &ListNode{3, &ListNode{5, &ListNode{6, nil}}}
	dumpList(mergeTwoLinkedLists(a, b))
}

// Definition for singly-linked list:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Value.(int) < l2.Value.(int)) {
			if tail != nil {
				tail.Next = l1
			}
			tail = l1
			l1 = l1.Next
		} else {
			if tail != nil {
				tail.Next = l2
			}
			tail = l2
			l2 = l2.Next
		}
		if head == nil {
			head = tail
		}
	}
	return head
}

func dumpList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v -> ", l.Value)
		l = l.Next
	}
	fmt.Println("nil")
}
