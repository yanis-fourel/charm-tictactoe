package main

import (
	"fmt"
	"log"
	"strconv"
)

type Tile rune

const (
	Tile_Empty Tile = '.'
	Tile_O     Tile = 'O'
	Tile_X     Tile = 'X'
)

// Constant
var TeamName = map[Tile]string{
	Tile_Empty: "No one",
	Tile_O:     "Circle",
	Tile_X:     "Cross",
}

type Board struct {
	data [9]Tile
}

func NewBoard() Board {
	data := [9]Tile{}
	for i := range data {
		data[i] = Tile_Empty
	}
	return Board{
		data,
	}
}

func (b *Board) get(x, y int) Tile {
	return b.data[x+3*y]
}

func (b *Board) set(x, y int, val Tile) {
	b.data[x+3*y] = val
}

func (b *Board) display() {
	for i := 0; i < 9; i++ {
		fmt.Printf("%c", b.data[i])
		if i%3 == 2 {
			fmt.Println()
		}
	}
}

// Returns Tile_Empty if no winner
func (b *Board) getWinner() Tile {
	// Columns
	for x := 0; x < 3; x++ {
		if b.get(x, 0) != Tile_Empty && b.get(x, 0) == b.get(x, 1) && b.get(x, 1) == b.get(x, 2) {
			return b.get(x, 0)
		}
	}

	// Rows
	for y := 0; y < 3; y++ {
		if b.get(0, y) != Tile_Empty && b.get(0, y) == b.get(1, y) && b.get(1, y) == b.get(2, y) {
			return b.get(0, y)
		}
	}

	// Diagonals
	if b.get(1, 1) != Tile_Empty && b.get(0, 0) == b.get(1, 1) && b.get(1, 1) == b.get(2, 2) {
		return b.get(1, 1)
	}
	if b.get(1, 1) != Tile_Empty && b.get(2, 0) == b.get(1, 1) && b.get(1, 1) == b.get(0, 2) {
		return b.get(1, 1)
	}

	return Tile_Empty
}

func getInput() (x, y int) {
	var input string
	ok := false

	for ok == false {
		fmt.Print("Row (1 top, 2 middle, 3 bottom): ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalln(err)
		}
		y, err = strconv.Atoi(input)
		if err != nil {
			log.Fatalln(err)
		}
		if y >= 1 && y <= 3 {
			ok = true
			y-- // User input is 1-based
		} else {
			fmt.Printf("%d is out of range\n", y)
		}
	}

	ok = false
	for ok == false {
		fmt.Print("Column (1 left, 2 middle, 3 right): ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalln(err)
		}
		x, err = strconv.Atoi(input)
		if err != nil {
			log.Fatalln(err)
		}
		if x >= 1 && x <= 3 {
			ok = true
			x-- // User input is 1-based
		} else {
			fmt.Printf("%d is out of range\n", x)
		}
	}

	return
}

func main() {
	b := NewBoard()

	turn := Tile_X
	winner := Tile_Empty

	for turn_count := 0; turn_count < 9; turn_count++ {
		b.display()
		fmt.Println("----------")
		fmt.Printf("%s to play\n", TeamName[turn])

		x, y := getInput()
		for b.get(x, y) != Tile_Empty {
			fmt.Println("Tile is not empty")
			x, y = getInput()
		}

		b.set(x, y, turn)

		switch turn {
		case Tile_O:
			turn = Tile_X
		case Tile_X:
			turn = Tile_O
		}

		winner = b.getWinner()
		if winner != Tile_Empty {
			break
		}
	}

	if winner != Tile_Empty {
		fmt.Println("---")
		b.display()
		fmt.Println("---")
		fmt.Printf("%s won!\n", TeamName[winner])
	} else {
		fmt.Println("It's a tie!")
	}
}
