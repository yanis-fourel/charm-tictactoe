package main

type Board struct {
	data [9]Team
}

func NewBoard() Board {
	data := [9]Team{}
	for i := range data {
		data[i] = Team_None
	}
	return Board{
		data,
	}
}

func (b *Board) get(x, y int) Team {
	return b.data[x+3*y]
}

func (b *Board) set(x, y int, val Team) {
	b.data[x+3*y] = val
}

// Returns Team_Empty if no winner
func (b *Board) getWinner() Team {
	// Columns
	for x := 0; x < 3; x++ {
		if b.get(x, 0) != Team_None && b.get(x, 0) == b.get(x, 1) && b.get(x, 1) == b.get(x, 2) {
			return b.get(x, 0)
		}
	}

	// Rows
	for y := 0; y < 3; y++ {
		if b.get(0, y) != Team_None && b.get(0, y) == b.get(1, y) && b.get(1, y) == b.get(2, y) {
			return b.get(0, y)
		}
	}

	// Diagonals
	if b.get(1, 1) != Team_None && b.get(0, 0) == b.get(1, 1) && b.get(1, 1) == b.get(2, 2) {
		return b.get(1, 1)
	}
	if b.get(1, 1) != Team_None && b.get(2, 0) == b.get(1, 1) && b.get(1, 1) == b.get(0, 2) {
		return b.get(1, 1)
	}

	return Team_None
}
