package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput("./input")
	topAndBottom := strings.Split(data, "\n\n")
	top := topAndBottom[0]
	bottom := topAndBottom[1]
	rows := strings.Split(top, "\n")
	onlyLetters := strings.Split(rows[len(rows)-2], " ")
	columnsLen := len(onlyLetters)
	var columns = make([][]string, columnsLen)
	rows = rows[:len(rows)-1]
	for _, row := range rows {
		for j := 1; j <= columnsLen; j++ {
			if string(row[j*4-3]) != " " {
				columns[j-1] = append([]string{string(row[j*4-3])}, columns[j-1]...)
			}
		}
	}
	tasks := strings.Split(bottom, "\n")
	tasks = tasks[:len(tasks)-1]

	for _, task := range tasks {
		move, from, to := getMoves(task)
		var els = []string{}
		for i := 1; i <= move; i++ {
			var el string
			el, columns[from-1] = removeEl(columns[from-1])
			els = append(els, el)
		}
		for i := 0; i <= len(els)-1; i++ {
			columns[to-1] = addEl(columns[to-1], els[len(els)-i-1])
		}
		fmt.Println(columns)

	}
	for _, column := range columns {
		fmt.Print(column[len(column)-1])
	}
}

func removeEl(array []string) (string, []string) {
	return string(array[len(array)-1]), array[:len(array)-1]
}
func addEl(array []string, el string) []string {
	return append(array, el)
}

func getMoves(sentence string) (move, from, to int) {
	moves := strings.Split(sentence, " ")
	move, _ = strconv.Atoi(moves[1])
	from, _ = strconv.Atoi(moves[3])
	to, _ = strconv.Atoi(moves[5])
	return move, from, to
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}
