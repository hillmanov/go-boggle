package board

type board struct {
	grid [][]string
	size int
}

func New(letters []string) *board {
	if !(len(letters) == 16 || len(letters) == 25) {
		panic("Board must be 4x4 or 5x5!")
	}
	b := new(board)

	if len(letters) == 16 {
		b.size = 4
	} else {
		b.size = 5
	}

	for row := 0; row < b.size; row++ {
		b.grid = append(b.grid, letters[(row*b.size):((row*b.size)+b.size)])
	}

	return b
}

func (b *board) Letter(row, col int) string {
	return b.grid[row][col]
}

func (b *board) Size() int {
	return b.size
}
