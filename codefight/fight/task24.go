package main

import "fmt"

func main24() {
	network := [][]int{
		{0, 2, 4, 8, 0, 0},
		{0, 0, 0, 9, 0, 0},
		{0, 0, 0, 0, 0, 10},
		{0, 0, 6, 0, 0, 10},
		{10, 10, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	res := dataRoute(4, 5, network)
	fmt.Printf("res = %+v\n", res)
}

type FlowEdge struct {
	v, w, capacity, flow int
}

func (e *FlowEdge) other(x int) int {
	if x == e.v {
		return e.w
	}
	if x == e.w {
		return e.v
	}
	panic("unknow edge")
}

func (e *FlowEdge) residualCapacityTo(x int) int {
	if e.v == x {
		return e.flow
	} else if e.w == x {
		return e.capacity - e.flow
	} else {
		panic(fmt.Sprintf("unknown edge %d %+v", x, e))
	}
}

func (e *FlowEdge) addResidualCapacityTo(x int, delta int) {
	if e.v == x {
		e.flow -= delta
	} else if e.w == x {
		e.flow += delta
	} else {
		panic("unknown edge")
	}
}

type FlowNetwork struct {
	edges [][]*FlowEdge
}

func NewFlowNetwork(n int) *FlowNetwork {
	network := FlowNetwork{}
	network.edges = make([][]*FlowEdge, n)
	for i := 0; i < n; i++ {
		network.edges[i] = make([]*FlowEdge, 0)
	}
	return &network
}

func (n *FlowNetwork) addEdge(edge *FlowEdge) {
	n.edges[edge.v] = append(n.edges[edge.v], edge)
}

func (n *FlowNetwork) adj(v int) []*FlowEdge {
	return n.edges[v]
}

func (n *FlowNetwork) V() int {
	return len(n.edges)
}

type FordFulkerson struct {
	edgeTo []*FlowEdge
}

func NewFordFulkerson() *FordFulkerson {
	return &FordFulkerson{}
}

func (f *FordFulkerson) MaxFlow(g *FlowNetwork, s, t int) int {
	value := 0
	cnt := 0
	for f.hasAugmentingPath(g, s, t) {
		// fmt.Printf("loop = %+v\n", cnt)
		bottle := 100000000
		for v := t; v != s; v = f.edgeTo[v].other(v) {
			// ed := f.edgeTo[v]
			// fmt.Printf("  %d -> %d (cap: %d, flow: %d)\n", ed.v, ed.w, ed.capacity, ed.flow)
			bottle = flowMin(bottle, f.edgeTo[v].residualCapacityTo(v))
		}
		for v := t; v != s; v = f.edgeTo[v].other(v) {
			f.edgeTo[v].addResidualCapacityTo(v, bottle)
		}
		value += bottle
		cnt++
	}
	return value
}

func (f *FordFulkerson) hasAugmentingPath(g *FlowNetwork, s, t int) bool {
	f.edgeTo = make([]*FlowEdge, g.V())

	marked := make([]bool, g.V())
	q := make([]int, 1)
	marked[s] = true
	q[0] = s
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		children := g.adj(v)
		for i := 0; i < len(children); i++ {
			e := children[i]
			w := e.other(v)
			if e.residualCapacityTo(w) > 0 && !marked[w] {
				f.edgeTo[w] = e
				marked[w] = true
				q = append(q, w)
			}
		}
	}
	return marked[t]
}

func dataRoute(resource int, server int, network [][]int) int {
	flowNetwork := NewFlowNetwork(len(network))
	for i := 0; i < len(network); i++ {
		for j := 0; j < len(network); j++ {
			if network[i][j] > 0 {
				flowNetwork.addEdge(&FlowEdge{i, j, network[i][j], 0})
			}
		}
	}
	fordFulkerson := NewFordFulkerson()
	return fordFulkerson.MaxFlow(flowNetwork, resource, server)
}

func flowMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
