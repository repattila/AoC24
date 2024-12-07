package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var res int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var testVal int
		var eq string
		splitLine := strings.Split(scanner.Text(), ": ")
		testVal, _ = strconv.Atoi(splitLine[0])
		eq = splitLine[1]

		operands := strings.Split(eq, " ")

		fmt.Printf("%d %v\n", testVal, operands)

		// Operators represented as a base 3 number where 2 is ||, 1 is + and 0 is *
		// 11 = 102 : +*||
		var opCount int = len(operands) - 1
		for opsInt := int64(math.Pow(3, float64(opCount))) - 1; opsInt >= 0; opsInt-- {
			// Convert to base 3 format string
			opsString := strconv.FormatInt(opsInt, 3)
			// Leftpad with zeros
			opsString = fmt.Sprintf("%0*s", opCount, opsString)

			if evalEq(operands, []rune(opsString)) == testVal {
				res += testVal

				fmt.Println(opsString)

				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func evalEq(operands []string, operators []rune) int {
	res, err := strconv.Atoi(operands[0])
	if err != nil {
		return -1
	}

	for i := 1; i < len(operands); i++ {
		operand, err := strconv.Atoi(operands[i])
		if err != nil {
			return -1
		}

		operator := operators[i-1]
		if operator == '0' {
			res *= operand
		} else if operator == '1' {
			res += operand
		} else {
			res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, operand))
		}
	}

	return res
}
