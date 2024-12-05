package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reportIsSafe(report []string) bool {
	direction := 0
	for i := 0; i < (len(report) - 1); i++ {
		first, err := strconv.Atoi(report[i])
		if err != nil {
			fmt.Println("Can't convert \"", report[i], "\"to int.")
		}
		second, err := strconv.Atoi(report[i+1])
		if err != nil {
			fmt.Println("Can't convert \"", report[i], "\"to int.")
		}
		diff := first - second
		if (diff == 0) || (diff > 3) || (diff < -3) {
			return false
		}
		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else if (direction == -1) && (diff > 0) {
			return false
		} else if (direction == 1) && (diff < 0) {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("inputa.txt")
	if err != nil {
		fmt.Println("file not found")
	}
	defer f.Close()

	safeReportCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		if reportIsSafe(report) {
			safeReportCount++
		}
	}
	fmt.Println("The number of safe reports is ", safeReportCount)
}
