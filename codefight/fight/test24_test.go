package main

import "testing"

func BenchmarkDataRoute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		network := [][]int{
			{0, 2, 4, 8, 0, 0},
			{0, 0, 0, 9, 0, 0},
			{0, 0, 0, 0, 0, 10},
			{0, 0, 6, 0, 0, 10},
			{10, 10, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}
		if dataRoute(4, 5, network) != 19 {
			b.Fail()
		}
	}
}

func TestDataRoute(t *testing.T) {
	network := [][]int{
		{0, 2, 4, 8, 0, 0},
		{0, 0, 0, 9, 0, 0},
		{0, 0, 0, 0, 0, 10},
		{0, 0, 6, 0, 0, 10},
		{10, 10, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	if dataRoute(4, 5, network) != 19 {
		t.Fail()
	}
}
