package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddOrMultiply(current int, remainingValues []int, target int) bool {
	if len(remainingValues) == 0 {
		if current == target {
			return true
		} else {
			return false
		}
	} else if current > target {
		// Early return: no negative values, so if we're already over, kill it.
		return false
	} else {
		// Add
		if AddOrMultiply(current+remainingValues[0], remainingValues[1:], target) {
			return true
		}

		// Multiply
		if AddOrMultiply(current*remainingValues[0], remainingValues[1:], target) {
			return true
		}

		return false
	}
}

func AddOrMultiplyOrConcatenate(current int, remainingValues []int, target int) bool {
	if len(remainingValues) == 0 {
		if current == target {
			fmt.Println("Found the target", target)
			return true
		}
		return false
	} else if current > target {
		// Early return: no negative values, so if we're already over, kill it.
		return false
	} else {
		// Add
		if AddOrMultiplyOrConcatenate(current+remainingValues[0], remainingValues[1:], target) {
			return true
		}

		// Multiply
		if AddOrMultiplyOrConcatenate(current*remainingValues[0], remainingValues[1:], target) {
			return true
		}

		// Concatenate
		firsthalf := strconv.Itoa(current)
		secondhalf := strconv.Itoa(remainingValues[0])
		value, err := strconv.Atoi(firsthalf + secondhalf)
		if err != nil {
			fmt.Println("strcat", firsthalf, secondhalf)
		}

		if AddOrMultiplyOrConcatenate(value, remainingValues[1:], target) {
			return true
		}

		return false
	}
}

func main() {
	filename := "inputa.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open file", filename)
	}
	defer f.Close()

	validCalibrations := 0
	sumOfCalibrations := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		total, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Bad Atoi", parts[0])
		}
		chunks := strings.Split(parts[1], " ")
		values := make([]int, len(chunks))
		for i, chunk := range chunks {
			value, err := strconv.Atoi(chunk)
			if err != nil {
				fmt.Println("Bad Atoi", chunk, "<--")
			}
			values[i] = value
		}
		if AddOrMultiplyOrConcatenate(0, values[:], total) {
			validCalibrations += 1
			sumOfCalibrations += total
		}
	}
	fmt.Println("ValidCalibrations", validCalibrations)
	fmt.Println("SumOfCalibrations", sumOfCalibrations)
}
