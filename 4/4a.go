package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	char byte
}

type puzzle struct {
	field         [][]node
	numberMatched int
}

func (p *puzzle) startFromX(row int, column int) {
	// Make the string in 8 directions
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			fmt.Println("I,J:", i, ",", j)
			// Edge cases:
			if ((i == 0) && (j == 0)) || // Nonsense construction
				((row + 3*i) < 0) || // Breaks left boundary
				((row + 3*i) >= len(p.field)) || // Breaks right boundary
				((column + 3*j) < 0) || // Breaks top boundary
				((column + 3*j) >= len(p.field[row])) { // Breaks bottom boundary
				println("-->BREAK")
				continue
			}
			var word [4]byte
			for index := 0; index < 4; index++ {
				word[index] = p.field[row+index*i][column+index*j].char
			}

			if string(word[:]) == "XMAS" {
				fmt.Println("Woot!")
				fmt.Println("Found at", column, ",", row, "in direction", i, ",", j)
				p.numberMatched++
			}
		}
	}
}

func (p *puzzle) solve() {
	// Look for 'X' Outer is row number, inner is column number.
	for i := 0; i < len(p.field); i++ {
		for j := 0; j < len(p.field[i]); j++ {
			if p.field[i][j].char == 'X' {
				fmt.Println("An 'X' at ", i, ",", j)
				p.startFromX(i, j)
			}
		}
	}
}

func main() {
	var p puzzle

	f, err := os.Open("inputa.txt")
	if err != nil {
		fmt.Println("file not found")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineLength := 0
	columnLength := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineLength == 0 {
			lineLength = len(line)
		} else if len(line) != lineLength {
			fmt.Println("Line lengths do not match")
		}

		row := make([]node, lineLength)
		for i := 0; i < lineLength; i++ {
			row[i].char = line[i]
		}
		p.field = append(p.field, row)
		columnLength++
	}

	p.solve()
	fmt.Println("The answer is", p.numberMatched)

	//	fmt.Print(puzzle)
}
