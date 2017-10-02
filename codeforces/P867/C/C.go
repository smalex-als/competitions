package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var n int
var s int64
var a [][]int64

func solve() {
	n = readInt()
	s = readInt64()

	a = make([][]int64, n)
	totalReq := int64(0)
	for i := 0; i < n; i++ {
		a[i] = make([]int64, 3)
		a[i][0] = readInt64()
		a[i][1] = readInt64()
		a[i][2] = readInt64()
		totalReq += int64(a[i][0])
	}
	pizzasReq := (totalReq + int64(s-1)) / int64(s)
	if pizzasReq == 1 {
		fmt.Println(max(getValue(0, pizzasReq), getValue(1, pizzasReq)))
		return
	}
	lo := int64(0)
	hi := int64(pizzasReq) + 1
	for hi-lo > 2 {
		mid := (lo + hi) / 2
		nl := lo + (hi-lo)/3
		nh := hi + (lo-hi)/3
		value1 := getValue((lo+mid)/2, pizzasReq)
		value2 := getValue((mid+hi)/2, pizzasReq)
		if value1 > value2 {
			hi = nh
		} else {
			lo = nl
		}
	}
	best := int64(0)
	for i := max(0, lo-3); i <= min(hi+3, pizzasReq); i++ {
		best = max(best, getValue(i, pizzasReq))
	}

	fmt.Println(best)
}

func getValue(i, pizzasReq int64) int64 {
	cntA := i * s
	cntB := (pizzasReq - i) * s
	heap := NewMinHeap(n * 2)
	cnt := make([]int64, n)
	for j := 0; j < n; j++ {
		heap.Add(j, -a[j][0]*a[j][1])
		heap.Add(j+n, -a[j][0]*a[j][2])
		cnt[j] = a[j][0]
	}

	res := int64(0)
	for heap.Size() > 0 {
		index := heap.ArgMin()
		heap.Remove(index)
		realIndex := index % n
		if index < n {
			h := min(cnt[realIndex], cntA)
			cnt[realIndex] -= h
			cntA -= h
			res += h * a[realIndex][1]
			if cnt[realIndex] == 0 {
				heap.Remove(index + n)
			}
		} else {
			h := min(cnt[realIndex], cntB)
			cnt[realIndex] -= h
			cntB -= h
			res += h * a[realIndex][2]
			if cnt[realIndex] == 0 {
				heap.Remove(index - n)
			}
		}
	}
	return res
}

type MinHeap struct {
	a    []int64
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
	heap.a = make([]int64, heap.n)
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

func (h *MinHeap) Add(ind int, x int64) int64 {
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

func (h *MinHeap) Update(ind int, x int64) int64 {
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

func (h *MinHeap) Remove(ind int) int64 {
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

func (h *MinHeap) Min() int64 {
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

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
