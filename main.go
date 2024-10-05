package main

import "fmt"

type Piece int

const (
	EMPTY Piece = iota
	BLACk
	WHITE
)

type Board struct {
	Board [][]Piece
}

func (b Board) set(i int, j int, p Piece) {
	b.Board[i][j] = p
}

func main() {
	board_size := 10

	fmt.Println("Go Game in Go Lang")

	board := Board{Board: make([][]Piece, board_size)}
	for i := 0; i < board_size; i++ {
		board.Board[i] = make([]Piece, board_size)
		fmt.Println(board.Board[i])
	}
}
