package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var locks [][]int = make([][]int, 0)
	var keys [][]int = make([][]int, 0)

	var lineNum int
	var isKey bool
	var currPins []int = make([]int, 5)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var lineNum8 int = lineNum % 8
		var line string = scanner.Text()
		if lineNum8 == 6 {
			if isKey {
				keys = append(keys, currPins)
			} else {
				locks = append(locks, currPins)
			}
			currPins = make([]int, 5)
		} else if lineNum8 == 0 {
			if line == "....." {
				isKey = true
			} else {
				isKey = false
			}
		} else if lineNum8 != 7 {
			for i := range 5 {
				if line[i] == '#' {
					currPins[i] += 1
				}
			}
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", locks)
	fmt.Printf("%v\n", keys)

	var res int
	for _, lock := range locks {
		for _, key := range keys {
			var match bool = true
			for i := range 5 {
				if key[i]+lock[i] > 5 {
					match = false
					break
				}
			}
			if match {
				res++
			}
		}
	}

	fmt.Println(res)
}
