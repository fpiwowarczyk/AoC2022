package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type File int

type Directory struct {
	Size int
	Dirs []string
	Deep int
}

func NewDirectory() *Directory {
	return &Directory{Size: 0}
}

func main() {
	data := readInput("./input")
	commandsAndOutputs := strings.Split(data, "\n")
	commandsAndOutputs = commandsAndOutputs[:len(commandsAndOutputs)-1]
	directories := make(map[string]*Directory)
	var currentDirectory []string
	deepnes := 0
	for _, command := range commandsAndOutputs {
		if strings.HasPrefix(command, "$ cd") {
			parts := strings.Split(command, " ")
			if parts[2] != ".." {
				currentDirectory = append(currentDirectory, parts[2])
				deepnes++
			} else if parts[2] == ".." {
				currentDirectory = currentDirectory[:len(currentDirectory)-1]
				deepnes--
			} else {
				fmt.Println("SHOUDNT BE HERE --------------")
			}
			curr := arrayToString(currentDirectory)
			if _, ok := directories[curr]; !ok {
				directories[curr] = NewDirectory()
				directories[curr].Deep = deepnes
			}
		}
		if strings.HasPrefix(command, "dir") {
			parts := strings.Split(command, " ")
			dirInside := arrayToString(currentDirectory) + " " + parts[1]
			directories[arrayToString(currentDirectory)].Dirs = append(directories[arrayToString(currentDirectory)].Dirs, dirInside)
		}
		if !strings.HasPrefix(command, "dir") && !strings.HasPrefix(command, "$") {
			parts := strings.Split(command, " ")
			value, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			directories[arrayToString(currentDirectory)].Size += value
		}
	}
	max := findMaxDeep(directories)
	fmt.Println("MAX:", max)
	for i := max; i >= 0; i-- {
		for k, v := range directories {
			if v.Deep == i {
				v.Size = calculateSize(k, directories, i)
			}
		}
	}

	diskSize := 70_000_000
	neededSpace := 30_000_000
	freeSpace := diskSize - directories["/"].Size
	fmt.Println(freeSpace)
	neededToFree := neededSpace - freeSpace
	fmt.Println(neededToFree)
	smallest := directories["/"].Size

	for _, v := range directories {
		if neededToFree-v.Size < 0 && v.Size < smallest {
			smallest = v.Size
		}
	}
	fmt.Println(smallest)
}

func findMaxDeep(thisMap map[string]*Directory) int {
	output := 0
	for _, v := range thisMap {
		if v.Deep > output {
			output = v.Deep
		}
	}
	return output
}

func calculateSize(name string, directories map[string]*Directory, deep int) int {
	directory := directories[name]
	output := 0
	if len(directory.Dirs) == 0 || directory.Deep > deep {
		return directory.Size
	}
	for _, insideDir := range directory.Dirs {
		output += calculateSize(insideDir, directories, deep)
	}
	return output + directory.Size
}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func arrayToString(array []string) string {
	output := ""
	for i, val := range array {
		if i == 0 {
			output += val
		} else {
			output += " " + val
		}
	}
	return output
}
