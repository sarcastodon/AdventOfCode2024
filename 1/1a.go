package main

import (
	"fmt"
	"os"
	"slices"
)

func absInt(first, second int) int {
	if first-second < 0 {
		return second - first
	}
	return first - second
}

func main() {
	f, err := os.Open("inputa.txt")

	if err != nil {
		fmt.Println("file not found")
	}

	//	fs := bufio.NewScanner(f)

	// scanf the file, adding the two numbers to the two lists.
	var listOne []int
	var listTwo []int

	count := 0

	for {
		num1 := 0
		num2 := 0
		_, err := fmt.Fscanf(f, "%d   %d\n", &num1, &num2)
		if err != nil {
			break
		}
		listOne = append(listOne, num1)
		listTwo = append(listTwo, num2)
	}
	fmt.Printf("Read %d lines\n", count)
	slices.Sort(listOne)
	slices.Sort(listTwo)

	totalDistance := 0
	for i := 0; i < len(listOne); i++ {
		totalDistance += absInt(listOne[i], listTwo[i])
	}

	fmt.Printf("The answer: %d!", totalDistance)
}
