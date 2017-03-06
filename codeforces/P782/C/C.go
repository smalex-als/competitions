package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

var res []int

type Vertex struct {
	color int
	adj   []*Vertex
}

func (v *Vertex) add(u *Vertex) {
	v.adj = append(v.adj, u)
}

func (v *Vertex) dfs(us int, skip int) {
	v.color = us
	nxt := 1
	for nxt == us || nxt == skip {
		nxt++
	}
	for _, next := range v.adj {
		if next.color == 0 {
			next.dfs(nxt, us)
			nxt++
			for nxt == us || nxt == skip {
				nxt++
			}
		}
	}
}

func solve() {
	n := readInt()
	vertexes := make([]*Vertex, n)
	for i := 0; i < n; i++ {
		vertexes[i] = &Vertex{0, make([]*Vertex, 0)}
	}
	for i := 0; i < n-1; i++ {
		x := readInt() - 1
		y := readInt() - 1
		vertexes[y].add(vertexes[x])
		vertexes[x].add(vertexes[y])
	}
	vertexes[0].dfs(1, 0)
	max := 0
	for _, v := range vertexes {
		if max < len(v.adj)+1 {
			max = len(v.adj) + 1
		}
	}
	fmt.Println(max)
	var buffer bytes.Buffer
	for i, v := range vertexes {
		if i > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(strconv.Itoa(v.color))
	}
	fmt.Println(buffer.String())
}

var scanner *bufio.Scanner

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
