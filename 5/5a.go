package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "inputa.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Couldn't open", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// For each first page, all of the pages that can only come after.
	rules := map[int]map[int]bool{}
	answer := 0

	beginProcessingUpdates := false
	for scanner.Scan() {
		// Build the map of ...
		line := scanner.Text()
		if !beginProcessingUpdates {
			chunks := strings.Split(line, "|")
			// Blank line separates rules and updates
			if len(chunks) == 1 {
				beginProcessingUpdates = true
				continue
			} else if len(chunks) != 2 {
				fmt.Println("Panic! Line doesn't match rule pattern.")
			}
			first, _ := strconv.Atoi(chunks[0])
			second, _ := strconv.Atoi(chunks[1])
			follower, ok := rules[first]
			if !ok {
				follower = make(map[int]bool)
				rules[first] = follower
			}
			rules[first][second] = true
		} else {
			validUpdate := true
			chunks := strings.Split(line, ",")
			updates := map[int]bool{}
			for _, chunk := range chunks {
				page, _ := strconv.Atoi(chunk)

				// Does this update page have rules?
				followers, ok := rules[page]
				if ok {
					// Make sure none of the follower pages have already been seen
					// as that would break the rule.
					for follower := range followers {
						if updates[follower] == true {
							validUpdate = false
							break
						}
					}
				}
				if !validUpdate {
					break
				}
				// Add this page to the seen updates.
				updates[page] = true
			}
			if validUpdate {
				mid, _ := strconv.Atoi(chunks[(len(chunks)-1)/2])
				fmt.Println("This update valid: ", line, " Adding", mid, "to total")
				answer += mid
			}
		}
	}

	fmt.Println("The answer is", answer)
}
