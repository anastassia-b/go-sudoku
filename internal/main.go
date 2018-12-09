package main

import (
	"fmt"

	"github.com/anastassia-b/go-sudoku/internal/sudoku"
)

func main() {
	fmt.Println("Welcome to Sudoku Go!")
	board := sudoku.NewBoard(9)
	sudoku.SolveBoard(&board)

	// fmt.Printf("%v\n", board)
}
