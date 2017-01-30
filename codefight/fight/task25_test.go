package main

import "testing"

func BenchmarkDataRouteShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		network := [][]int{
			{0, 2, 4, 8, 0, 0},
			{0, 0, 0, 9, 0, 0},
			{0, 0, 0, 0, 0, 10},
			{0, 0, 6, 0, 0, 10},
			{10, 10, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}
		if dataRouteShort(4, 5, network) != 19 {
			b.Fail()
		}
	}
}

func TestDataRouteShort(t *testing.T) {
	network := [][]int{
		{0, 2, 4, 8, 0, 0},
		{0, 0, 0, 9, 0, 0},
		{0, 0, 0, 0, 0, 10},
		{0, 0, 6, 0, 0, 10},
		{10, 10, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	if dataRouteShort(4, 5, network) != 19 {
		t.Fail()
	}
}
