package sudoku

import "fmt"

type Board struct {
	grid [][]Cell
}

func NewBoard(size int) Board {
	grid := make([][]Cell, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]Cell, size)
	}

	return Board{grid}
}

func (b *Board) Print() {
	for x := 0; x < 9; x++ {

		for y := 0; y < 9; y++ {
			fmt.Printf("%v", b.grid[x][y])
		}

		fmt.Println()
	}
}

func (b *Board) Validate() bool {
	// non-optimized way just blindly checks all rows, cols, boxes until false.

	// Validate Row
	for x := 0; x < 9; x++ {
		row := [9]int{}

		for y := 0; y < 9; y++ {
			val := b.grid[x][y].Val - 1

			if val == -1 {
				break
			}

			row[val] += row[val] + 1

			if row[val] > 1 {
				return false
			}
		}
	}

	// Validate Column
	for y := 0; y < 9; y++ {
		column := [9]int{}

		for x := 0; x < 9; x++ {
			val := b.grid[x][y].Val - 1

			if val == -1 {
				break
			}

			column[val] += column[val] + 1

			if column[val] > 1 {
				return false
			}
		}

		// TODO: Validate boxes.
	}

	return true
}
