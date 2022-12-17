package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput("./input")
	instructions := strings.Split(data, "\n")
	cycleNumber := 0
	registerX := 1
	instructionIndex := -1
	bussy := 0
	currentInstruction := ""
	toAdd := 0
	drawing := ""
	for instructionIndex < len(instructions)-1 {

		if cycleNumber > 0 && cycleNumber%40 == 0 {
			drawing += "\n"
		}
		if (cycleNumber%40) == registerX-1 || cycleNumber%40 == registerX || cycleNumber%40 == registerX+1 {
			drawing += "#"
		} else {
			drawing += "."
		}
		//DrawCRT(registerX)
		//fmt.Println(drawing, "Cycle:", cycleNumber)
		if bussy > 0 {
			cycleNumber++
			bussy--
			if bussy == 0 {
				registerX += toAdd
			}
			continue
		}
		if bussy == 0 {
			instructionIndex++
		}
		currentInstruction = instructions[instructionIndex]
		if strings.HasPrefix(currentInstruction, "noop") {
			bussy = 1
			toAdd = 0
			bussy--
			cycleNumber++
		}
		if strings.HasPrefix(currentInstruction, "addx") {
			parts := strings.Split(currentInstruction, " ")
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			toAdd = val
			bussy = 2
			bussy--
			cycleNumber++
		}

	}
	fmt.Print(drawing)
}

func DrawCRT(registerX int) {
	dr := ""
	for i := 0; i <= 40; i++ {
		if i == registerX-1 || i == registerX || i == registerX+1 {
			dr += "#"
		} else {
			dr += "."
		}
	}
	fmt.Println("Sprite "+dr+" registerX:", registerX)
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}
