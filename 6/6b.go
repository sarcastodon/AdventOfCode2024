package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	visited     bool
	obstacle    bool
	visitedDirs []byte
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

func (p *puzzle) deepCopy() puzzle {
	var pCopy puzzle
	pCopy.row = p.row
	pCopy.column = p.column
	pCopy.direction = p.direction
	pCopy.numRows = p.numRows
	pCopy.numColumns = p.numColumns
	pCopy.field = make([][]node, p.numRows)
	for i := range pCopy.field {
		pCopy.field[i] = make([]node, p.numColumns)
		copy(pCopy.field[i], p.field[i])
	}
	return pCopy
}

// Returns false if it char has moved off field.
func (p *puzzle) move() string {
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
		return "exited"
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
		if p.updatePosition(nextRow, nextColumn, p.direction) {
			return "looped"
		}
	}
	return "normal"
}

func (p *puzzle) printField() {
	for i := 0; i < p.numRows; i++ {
		line := make([]byte, p.numColumns)
		for j := 0; j < p.numColumns; j++ {
			if p.row == i && p.column == j {
				line[j] = p.direction
			} else if p.field[i][j].obstacle {
				line[j] = '#'
			} else if p.field[i][j].visited || len(p.field[i][j].visitedDirs) > 0 {
				line[j] = 'X'
			} else {
				line[j] = '.'
			}
		}
		fmt.Println(string(line[:]))
	}
	fmt.Println()
}

// Returns true if new Position has been visited in same direction before (which indicates loop)
func (p *puzzle) updatePosition(row int, column int, direction byte) bool {
	p.row = row
	p.column = column
	p.direction = direction

	for _, dirs := range p.field[row][column].visitedDirs {
		if dirs == direction {
			return true
		}
	}
	p.field[row][column].visitedDirs = append(p.field[row][column].visitedDirs, direction)
	return false
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
			row[i].visitedDirs = make([]byte, 0)
		}
		p.field = append(p.field, row)
	}
	p.updatePosition(firstRow, firstColumn, firstDirection)
	p.numRows = len(p.field)
}

func main() {
	var p puzzle

	p.initializePuzzle("inputa.txt")

	numLooped := 0
	for i := 0; i < p.numRows; i++ {
		for j := 0; j < p.numColumns; j++ {
			if i == p.row && j == p.column {
				// Can't place an article on top of starting position.
				fmt.Println("NEXT!!!")
				continue
			}
			pCopy := p.deepCopy()

			pCopy.field[i][j].obstacle = true
			var status string
			for status = "normal"; status == "normal"; status = pCopy.move() {
			}
			if status == "looped" {
				numLooped++
			}
			fmt.Println("Placing at", i, ",", j, "resulted in", status)
		}
	}

	fmt.Println("NumLooped:", numLooped)

}
