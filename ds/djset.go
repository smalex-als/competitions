package ds

type DJSet struct {
	upper []int
}

func NewDJSet(n int) *DJSet {
	s := DJSet{upper: make([]int, n)}
	for i := 0; i < n; i++ {
		s.upper[i] = -1
	}
	return &s
}

func (s *DJSet) Root(x int) int {
	if s.upper[x] < 0 {
		return x
	} else {
		s.upper[x] = s.Root(s.upper[x])
	}
	return s.upper[x]
}

func (s *DJSet) Union(x, y int) bool {
	x = s.Root(x)
	y = s.Root(y)
	if x != y {
		if s.upper[y] < s.upper[x] {
			x, y = y, x
		}
		s.upper[x] += s.upper[y]
		s.upper[y] = x
	}
	return x == y
}

func (s *DJSet) Equiv(x, y int) bool {
	return s.Root(x) == s.Root(y)
}

func (s *DJSet) Upper(y int) int {
	return s.upper[y]
}

func (s *DJSet) Count(root int) int {
	return -s.upper[root]
}

func (s *DJSet) Components() int {
	ct := 0
	for _, u := range s.upper {
		if u < 0 {
			ct++
		}
	}
	return ct
}
