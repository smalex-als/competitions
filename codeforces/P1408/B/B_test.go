package main

import (
	"testing"

	"github.com/smalex-als/competitions/test"
)

func TestA(t *testing.T) {
	test.RunTest("B", t)
}

func TestFunc(t *testing.T) {
	res := solveSingle([]int{0, 1, 1}, 2)
	if 1 != res {
		t.Fatalf("expected %d, but got %d", 1, res)
	}
	res = solveSingle([]int{0, 1, 2}, 2)
	if 2 != res {
		t.Fatalf("expected %d, but got %d", 2, res)
	}
	res = solveSingle([]int{0, 1, 2, 3}, 3)
	if 2 != res {
		t.Fatalf("expected %d, but got %d", 2, res)
	}
}
