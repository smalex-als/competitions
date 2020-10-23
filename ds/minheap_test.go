package ds

import (
	"math/rand"
	"sort"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	h := NewMinHeap(10)
	h.Add(0, 10)
	h.Add(1, 5)
	h.Add(2, 6)
	h.Add(3, 2)
	MinHeapAssertSize(t, h, 4)
	MinHeapAssertMin(t, h, 2)
	h.Remove(2)
	MinHeapAssertSize(t, h, 3)
	MinHeapAssertMin(t, h, 2)
	h.Remove(3)
	MinHeapAssertSize(t, h, 2)
	MinHeapAssertMin(t, h, 5)
}

func TestNewMinHeapReverse(t *testing.T) {
	h := NewMinHeap(10)
	h.Add(0, -10)
	h.Add(1, -5)
	h.Add(2, -6)
	h.Add(3, -2)
	MinHeapAssertSize(t, h, 4)
	MinHeapAssertMin(t, h, -10)
	h.Remove(2)
	MinHeapAssertSize(t, h, 3)
	MinHeapAssertMin(t, h, -10)
	h.Remove(0)
	MinHeapAssertSize(t, h, 2)
	MinHeapAssertMin(t, h, -5)
}

func TestNewMinHeapRand(t *testing.T) {
	n := 10000
	h := NewMinHeap(n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		value := rand.Intn(n)
		if rand.Intn(2) == 0 {
			value = -value
		}
		values[i] = value
		h.Add(i, value)
	}
	cntDelete := n / 4
	deleted := make([]bool, n)
	for i := 0; i < cntDelete; i++ {
		value := rand.Intn(n)
		if deleted[value] {
			i--
			continue
		}
		deleted[value] = true
		h.Remove(value)
	}
	MinHeapAssertSize(t, h, n-cntDelete)
	values2 := make([]int, 0)
	for i := 0; i < n; i++ {
		if !deleted[i] {
			values2 = append(values2, values[i])
		}
	}
	sort.Ints(values2)
	for i := 0; i < len(values2); i++ {
		value := h.Min()
		h.Remove(h.ArgMin())
		if value != values2[i] {
			t.Fatalf("%+v: expected value %+v but was %+v", i, values2[i], value)
		}
	}
	MinHeapAssertSize(t, h, 0)
}

func MinHeapAssertSize(t *testing.T, h *MinHeap, expected int) {
	sz := h.Size()
	if sz != expected {
		t.Fatalf("Expected size of heap %+v, but was %+v", expected, sz)
	}
}

func MinHeapAssertMin(t *testing.T, h *MinHeap, expected int) {
	m := h.Min()
	if m != expected {
		t.Fatalf("Expected min %+v, but was %+v", expected, m)
	}
}

func MinHeapAssertArgFirst(t *testing.T, h *MinHeap, expected int) {
	m := h.ArgMin()
	if m != expected {
		t.Fatalf("Expected arg max %+v, but was %+v", expected, m)
	}
}
