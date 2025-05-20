package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	splitLine := strings.Split(scanner.Text(), " ")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var stones []int = make([]int, 0, 80028872400)
	for _, s := range splitLine {
		i, _ := strconv.Atoi(s)
		stones = append(stones, i)
	}

	start := time.Now()

	var lenGeneratedFromDigit map[int]map[int]int = make(map[int]map[int]int)
	var maxDigit int = 10
	for n := range maxDigit {
		var lenGeneratedFromCurr map[int]int = make(map[int]int)
		var stonesFromDigit []int = make([]int, 0, 1000000)
		stonesFromDigit = append(stonesFromDigit, n)

		for i := range 40 {
			for j := range len(stonesFromDigit) {
				var s int = stonesFromDigit[j]
				if s == 0 {
					stonesFromDigit[j] = 1
				} else {
					engString := strconv.FormatInt(int64(s), 10)
					if len(engString)%2 == 0 {
						newEngraving, _ := strconv.Atoi(engString[:len(engString)/2])
						stonesFromDigit = append(stonesFromDigit, newEngraving)

						newEngraving, _ = strconv.Atoi(engString[len(engString)/2:])
						stonesFromDigit[j] = newEngraving
					} else {
						stonesFromDigit[j] = s * 2024
					}
				}
			}

			lenGeneratedFromCurr[i] = len(stonesFromDigit)
		}

		lenGeneratedFromDigit[n] = lenGeneratedFromCurr

		fmt.Printf("Digit done:%v\n", n)
	}

	var stepNum int = 75
	var lenFromSkipped int = 0

	for i := range stepNum {
		for j, s := range stones {
			if s == -1 {
				// ignore
			} else if s < maxDigit {
				var skipped bool = false
				lengeneratedFromDigit, ok1 := lenGeneratedFromDigit[s]
				if ok1 {
					generatedLen, ok2 := lengeneratedFromDigit[stepNum-1-i]

					if ok2 {
						lenFromSkipped += generatedLen
						stones[j] = -1

						skipped = true
					}
				}

				if !skipped {
					if s == 0 {
						stones[j] = 1
					} else {
						stones[j] = s * 2024
					}
				}
			} else {
				engString := strconv.FormatInt(int64(s), 10)
				if len(engString)%2 == 0 {
					newEngraving, _ := strconv.Atoi(engString[:len(engString)/2])
					stones = append(stones, newEngraving)

					newEngraving, _ = strconv.Atoi(engString[len(engString)/2:])
					stones[j] = newEngraving
				} else {
					stones[j] = s * 2024
				}
			}
		}

		fmt.Println(i)
	}

	lenStones := 0
	for _, s := range stones {
		if s != -1 {
			lenStones += 1
		}
	}

	fmt.Printf("%v\n", len(stones))
	fmt.Printf("%v\n", lenStones)
	fmt.Printf("%v\n", lenFromSkipped)
	fmt.Printf("%v\n", lenStones+lenFromSkipped)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
