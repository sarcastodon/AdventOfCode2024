package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	visited  bool
	obstacle bool
}

type puzzle struct {
	field       [][]node
	row         int
	column      int
	direction   byte
	numRows     int
	numColumns  int
	vistedCount int
}

// Returns false if it char has moved off field.
func (p *puzzle) move() bool {
	nextRow := p.row
	nextColumn := p.column
	switch p.direction {
	case '>':
		nextColumn += 1
	case '<':
		nextColumn -= 1
	case '^':
		nextRow -= 1
	case 'v':
		nextRow += 1
	}
	// Next position is off the map: Consider this game finished
	if nextRow < 0 || nextRow >= p.numRows || nextColumn < 0 || nextColumn >= p.numColumns {
		return false
	}
	if p.field[nextRow][nextColumn].obstacle {
		// Turn to the right
		switch p.direction {
		case '>':
			p.direction = 'v'
		case '<':
			p.direction = '^'
		case '^':
			p.direction = '>'
		case 'v':
			p.direction = '<'
		}
	} else {
		p.updatePosition(nextRow, nextColumn, p.direction)
	}
	return true
}

func (p *puzzle) printField() {
	for i := 0; i < p.numRows; i++ {
		line := make([]byte, p.numColumns)
		for j := 0; j < p.numColumns; j++ {
			if p.row == i && p.column == j {
				line[j] = p.direction
			} else if p.field[i][j].obstacle {
				line[j] = '#'
			} else if p.field[i][j].visited {
				line[j] = 'X'
			} else {
				line[j] = '.'
			}
		}
		fmt.Println(string(line[:]))
	}
	fmt.Println()
}

func (p *puzzle) updatePosition(row int, column int, direction byte) {
	p.row = row
	p.column = column
	p.direction = direction
	if p.field[row][column].visited == false {
		p.vistedCount++
	}
	p.field[row][column].visited = true
}

func (p *puzzle) initializePuzzle(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open file", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var firstRow, firstColumn int
	var firstDirection byte
	for scanner.Scan() {
		line := scanner.Text()
		if p.numColumns == 0 {
			p.numColumns = len(line)
		} else if len(line) != p.numColumns {
			fmt.Println("Line lengths do not match")
		}

		row := make([]node, p.numColumns)
		for i := 0; i < p.numColumns; i++ {
			switch line[i] {
			case '#':
				row[i].obstacle = true
			case '>':
				fallthrough
			case '<':
				fallthrough
			case '^':
				fallthrough
			case 'v':
				firstDirection = line[i]
				firstRow = len(p.field)
				firstColumn = i
			default:
			}
		}
		p.field = append(p.field, row)
	}
	p.updatePosition(firstRow, firstColumn, firstDirection)
	p.numRows = len(p.field)
}

func main() {
	var p puzzle

	p.initializePuzzle("inputa.txt")

	for p.move() {
		fmt.Println("Visited:", p.vistedCount)
	}
}
