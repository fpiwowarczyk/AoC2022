package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3

	win  = 6
	draw = 3
	lose = 0
)

func main() {
	data := readInput("./input")
	inputs := strings.Split(data, "\n")
	points := 0
	for _, game := range inputs {
		if len(game) == 0 {
			continue
		}
		moves := strings.Split(game, " ")
		enemy := moves[0]
		my := moves[1]
		//points += calculatePointsStareOne(enemy, my)
		points += calculatePointsStartTwo(enemy, my)
	}
	fmt.Println(points)
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// X - lose
// Y - draw
// Z - win

func calculatePointsStartTwo(enemy, endResult string) int {
	output := 0
	switch enemy {
	case "A": // rock
		switch endResult {
		case "X": // lose
			output = lose + scissors
		case "Y": // draw
			output = draw + rock
		case "Z": // win
			output = win + paper
		}
	case "B": // paper
		switch endResult {
		case "X": // lose
			output = lose + rock
		case "Y": // draw
			output = draw + paper
		case "Z": // win
			output = win + scissors
		}
	case "C": // scissors
		switch endResult {
		case "X": // lose
			output = lose + paper
		case "Y": // draw
			output = draw + scissors
		case "Z": // win
			output = win + rock
		}
	}
	return output
}

// X - rock
// Y - paper
// Z - scissors
func calculatePointsStareOne(enemy, my string) int {
	output := 0
	switch my {
	case "X": //rock
		output += rock
		switch enemy {
		case "A": // rock
			output += draw
		case "B": // paper
			output += lose
		case "C": // scissors
			output += win
		}
	case "Y": //paper
		output += paper
		switch enemy {
		case "A": // rock
			output += win
		case "B": // paper
			output += draw
		case "C": //scissors
			output += lose
		}
	case "Z": // scissors
		output += scissors
		switch enemy {
		case "A": // rock
			output += lose
		case "B": //paper
			output += win
		case "C": //scissors
			output += draw
		}
	}
	fmt.Println("add ", output)
	return output
}
