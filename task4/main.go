package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput("./input")
	pairs := strings.Split(data, "\n")
	pairs = removeEmpty(pairs)
	overlapsCout := 0
	for _, pair := range pairs {
		sections := strings.Split(pair, ",")
		if checkOverlap(sections[0], sections[1]) {
			overlapsCout++
		}
	}

	fmt.Println(overlapsCout)
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func removeEmpty(data []string) []string {
	for i, row := range data {
		if len(row) == 0 {
			data = append(data[:i], data[i+1:]...)
		}
	}
	return data
}

func checkOverlap(section1, section2 string) bool {
	vals1 := strings.Split(section1, "-")
	vals2 := strings.Split(section2, "-")
	startF, err := strconv.Atoi(vals1[0])
	if err != nil {
		panic(err)
	}
	endF, err := strconv.Atoi(vals1[1])
	if err != nil {
		panic(err)
	}
	startS, err := strconv.Atoi(vals2[0])
	if err != nil {
		panic(err)
	}
	endS, err := strconv.Atoi(vals2[1])
	if err != nil {
		panic(err)
	}
	if (endF >= startS && startF <= startS) ||
		(endS >= startF && startS <= startF) ||
		(startF <= startS && endF >= endS) ||
		(startS <= startF && endS >= endF) {
		fmt.Println((endF >= startS && startF <= startS))
		fmt.Println((endS >= startF && startS <= startF))
		fmt.Println((startF <= startS && endF >= endS))
		fmt.Println((startS <= startF && endS >= endF))
		fmt.Println(startS, endS, startF, endF)
		return true
	}
	//	if (startF <= startS) && (endF >= endS) {
	//		fmt.Println(startF, startS, startF <= startS, endF, endS, endF >= endS)
	//		return true
	//	}
	return false
}
