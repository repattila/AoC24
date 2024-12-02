package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var safe int = 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
lines:
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		intVals := make([]int, 0)
		for _, s := range vals {
			i, err := strconv.Atoi(s)
			if err != nil {
				continue lines
			} else {
				intVals = append(intVals, i)
			}
		}

		if tryIncDec(intVals) {
			safe += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", safe)
}

func tryIncDec(vals []int) bool {
	if isSafe(vals, true) {
		return true
	} else {
		return isSafe(vals, false)
	}
}

func isSafe(vals []int, inc bool) bool {
	var res bool = true
	var prevVal int = vals[0]
	for i := 1; i < len(vals); i++ {
		var currVal int = vals[i]

		if inc {
			if currVal <= prevVal || currVal-prevVal > 3 {
				res = false
				break
			}
		} else {
			if currVal >= prevVal || prevVal-currVal > 3 {
				res = false
				break
			}
		}

		prevVal = currVal
	}
	return res
}
