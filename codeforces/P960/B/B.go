package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type AB struct {
	a int
	b int
}

func solve() {
	n := readInt()
	k1 := readInt()
	k2 := readInt()
	a := readIntArray(n)
	b := readIntArray(n)
	heap := NewMinHeap(n)
	for i := 0; i < n; i++ {
		heap.Add(i, -abs(a[i]-b[i]))
	}
	for k1 > 0 || k2 > 0 {
		ind := heap.ArgMin()
		f := a[ind]
		v := b[ind]
		if f > v {
			if k1 > 0 {
				f--
				k1--
			} else if k2 > 0 {
				v++
				k2--
			}
		} else if v > f {
			if k1 > 0 {
				f++
				k1--
			} else if k2 > 0 {
				v--
				k2--
			}
		} else {
			if k1 > 0 {
				f++
				k1--
			} else if k2 > 0 {
				v--
				k2--
			}
		}
		a[ind] = f
		b[ind] = v

		heap.Update(ind, -abs(f-v))
	}
	total := int64(0)
	for i := 0; i < n; i++ {
		total += int64(a[i]-b[i]) * int64(a[i]-b[i])
	}
	fmt.Println(total)
}

type MinHeap struct {
	a    []int
	vmap []int
	imap []int
	n    int
	pos  int
}

const INF = math.MaxInt32

func NewMinHeap(m int) *MinHeap {
	heap := &MinHeap{}
	m += 2
	heap.n = m
	heap.a = make([]int, heap.n)
	heap.vmap = make([]int, heap.n)
	heap.imap = make([]int, heap.n)
	heap.pos = 1
	for i := 0; i < heap.n; i++ {
		heap.a[i] = INF
		heap.vmap[i] = -1
		heap.imap[i] = -1
	}
	return heap
}

func (h *MinHeap) Add(ind, x int) int {
	ret := h.imap[ind]
	if h.imap[ind] < 0 {
		h.a[h.pos] = x
		h.vmap[h.pos] = ind
		h.imap[ind] = h.pos
		h.pos++
		h.up(h.pos - 1)
	}
	if ret != -1 {
		return h.a[ret]
	}
	return x
}

func (h *MinHeap) Update(ind, x int) int {
	ret := h.imap[ind]
	if h.imap[ind] < 0 {
		h.a[h.pos] = x
		h.vmap[h.pos] = ind
		h.imap[ind] = h.pos
		h.pos++
		h.up(h.pos - 1)
	} else {
		h.a[ret] = x
		h.up(ret)
		h.down(ret)
	}
	return x
}

func (h *MinHeap) Remove(ind int) int {
	if h.pos == 1 {
		return INF
	}
	if h.imap[ind] == -1 {
		return INF
	}

	h.pos--
	rem := h.imap[ind]
	ret := h.a[rem]
	h.vmap[rem] = h.vmap[h.pos]
	h.imap[h.vmap[h.pos]] = rem
	h.imap[ind] = -1
	h.a[rem] = h.a[h.pos]
	h.a[h.pos] = INF
	h.vmap[h.pos] = -1

	h.up(rem)
	h.down(rem)
	return ret
}

func (h *MinHeap) Min() int {
	return h.a[1]
}

func (h *MinHeap) ArgMin() int {
	return h.vmap[1]
}

func (h *MinHeap) Size() int {
	return h.pos - 1
}

func (h *MinHeap) up(cur int) {
	for c, p := cur, cur>>1; p >= 1 && h.a[p] > h.a[c]; c, p = c>>1, p>>1 {
		h.a[p], h.a[c] = h.a[c], h.a[p]
		h.imap[h.vmap[p]], h.imap[h.vmap[c]] = h.imap[h.vmap[c]], h.imap[h.vmap[p]]
		h.vmap[p], h.vmap[c] = h.vmap[c], h.vmap[p]
	}
}

func (h *MinHeap) down(cur int) {
	for c := cur; 2*c < h.pos; {
		b := 0
		if h.a[2*c] < h.a[2*c+1] {
			b = 2 * c
		} else {
			b = 2*c + 1
		}
		if h.a[b] < h.a[c] {
			h.a[c], h.a[b] = h.a[b], h.a[c]
			h.imap[h.vmap[c]], h.imap[h.vmap[b]] = h.imap[h.vmap[b]], h.imap[h.vmap[c]]
			h.vmap[c], h.vmap[b] = h.vmap[b], h.vmap[c]
			c = b
		} else {
			break
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

// IO

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readInt64() int64 {
	v, _ := strconv.ParseInt(readString(), 10, 64)
	return v
}

func readIntArray(sz int) []int {
	a := make([]int, sz)
	for i := 0; i < sz; i++ {
		a[i] = readInt()
	}
	return a
}

func readInt64Array(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = readInt64()
	}
	return res
}
