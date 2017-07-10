package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	// a := &ListNode{
	// 	Value: 9876,
	// 	Next: &ListNode{
	// 		Value: 5432,
	// 		Next: &ListNode{
	// 			Value: 1999,
	// 		},
	// 	},
	// }
	// b := &ListNode{
	// 	Value: 1,
	// 	Next: &ListNode{
	// 		Value: 8001,
	// 	},
	// }
	a := &ListNode{
		Value: 1,
	}
	b := &ListNode{
		Value: 9999,
		Next: &ListNode{
			Value: 9999,
			Next: &ListNode{
				Value: 9999,
			},
		},
	}
	dumpList(addTwoHugeNumbers(a, b))
}

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	lenA := ListLen(a)
	lenB := ListLen(b)
	if lenA < lenB {
		a, b = b, a
	}
	a = reverse(a)
	b = reverse(b)
	res := a
	cur := a
	rest := 0
	var last *ListNode
	for cur != nil {
		newvalue := rest + (cur.Value.(int))
		if b != nil {
			newvalue += (b.Value.(int))
			b = b.Next
		}
		rest = newvalue / 10000
		cur.Value = newvalue % 10000
		last = cur
		cur = cur.Next
	}
	if rest > 0 {
		last.Next = &ListNode{Value: rest}
	}
	res = reverse(res)
	return res
}

func ListLen(a *ListNode) int {
	cnt := 0
	for a != nil {
		a = a.Next
		cnt++
	}
	return cnt
}

func dumpList(a *ListNode) {
	for a != nil {
		fmt.Printf("%d -> ", a.Value)
		a = a.Next
		if a == nil {
			fmt.Println("nil")
		}
	}
}

func reverse(a *ListNode) *ListNode {
	head := a
	cur := head
	for cur.Next != nil {
		toRemove := cur.Next
		cur.Next = cur.Next.Next
		toRemove.Next = head
		head = toRemove
	}
	return head
}
