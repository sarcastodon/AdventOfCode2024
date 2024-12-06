package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseAndMultiply(phrase string) int {
	numRE := regexp.MustCompile(`[\d]{1,3}`)
	nums := numRE.FindAllString(phrase, -1)
	if len(nums) != 2 {
		fmt.Println("Parse failure: ", phrase)
	}
	first, err := strconv.Atoi(nums[0])
	if err != nil {
		fmt.Println("Can't convert ", nums[0])
	}
	second, err := strconv.Atoi(nums[1])
	if err != nil {
		fmt.Println("Can't convert ", nums[1])
	}
	return first * second
}

func main() {
	doRE := regexp.MustCompile(`do\(\)`)
	dontRE := regexp.MustCompile(`don\'t\(\)`)
	phraseRE := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)

	f, err := os.Open("inputa.txt")
	if err != nil {
		fmt.Println("file not found")
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	enabled := true

	for scanner.Scan() {
		line := scanner.Text()
		currentPosition := 0

		for currentPosition < len(line) {
			fmt.Println("Next loop. CurrentPosition:", currentPosition)
			// If not enabled, try to find the next "do()" on the current line.
			if enabled != true {
				chunkIndex := doRE.FindStringIndex(line[currentPosition:])
				// Reached the end of the line without enabling
				if chunkIndex == nil {
					currentPosition = len(line)
					fmt.Println("-->Disabled through EOL")
				} else {
					enabled = true
					currentPosition += chunkIndex[1]
					fmt.Println("-->Disabled until", currentPosition)
				}
			}

			// If we are now enabled, try to find the next "don't()" on the current line.
			if enabled {
				endPosition := len(line)
				chunkIndex := dontRE.FindStringIndex(line[currentPosition:])
				// Found a "don't", extract the substring and mark disabled for next.
				if chunkIndex != nil {
					enabled = false
					endPosition = currentPosition + chunkIndex[1]
				}
				fmt.Println("-->Enabled from", currentPosition, "to", endPosition)
				// Find the chunks in the substring
				chunks := phraseRE.FindAllString(line[currentPosition:endPosition], -1)
				fmt.Println("Chunks: ", chunks)
				currentPosition = endPosition

				for _, chunk := range chunks {
					total += parseAndMultiply(chunk)
				}
			}
		}
	}

	fmt.Println("The answer: ", total)
}
