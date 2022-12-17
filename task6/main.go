package main

import (
	"fmt"
	"os"
)

func main() {
	data := readInput("./input")
	seq := ""
	for i, characher := range data {
		seq += string(characher)
		if len(seq) > 14 {
			seq = seq[1:]
		}
		if len(seq) == 14 && allDifferentChars(seq) {
			fmt.Println(seq, i+1)
			break
		}
	}
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func allDifferentChars(input string) bool {
	for i, _ := range input {
		for j, _ := range input {
			if i == j {
				continue
			}
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}
