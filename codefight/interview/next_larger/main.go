package main

import "fmt"

func main() {
	// fmt.Println(nextLarger([]int{6, 2, 7, 3, 1, 0, 4, 5}))
	// [7, 7, -1, 4, 4, 4, 5, -1]
	fmt.Println(nextLarger([]int{10, 3, 12, 4, 2, 9, 13, 0, 8, 11, 1, 7, 5, 6}))
	// [12, 12, 13, 9, 9, 13, -1, 8, 11, -1, 7, -1, 6, -1]
}

type Stack struct {
	st []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (t *Stack) Peek() int {
	return t.st[len(t.st)-1]
}

func (t *Stack) Pop() int {
	ret := t.st[len(t.st)-1]
	t.st = t.st[:len(t.st)-1]
	return ret
}

func (t *Stack) Push(el int) {
	t.st = append(t.st, el)
}

func (t *Stack) Empty() bool {
	return len(t.st) == 0
}

func nextLarger(a []int) []int {
	stack := NewStack()
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = -1
	}
	for i := 0; i < len(a); i++ {
		for !stack.Empty() && a[stack.Peek()] < a[i] {
			res[stack.Pop()] = a[i]
		}
		stack.Push(i)
	}
	return res
}
