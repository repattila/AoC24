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

	var safeCount int = 0
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

		safe, pos := tryIncDec(intVals)
		if safe {
			safeCount += 1
		} else {
			if pos != -1 {
				intValsDampened := removeIndex(intVals, pos)
				safe, _ = tryIncDec(intValsDampened)
				if safe {
					safeCount += 1
				} else {
					intValsDampened = removeIndex(intVals, pos-1)
					safe, _ := tryIncDec(intValsDampened)
					if safe {
						safeCount += 1
					}
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", safeCount)
}

func tryIncDec(vals []int) (bool, int) {
	var prevVal int = vals[0]
	var eqCount int = 0
	var incCount int = 0
	var currVal int
	for i := 1; i < 4; i++ {
		currVal = vals[i]

		if currVal == prevVal {
			eqCount += 1
		} else if currVal > prevVal {
			incCount += 1
		}

		prevVal = currVal
	}

	var inc bool
	if eqCount > 1 {
		return false, -1
	} else if incCount > 1 {
		inc = true
	} else {
		inc = false
	}

	return isSafe(vals, inc)
}

func isSafe(vals []int, inc bool) (bool, int) {
	var res bool = true
	var pos int = 1
	var prevVal int = vals[0]
	for ; pos < len(vals); pos++ {
		var currVal int = vals[pos]

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
	return res, pos
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
