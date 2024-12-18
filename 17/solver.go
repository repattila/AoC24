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

	var reg map[rune]int = make(map[rune]int)
	var program []int = make([]int, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for range 3 {
		scanner.Scan()
		var regVal int
		var regAddr rune
		fmt.Sscanf(scanner.Text(), "Register %c: %d", &regAddr, &regVal)
		reg[regAddr] = regVal
	}

	scanner.Scan()
	scanner.Scan()
	var prog string
	fmt.Sscanf(scanner.Text(), "Program: %s", &prog)
	for _, inst := range strings.Split(prog, ",") {
		i, err := strconv.Atoi(inst)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", reg)
	fmt.Printf("%v\n", program)

	var output []int = make([]int, 0)
	var progCounter int
	for progCounter < len(program) {
		var inst int = program[progCounter]
		var op int = program[progCounter+1]

		switch inst {
		case 0:
			reg['A'] = reg['A'] / int(math.Pow(2, float64(getComboVal(op, reg))))
			progCounter += 2
		case 1:
			reg['B'] = reg['B'] ^ op
			progCounter += 2
		case 2:
			reg['B'] = getComboVal(op, reg) % 8
			progCounter += 2
		case 3:
			if reg['A'] == 0 {
				progCounter += 2
			} else {
				progCounter = op
			}
		case 4:
			reg['B'] = reg['B'] ^ reg['C']
			progCounter += 2
		case 5:
			output = append(output, getComboVal(op, reg)%8)
			progCounter += 2
		case 6:
			reg['B'] = reg['A'] / int(math.Pow(2, float64(getComboVal(op, reg))))
			progCounter += 2
		case 7:
			reg['C'] = reg['A'] / int(math.Pow(2, float64(getComboVal(op, reg))))
			progCounter += 2
		}
	}

	for _, out := range output {
		fmt.Printf("%d,", out)
	}
	fmt.Printf("\n")
}

func getComboVal(val int, reg map[rune]int) int {
	var res int

	if val < 4 {
		res = val
	} else {
		switch val {
		case 4:
			res = reg['A']
		case 5:
			res = reg['B']
		case 6:
			res = reg['C']
		default:
			log.Fatal("Unexpected val!")
		}
	}

	return res
}
