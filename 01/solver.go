package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var firstList []int = make([]int, 1)
	var secondList []int = make([]int, 1)
	var init bool = true
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var first, second int
		fmt.Sscanf(scanner.Text(), "%d%d", &first, &second)

		fmt.Printf("%v %v\n", first, second)

		if init {
			firstList[0] = first
			secondList[0] = second
			init = false
		} else {
			firstList = insert(firstList, first)
			secondList = insert(secondList, second)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result1 int = 0
	for i := 0; i < len(firstList); i++ {
		result1 += abs(firstList[i] - secondList[i])
	}

	var result2 int = 0
	for i := 0; i < len(firstList); i++ {
		var currVal int = firstList[i]
		pos, found := slices.BinarySearch(secondList, currVal) // find slot
		if found {
			var count int = 0
			for ; secondList[pos] == currVal; pos++ {
				count++
			}
			result2 += currVal * count
		}
	}

	fmt.Printf("%v\n", len(firstList))
	fmt.Printf("%v\n", result1)
	fmt.Printf("%v\n", result2)
}

func insert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}
