package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData := readInput("./input")

	caloriesSplitted := strings.Split(inputData, "\n")

	var elvesSummedCallories []int
	tempCalories := 0
	for _, calories := range caloriesSplitted {
		if calories == "" {
			elvesSummedCallories = append(elvesSummedCallories, tempCalories)
			tempCalories = 0
			continue
		}
		cal, err := strconv.Atoi(calories)
		if err != nil {
			panic(err)
		}
		tempCalories += cal
	}
	top3Elves := max3(elvesSummedCallories)
	fmt.Println(max3(elvesSummedCallories))
	fmt.Println(sumEls(top3Elves))
}

func sumEls(slice []int) int {
	output := 0
	for _, el := range slice {
		output += el
	}
	return output
}

func readInput(dir string) string {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func max3(slice []int) []int {
	topN := 0
	var output []int
	for i := 0; i < 3; i++ {
		for _, s := range slice {
			if s > topN {
				if !contains(output, s) {
					topN = s
				}
			}
		}
		output = append(output, topN)
		topN = 0
	}
	return output
}

func contains(slice []int, val int) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}
