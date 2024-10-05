package main

import "fmt"

type Piece int

const (
	EMPTY Piece = iota
	BLACK
	WHITE
)

type Board struct {
	Board [][]Piece
	Size  int
}

func (b Board) Set(i int, j int, p Piece) {
	b.Board[i][j] = p
}

func (b Board) Display() {
	for i := 0; i < b.Size; i++ {
		fmt.Println(b.Board[i])
	}
}

func main() {
	board_size := 10

	fmt.Println("Go Game in Go Lang")

	board := Board{
		Board: make([][]Piece, board_size),
		Size:  board_size,
	}

	for i := 0; i < board_size; i++ {
		board.Board[i] = make([]Piece, board_size)
	}

	board.Display()

	board.Set(0, 0, BLACK)
	fmt.Println("Set (0, 0) to BLACK")

	board.Display()
}
