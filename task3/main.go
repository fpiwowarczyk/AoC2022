package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	priorities := getPraiorities()
	data := readInput("./input")
	ruckstacks := strings.Split(data, "\n")
	commonItem := make(map[string]string)
	var rocksackGroups []string
	numIngroup := 0
	tempString := ""
	for _, ruckstack := range ruckstacks {
		numIngroup++
		tempString += ruckstack + "\n"
		if numIngroup == 3 {
			numIngroup = 0
			rocksackGroups = append(rocksackGroups, tempString)
			tempString = ""
		}
	}
	for _, group := range rocksackGroups {
		groupSacks := strings.Split(group, "\n")

		for _, item := range groupSacks[0] {
			if strings.Contains(groupSacks[1], string(item)) && strings.Contains(groupSacks[2], string(item)) {
				commonItem[group] = string(item)
			}
		}
	}

	fmt.Println(len(ruckstacks))
	sum := 0
	for key, val := range commonItem {
		fmt.Println(key, val)
		sum += priorities[val]
	}
	fmt.Println(len(ruckstacks), len(commonItem))
	fmt.Println(sum)
}

func contains(array []string, letter string) bool {
	for _, el := range array {
		if el == letter {
			return true
		}
	}
	return false
}

func getPraiorities() map[string]int {
	prio := make(map[string]int)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetUppder := strings.ToUpper(alphabet)
	alphabetBoth := alphabet + alphabetUppder
	for ind, val := range alphabetBoth {
		prio[string(val)] = ind + 1
	}
	return prio
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// FIrst task answer
// 	comp1 := ruckstack[0 : len(ruckstack)/2]
//		comp2 := ruckstack[len(ruckstack)/2:]
//		for _, item1 := range comp1 {
//			if strings.Contains(comp2, string(item1)) {
//				commonItem[ruckstack] = string(item1)
//			}
//
//		}
