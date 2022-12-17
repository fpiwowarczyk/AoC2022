package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Tail struct {
	X    int
	Y    int
	HEAD *Tail
}

const (
	startX = 30
	startY = 30
)

// 2531 wrong but close
func main() {
	data := readInput("./input")
	seriesOfMoves := strings.Split(data, "\n")
	seriesOfMoves = seriesOfMoves[:len(seriesOfMoves)-1]
	HEAD := Tail{X: startX, Y: startY, HEAD: nil}
	TAIL1 := Tail{X: startX, Y: startY, HEAD: &HEAD}
	TAIL2 := Tail{X: startX, Y: startY, HEAD: &TAIL1}
	TAIL3 := Tail{X: startX, Y: startY, HEAD: &TAIL2}
	TAIL4 := Tail{X: startX, Y: startY, HEAD: &TAIL3}
	TAIL5 := Tail{X: startX, Y: startY, HEAD: &TAIL4}
	TAIL6 := Tail{X: startX, Y: startY, HEAD: &TAIL5}
	TAIL7 := Tail{X: startX, Y: startY, HEAD: &TAIL6}
	TAIL8 := Tail{X: startX, Y: startY, HEAD: &TAIL7}
	TAIL9 := Tail{X: startX, Y: startY, HEAD: &TAIL8}
	uniquePosition9 := make(map[string]bool)
	uniquePosition8 := make(map[string]bool)
	uniquePosition7 := make(map[string]bool)
	uniquePosition6 := make(map[string]bool)
	uniquePosition5 := make(map[string]bool)
	uniquePosition4 := make(map[string]bool)
	uniquePosition3 := make(map[string]bool)
	uniquePosition2 := make(map[string]bool)
	uniquePosition1 := make(map[string]bool)
	for _, series := range seriesOfMoves {
		parts := strings.Split(series, " ")
		direction := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		times := value
		for i := 1; i <= times; i++ {
			HEAD.updateHEADPosition(direction)
			TAIL1.updatePosition()
			TAIL2.updatePosition()
			TAIL3.updatePosition()
			TAIL4.updatePosition()
			TAIL5.updatePosition()
			TAIL6.updatePosition()
			TAIL7.updatePosition()
			TAIL8.updatePosition()
			TAIL9.updatePosition()
			uniquePosition9[fmt.Sprintf("%d,%d", TAIL9.X, TAIL9.Y)] = true
			uniquePosition8[fmt.Sprintf("%d,%d", TAIL8.X, TAIL8.Y)] = true
			uniquePosition7[fmt.Sprintf("%d,%d", TAIL7.X, TAIL7.Y)] = true
			uniquePosition6[fmt.Sprintf("%d,%d", TAIL6.X, TAIL6.Y)] = true
			uniquePosition5[fmt.Sprintf("%d,%d", TAIL5.X, TAIL5.Y)] = true
			uniquePosition4[fmt.Sprintf("%d,%d", TAIL4.X, TAIL4.Y)] = true
			uniquePosition3[fmt.Sprintf("%d,%d", TAIL3.X, TAIL3.Y)] = true
			uniquePosition2[fmt.Sprintf("%d,%d", TAIL2.X, TAIL2.Y)] = true
			uniquePosition1[fmt.Sprintf("%d,%d", TAIL1.X, TAIL1.Y)] = true

			//PrintMap(uniquePosition9, uniquePosition8, uniquePosition7, uniquePosition6, uniquePosition5, uniquePosition4, uniquePosition3, uniquePosition2, uniquePosition1)
		}
	}
	output := 0
	for range uniquePosition9 {
		output++
	}
	fmt.Println(output)
}

type Pair struct {
	X int
	Y int
}

func PrintMap(uniquePosition9 map[string]bool, uniquePosition8 map[string]bool, uniquePosition7 map[string]bool, uniquePosition6 map[string]bool, uniquePosition5 map[string]bool, uniquePosition4 map[string]bool, uniquePosition3 map[string]bool, uniquePosition2 map[string]bool, uniquePosition1 map[string]bool) {
	cmd := exec.Command("clear")
	cmd.Run()
	var uniques9 []Pair
	var uniques8 []Pair
	var uniques7 []Pair
	var uniques6 []Pair
	var uniques5 []Pair
	var uniques4 []Pair
	var uniques3 []Pair
	var uniques2 []Pair
	var uniques1 []Pair
	for k := range uniquePosition9 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques9 = append(uniques9, Pair{X: val1, Y: val2})
	}

	for k := range uniquePosition8 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques8 = append(uniques8, Pair{X: val1, Y: val2})
	}

	for k := range uniquePosition7 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques7 = append(uniques7, Pair{X: val1, Y: val2})
	}
	for k := range uniquePosition6 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques6 = append(uniques6, Pair{X: val1, Y: val2})
	}
	for k := range uniquePosition5 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques5 = append(uniques5, Pair{X: val1, Y: val2})
	}
	for k := range uniquePosition4 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques4 = append(uniques4, Pair{X: val1, Y: val2})
	}
	for k := range uniquePosition3 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques3 = append(uniques3, Pair{X: val1, Y: val2})
	}
	for k := range uniquePosition2 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques2 = append(uniques2, Pair{X: val1, Y: val2})
	}
	for k := range uniquePosition1 {
		parts := strings.Split(k, ",")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		uniques1 = append(uniques1, Pair{X: val1, Y: val2})
	}
	output := ""
	for i := 0; i < 70; i++ {
		for j := -75; j < 75; j++ {
			printed := false
			toAdd := ""
			for _, pair := range uniques9 {
				if j == pair.X && i == pair.Y {
					toAdd = "#"
					printed = true
					break
				}
			}
			for _, pair := range uniques8 {
				if j == pair.X && i == pair.Y {
					toAdd = "8"
					printed = true
					break
				}
			}
			for _, pair := range uniques7 {
				if j == pair.X && i == pair.Y {
					toAdd = "7"
					printed = true
					break
				}
			}
			for _, pair := range uniques6 {
				if j == pair.X && i == pair.Y {
					toAdd = "6"
					printed = true
					break
				}
			}
			for _, pair := range uniques5 {
				if i == pair.Y && j == pair.X {
					toAdd = "5"
					printed = true
					break
				}
			}
			for _, pair := range uniques4 {
				if i == pair.Y && j == pair.X {
					toAdd = "4"
					printed = true
					break
				}
			}
			for _, pair := range uniques3 {
				if i == pair.Y && j == pair.X {
					toAdd = "3"
					printed = true
					break
				}
			}
			for _, pair := range uniques2 {
				if i == pair.Y && j == pair.X {
					toAdd = "2"
					printed = true
					break
				}
			}
			for _, pair := range uniques1 {
				if i == pair.Y && j == pair.X {
					toAdd = "1"
					printed = true
					break
				}
			}
			if !printed {
				output += "."
			} else {
				output += toAdd
			}
		}
		output += "\n"
	}
	time.Sleep(time.Millisecond * 20)
	fmt.Print(output)
}

func (t *Tail) updateHEADPosition(dir string) {
	switch dir {
	case "U":
		t.Y--
	case "R":
		t.X++
	case "D":
		t.Y++
	case "L":
		t.X--
	}
}

func (t *Tail) updatePosition() {
	var Xsubstraction int
	var Ysubstraction int
	Xsubstraction = t.X - t.HEAD.X
	Ysubstraction = t.Y - t.HEAD.Y

	if Xsubstraction == -2 && Ysubstraction == -2 {
		t.X++
		t.Y++
		return
	} else if Xsubstraction == -2 && Ysubstraction == 2 {
		t.X++
		t.Y--
		return
	} else if Xsubstraction == 2 && Ysubstraction == -2 {
		t.X--
		t.Y++
		return
	} else if Xsubstraction == 2 && Ysubstraction == 2 {
		t.X--
		t.Y--
		return
	}

	if Xsubstraction <= -2 {
		t.X++
		if Ysubstraction <= -1 {
			t.Y++
		} else if Ysubstraction >= 1 {
			t.Y--
		}
	} else if Xsubstraction >= 2 {
		t.X--
		if Ysubstraction <= -1 {
			t.Y++
		} else if Ysubstraction >= 1 {
			t.Y--
		}
		return
	}
	if Ysubstraction <= -2 {
		t.Y++
		if Xsubstraction <= -1 {
			t.X++
		} else if Xsubstraction >= 1 {
			t.X--
		}
	} else if Ysubstraction >= 2 {
		t.Y--
		if Xsubstraction <= -1 {
			t.X++
		} else if Xsubstraction >= 1 {
			t.X--
		}
		return
	}

	//switch Xsubstraction {
	//case -2:
	//	t.X++
	//	switch Ysubstraction {
	//	case -1:
	//		t.Y++
	//	case 1:
	//		t.Y--
	//	}
	//	return
	//case 2:
	//	t.X--
	//	switch Ysubstraction {
	//	case -1:
	//		t.Y++
	//	case 1:
	//		t.Y--
	//	}
	//	return
	//}
	//switch Ysubstraction {
	//case -2:
	//	t.Y++
	//	switch Xsubstraction {
	//	case -1:
	//		t.X++
	//	case 1:
	//		t.X--
	//	}
	//	return
	//case 2:
	//	t.Y--
	//	switch Xsubstraction {
	//	case -1:
	//		t.X++
	//	case 1:
	//		t.X--
	//	}
	//	return
	//}

}

func readInput(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	return string(data)
}
