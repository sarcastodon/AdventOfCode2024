package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	phraseRE := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	numRE := regexp.MustCompile(`[\d]{1,3}`)
	//	re := regexp.MustCompile(`mul\([\d]{1,3}`)
	f, err := os.Open("inputa.txt")
	if err != nil {
		fmt.Println("file not found")
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		chunks := phraseRE.FindAllString(scanner.Text(), -1)
		for _, chunk := range chunks {
			nums := numRE.FindAllString(chunk, -1)
			if len(nums) != 2 {
				fmt.Println("Parse failure: ", chunk)
			}
			first, err := strconv.Atoi(nums[0])
			if err != nil {
				fmt.Println("Can't convert ", nums[0])
			}
			second, err := strconv.Atoi(nums[1])
			if err != nil {
				fmt.Println("Can't convert ", nums[1])
			}
			total += first * second
		}
	}
	fmt.Println("The answer: ", total)
}
