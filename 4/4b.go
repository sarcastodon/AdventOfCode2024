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

func testPairwise(first byte, second byte) bool {
	if (first == 'M' && second == 'S') ||
		(first == 'S' && second == 'M') {
		return true
	}
	return false
}

func (p *puzzle) lookForTMatch(row int, column int) {
	if testPairwise(p.field[row+1][column].char, p.field[row-1][column].char) &&
		testPairwise(p.field[row][column+1].char, p.field[row][column-1].char) {
		fmt.Println("T match found at", row, column)
		//		p.numberMatched++
	}
}

func (p *puzzle) lookForXMatch(row int, column int) {
	if testPairwise(p.field[row+1][column-1].char, p.field[row-1][column+1].char) &&
		testPairwise(p.field[row+1][column+1].char, p.field[row-1][column-1].char) {
		fmt.Println("X match found at", row, column)
		p.numberMatched++
	}
}

func (p *puzzle) solve() {
	for i := 1; i < len(p.field)-1; i++ {
		for j := 1; j < len(p.field[i])-1; j++ {
			if p.field[i][j].char == 'A' {
				p.lookForTMatch(i, j)
				p.lookForXMatch(i, j)
			}
		}
	}
}

func main() {
	var p puzzle

	f, err := os.Open("inputa.txt")
	if err != nil {
		fmt.Println("Unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineLength := 0
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

	}
	p.solve()
	fmt.Println("The answer is", p.numberMatched)
}
