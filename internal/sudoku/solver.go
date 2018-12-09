package sudoku

import "fmt"

func SolveBoard(b *Board) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			valid := false
			i := b.grid[x][y].Val

			for i < 9 {
				i++
				b.grid[x][y] = Cell{i}

				// check validity
				if b.Validate() {
					valid = true
					break
				}
			}

			if !valid {
				b.grid[x][y] = Cell{0}
				fmt.Printf("begin backtracking! pos: %v, %v \n", x, y)
				y = y - 2
				// TODO: Fix valid range for going back
			}
		}

	}

	b.Print()
}
