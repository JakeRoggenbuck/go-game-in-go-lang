package main

import (
	"fmt"
	"os"
)

type Piece int

const (
	EMPTY Piece = iota
	BLACK
	WHITE
)

func (p Piece) Name() string {
	strings := [...]string{"EMPTY", "Black", "White"}

	if p < EMPTY || p > WHITE {
		return "Unknown"
	}

	return strings[p]
}

func (p Piece) String() string {
	strings := [...]string{"+ ", "⚪", "⚫"}


	if p < EMPTY || p > WHITE {
		return "?"
	}

	return strings[p]
}

func game_won(winner int) {
	fmt.Printf("Player %d wins!\n", winner)
	os.Exit(0)
}

type Board struct {
	Board [][]Piece
	Size  int
}

func (b Board) Set(i int, j int, p Piece) {
	if i < b.Size && j < b.Size {
		b.Board[i][j] = p
	}
}

func (b Board) Get(i int, j int) Piece {
	if i < b.Size && j < b.Size {
		return b.Board[i][j]
	}

	return EMPTY
}

func (b Board) CanPlace(i int, j int) bool {
	return b.Get(i, j) == EMPTY
}

func (b Board) Unset(i int, j int) {
	b.Board[i][j] = EMPTY
}

func (b Board) Display() {
	fmt.Print("    ")
	for i := 0; i < b.Size; i++ {
		val := ((i % 10) + 1)
		if val > 9 {
			val = 0
		}
		fmt.Print(val, " ")
	}
	fmt.Println()
	for i := 0; i < b.Size; i++ {
		padding := ""
		// not < 10 because index + 1 for display
		if i < 9 {
			padding = " "
		}

		fmt.Printf(" %s%d ", padding, i+1)
		for j := 0; j < b.Size; j++ {
			fmt.Print(b.Board[i][j])
		}
		fmt.Print("\n")
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
	fmt.Println("==================")
	fmt.Println("1. For each turn, type in coordinates (e.g. 2, 3) to place your piece")
	fmt.Println("2. Type -1 to pass")
	fmt.Println()

	current_player := 1

	consecutive_passes := []int{0, 0}

	// Game Loop
	for {
		fmt.Printf("Play %d's turn to place a %v piece\n", current_player, Piece(current_player).Name())

		var x int
		var y int

		board.Display()

		fmt.Print("Type in coordinates (e.g. 2, 3): ")
		fmt.Scan(&x)

		if x == -1 {
			consecutive_passes[current_player-1] += 1
			if consecutive_passes[current_player-1] == 2 {
				var winner int

				if current_player == 1 {
					winner = 2
				} else {
					winner = 1
				}

				// Game lost for current_player
				game_won(winner)
			}
			continue
		}

		fmt.Scan(&y)

		// Note the swap of x and y for i and j and the 1 indexed input coords
		if board.CanPlace(y-1, x-1) {
			board.Set(y-1, x-1, Piece(current_player))
		}

		if current_player == 1 {
			current_player = 2
		} else {
			current_player = 1
		}
	}
}
