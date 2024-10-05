package main

import "fmt"

type board struct {
	board [][]int
}

func set ()

func main() {
	board_size := 10

	fmt.Println("Go Game in Go Lang")

	board := make([][]int, board_size)
	for i := 0; i < board_size; i++ {
		board[i] = make([]int, board_size)
		fmt.Println(board[i])
	}
}
