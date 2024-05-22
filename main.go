package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	skills "skills/skills"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	file := os.Args[1]
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(content)
	text = strings.TrimSuffix(text, "\n")
	strArr := strings.Split(text, "\n")
	var numbers []float64
	for _, v := range strArr {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		numbers = append(numbers, n)
	}

	average := (skills.Mean(numbers))
	median := (skills.Median(numbers))
	variance := (skills.Variance(numbers, average))
	standardDeviation := (skills.StandardDeviation(variance))
	fmt.Printf("Average: %.f\nMedian: %.f\nVariance: %.f\nStandard Deviation: %.f\n", average, median, variance, standardDeviation)
}
