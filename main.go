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

func (b Board) Get(i int, j int) Piece {
	return b.Board[i][j]
}

func (b Board) CanPlace(i int, j int) bool {
	return b.Get(i, j) == EMPTY
}

func (b Board) Unset(i int, j int) {
	b.Board[i][j] = EMPTY;
}

func (b Board) Display() {
	for i := 0; i < b.Size; i++ {
		fmt.Println(b.Board[i])
	}
}

func New(size int) Board {
	board := Board{
		Board: make([][]Piece, size),
		Size:  size,
	}

	for i := 0; i < size; i++ {
		board.Board[i] = make([]Piece, size)
	}

	return board
}

func main() {
	// Standard go board size
	board_size := 19
	board := New(board_size)

	fmt.Println("Go Game in Go Lang")

	board.Display()

	board.Set(0, 0, BLACK)
	fmt.Println("Set (0, 0) to BLACK")

	board.Display()
}
