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

	var reg []int = make([]int, 3)
	var program []int = make([]int, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for range 3 {
		scanner.Scan()
		var regVal int
		var regAddr rune
		fmt.Sscanf(scanner.Text(), "Register %c: %d", &regAddr, &regVal)
		switch regAddr {
		case 'A':
			reg[0] = regVal
		case 'B':
			reg[1] = regVal
		case 'C':
			reg[2] = regVal
		}
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

	var output []int = make([]int, len(program))

	runProgram(program, reg, output)

	fmt.Printf("%v\n", output)

	var dynReg []int = make([]int, 3)
	for i := 1; true; i++ {
		dynReg[0] = i
		dynReg[1] = reg[1]
		dynReg[2] = reg[2]

		if !runProgram(program, dynReg, output) {
			continue
		}

		var isSame bool = true
		for j := len(output) - 1; j >= 0; j-- {
			if program[j] != output[j] {
				isSame = false
				break
			}
		}

		if isSame {
			fmt.Println(i)
			break
		} else if i%100000 == 0 {
			fmt.Println(i)
		}
	}
}

func getComboVal(val int, reg []int) int {
	var res int

	if val < 4 {
		res = val
	} else {
		switch val {
		case 4:
			res = reg[0]
		case 5:
			res = reg[1]
		case 6:
			res = reg[2]
		default:
			log.Fatal("Unexpected val!")
		}
	}

	return res
}

func runProgram(program []int, reg []int, output []int) bool {
	for j := range len(output) {
		output[j] = -1
	}

	var currOut int
	var progCounter int

	for progCounter < len(program) {
		var inst int = program[progCounter]
		var op int = program[progCounter+1]

		switch inst {
		case 0:
			reg[0] = reg[0] / (1 << getComboVal(op, reg))
			progCounter += 2
		case 1:
			reg[1] = reg[1] ^ op
			progCounter += 2
		case 2:
			reg[1] = getComboVal(op, reg) % 8
			progCounter += 2
		case 3:
			if reg[0] == 0 {
				progCounter += 2
			} else {
				progCounter = op
			}
		case 4:
			reg[1] = reg[1] ^ reg[2]
			progCounter += 2
		case 5:
			if currOut == len(output) {
				return false
			}
			output[currOut] = getComboVal(op, reg) % 8
			currOut++
			progCounter += 2
		case 6:
			reg[1] = reg[0] / (1 << getComboVal(op, reg))
			progCounter += 2
		case 7:
			reg[2] = reg[0] / (1 << getComboVal(op, reg))
			progCounter += 2
		}
	}

	return true
}
