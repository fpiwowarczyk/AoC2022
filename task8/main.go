package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput("./input")
	rows := strings.Split(data, "\n")
	var grid [][]int
	for _, row := range rows {
		var line []int
		for _, tree := range row {
			val, err := strconv.Atoi(string(tree))
			if err != nil {
				panic(err)
			}
			line = append(line, val)
		}
		if len(line) == 0 {
			continue
		}
		grid = append(grid, line)
	}
	visible := 0
	bestView, bestViewI, bestViewJ := 0, 0, 0
	tempJ := 0
	for i, row := range grid {
		leftView, rightView, upView, downView := 0, 0, 0, 0
		for j, tree := range row {
			tempJ = j
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1 {
				leftView, rightView, upView, downView = 0, 0, 0, 0
				visible++
			} else {
				visibleLeft, visibleRight, visibleUp, visibleBottom := false, false, false, false
				visibleLeft, leftView = left(tree, i, j, grid)
				visibleRight, rightView = right(tree, i, j, grid)
				visibleUp, upView = up(tree, i, j, grid)
				if i == 2 && j == 1 {
					fmt.Printf("%d %d upView %d\n", i, j, upView)
				}
				visibleBottom, downView = down(tree, i, j, grid)
				if visibleLeft || visibleRight || visibleUp || visibleBottom {
					visible++
				}
			}
			if leftView*rightView*upView*downView != 0 {
				fmt.Printf("%d*%d*%d*%d = %d [%d,%d]\n", leftView, rightView, upView, downView, leftView*rightView*upView*downView, i, j)
			}
			if bestView < leftView*rightView*upView*downView {
				bestView, bestViewI, bestViewJ = leftView*rightView*upView*downView, i, tempJ
			}
		}

	}
	fmt.Printf("Visible %d, best view scopre %d i: %d j: %d \n", visible, bestView, bestViewI, bestViewJ)
}

func left(tree, i, j int, grid [][]int) (bool, int) {
	check := j - 1
	viewScore := 0
	for check >= 0 {
		viewScore++
		if tree <= grid[i][check] {
			fmt.Printf("LEFT return view %d\n", viewScore)

			return false, viewScore
		}
		check--
	}
	fmt.Printf("LEFT return view %d\n", viewScore)

	return true, viewScore
}

func right(tree, i, j int, grid [][]int) (bool, int) {
	check := j + 1
	viewScore := 0
	for check < len(grid[0]) {
		viewScore++
		if tree <= grid[i][check] {
			fmt.Printf("RIGHt return view %d\n", viewScore)

			return false, viewScore
		}
		check++
	}
	fmt.Printf("RIFGHT return view %d\n", viewScore)

	return true, viewScore
}

func up(tree, i, j int, grid [][]int) (bool, int) {
	check := i - 1
	viewScore := 0
	for check >= 0 {
		viewScore++
		if tree <= grid[check][j] {
			fmt.Printf("UP return view %d\n", viewScore)
			return false, viewScore
		}
		check--
	}
	fmt.Printf("UP return view %d\n", viewScore)

	return true, viewScore
}

func down(tree, i, j int, grid [][]int) (bool, int) {
	check := i + 1
	viewScore := 0
	for check < len(grid) {
		viewScore++
		if tree <= grid[check][j] {
			fmt.Printf("DOWN return view %d\n", viewScore)

			return false, viewScore
		}
		check++
	}
	fmt.Printf("DOWN return view %d\n", viewScore)

	return true, viewScore
}
func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// 1005 too low
