package main

import (
	"fmt"
	"math"
	"sort"
)

type Monkey struct {
	number         int
	inspectedItems int
	StartItems     []uint64
	Operation      func(old uint64) uint64
	Test           func(item uint64) (int, int)
}

const commonMultiply = 5 * 17 * 2 * 7 * 3 * 11 * 13 * 19

func main() {
	var monkeys = createMonkeys()
	PrintMonkeysState(monkeys)
	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			monkey.InspectItems(monkeys)
			//PrintMonkeysState(monkeys)
		}
	}

	PrintMonkeysState(monkeys)
	var allInspections []int
	for _, monkey := range monkeys {
		allInspections = append(allInspections, monkey.inspectedItems)
		fmt.Printf("Monkey %d inspectedItems %d\n", monkey.number, monkey.inspectedItems)
	}
	sort.Ints(allInspections)
	fmt.Println(allInspections)
	fmt.Println(allInspections[len(allInspections)-1] * allInspections[len(allInspections)-2])
}

func (m *Monkey) InspectItems(monkeys []*Monkey) {
	//fmt.Println("Monkey ", m.number)
	for len(m.StartItems) > 0 {
		var currentItem uint64
		if len(m.StartItems) > 0 {
			currentItem = m.StartItems[0]
		}
		m.inspectedItems++
		prevItemVal := currentItem
		currentItem = m.Operation(currentItem)
		if currentItem < prevItemVal {
			fmt.Println("Overflow")
		}
		currentItem = GetBored(currentItem)
		nextThrow, _ := m.Test(currentItem)
		monkeyToGetItem := GetMonkeyToThrow(monkeys, nextThrow)
		if monkeyToGetItem == nil {
			panic("WTF")
		}
		monkeyToGetItem.StartItems = append(monkeyToGetItem.StartItems, currentItem)
		//fmt.Printf("Item with level %d is thrown to moneky %d\n", currentItem, monkeyToGetItem.number)
		m.StartItems = m.StartItems[1:]
	}
}

// 14400119958 too low
// 14400119958 kurwa

func PrintMonkeysState(monkeys []*Monkey) {
	for _, m := range monkeys {
		fmt.Printf("Monkey %d\n\tItems %v\n", m.number, m.StartItems)
	}
}

func GetBored(item uint64) uint64 {
	return item % commonMultiply
}

func GetMonkeyToThrow(monkeys []*Monkey, nextThrow int) *Monkey {
	for _, monkey := range monkeys {
		if monkey.number == nextThrow {
			return monkey
		}
	}
	return nil
}

func createMonkeys() []*Monkey {
	var monkeys []*Monkey
	monkeys = append(monkeys, &Monkey{
		number:         0,
		inspectedItems: 0,
		StartItems:     []uint64{77, 69, 76, 77, 50, 58},
		Operation: func(old uint64) uint64 {
			return old * 11
		},
		Test: func(item uint64) (int, int) {
			if item%5 == 0 {
				return 1, int(math.Floor(float64(item) / 5))
			}
			return 5, int(math.Floor(float64(item) / 5))
		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         1,
		inspectedItems: 0,
		StartItems:     []uint64{75, 70, 82, 83, 96, 64, 62},
		Operation: func(old uint64) uint64 {
			return old + 8
		},
		Test: func(item uint64) (int, int) {
			if item%17 == 0 {
				return 5, int(math.Floor(float64(item) / 17))
			}
			return 6, int(math.Floor(float64(item) / 17))
		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         2,
		inspectedItems: 0,
		StartItems:     []uint64{53},
		Operation: func(old uint64) uint64 {
			return old * 3
		},
		Test: func(item uint64) (int, int) {
			if item%2 == 0 {
				return 0, int(math.Floor(float64(item) / 2))
			}
			return 7, int(math.Floor(float64(item) / 2))
		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         3,
		inspectedItems: 0,
		StartItems:     []uint64{85, 64, 93, 64, 99},
		Operation: func(old uint64) uint64 {
			return old + 4
		},
		Test: func(item uint64) (int, int) {
			if item%7 == 0 {
				return 7, int(math.Floor(float64(item) / 7))
			}
			return 2, int(math.Floor(float64(item) / 7))

		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         4,
		inspectedItems: 0,
		StartItems:     []uint64{61, 92, 71},
		Operation: func(old uint64) uint64 {
			return old * old
		},
		Test: func(item uint64) (int, int) {
			if item%3 == 0 {
				return 2, int(math.Floor(float64(item) / 3))
			}
			return 3, int(math.Floor(float64(item) / 3))
		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         5,
		inspectedItems: 0,
		StartItems:     []uint64{79, 73, 50, 90},
		Operation: func(old uint64) uint64 {
			return old + 2
		},
		Test: func(item uint64) (int, int) {
			if item%11 == 0 {
				return 4, int(math.Floor(float64(item) / 11))
			}
			return 6, int(math.Floor(float64(item) / 11))
		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         6,
		inspectedItems: 0,
		StartItems:     []uint64{50, 89},
		Operation: func(old uint64) uint64 {
			return old + 3
		},
		Test: func(item uint64) (int, int) {
			if item%13 == 0 {
				return 4, int(math.Floor(float64(item) / 13))
			}
			return 3, int(math.Floor(float64(item) / 13))
		},
	})
	monkeys = append(monkeys, &Monkey{
		number:         7,
		inspectedItems: 0,
		StartItems:     []uint64{83, 56, 64, 58, 93, 91, 56, 65},
		Operation: func(old uint64) uint64 {
			return old + 5
		},
		Test: func(item uint64) (int, int) {
			if item%19 == 0 {
				return 1, int(math.Floor(float64(item) / 19))
			}
			return 0, int(math.Floor(float64(item) / 19))
		},
	})
	return monkeys
}
