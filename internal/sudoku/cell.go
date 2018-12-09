package sudoku

import "fmt"

type Cell struct {
	Val int
}

func (c Cell) String() string {
	return fmt.Sprintf("%v", c.Val)
}
