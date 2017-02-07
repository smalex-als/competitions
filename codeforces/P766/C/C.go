package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	str := readString()
	a := make([]int, 26)
	b := make([]int, 26)
	for i := 0; i < 26; i++ {
		a[i] = readInt()
	}
	r := []rune(str)
	h := NewMinHeap(26)
	start := 0
	end := 0
	maxlen := 0
	jumplen := make([]int, n)
	for end < n && start < n {
		ch := int(r[end] - 'a')
		b[ch]++
		end++
		h.Add(ch, a[ch])
		curlen := end - start
		for end-start > 0 && h.Min() < curlen {
			oldch := int(r[start] - 'a')
			b[oldch]--
			curlen--
			start++
			if b[oldch] == 0 {
				h.Remove(oldch)
			}
		}
		if maxlen < curlen {
			maxlen = curlen
		}
		for i := start; i < end; i++ {
			if jumplen[i] < end {
				jumplen[i] = end
			}
		}
	}
	res := dfs(str, 0, jumplen, 0)
	fmt.Printf("%d\n", res)
	fmt.Printf("%d\n", maxlen)
	fmt.Printf("%d\n", minJumps)
}

var memo = map[int]int{}
var minJumps = 100000

func dfs(str string, pos int, tmplen []int, jumps int) int {
	if prev, ok := memo[pos]; ok {
		return prev
	}
	if pos == len(str) {
		if minJumps > jumps {
			minJumps = jumps
		}
		return 1
	}
	res := 0
	for i := tmplen[pos]; i >= pos+1; i-- {
		if i <= len(str) {
			res += dfs(str, i, tmplen, jumps+1)
			res %= 1000000007
		}
	}
	memo[pos] = res
	return res
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
