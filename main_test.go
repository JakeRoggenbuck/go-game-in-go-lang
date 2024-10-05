package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	b := New(19)

	j := 0

	for i := range 10 {

		if b.Board[i][j] != EMPTY {
			t.Fatalf("Board(%d, %d) != EMPTY", i, j)
		}

		b.Set(i, j, BLACK)

		if b.Board[i][j] != BLACK {
			t.Fatalf("Board(%d, %d) != BLACK", i, j)
		}

	}
}

func TestGet(t *testing.T) {
	b := New(19)

	j := 2

	for i := range 10 {
		if b.Board[i][j] != EMPTY {
			t.Fatalf("Board(%d, %d) != EMPTY", i, j)
		}

		b.Set(i, j, BLACK)

		if b.Board[i][j] != BLACK || b.Get(i, j) != BLACK {
			t.Fatalf("Board(%d, %d) != BLACK", i, j)
		}
	}
}
