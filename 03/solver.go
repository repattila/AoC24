package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
)

type condition struct {
	pos     int
	enabled bool
}

func cmpConditions(x condition, y condition) int {
	if x.pos == y.pos {
		return 0
	} else if x.pos < y.pos {
		return -1
	} else {
		return 1
	}
}

func insert(cs []condition, c condition) []condition {
	i, _ := slices.BinarySearchFunc(cs, c, cmpConditions)
	return slices.Insert(cs, i, c)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	mul := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	do := regexp.MustCompile(`do\(\)`)
	dont := regexp.MustCompile(`don\'t\(\)`)
	var res int
	var enabled bool = true
	for scanner.Scan() {
		var line string = scanner.Text()
		match := mul.FindAllStringIndex(line, -1)
		matchDo := do.FindAllStringIndex(line, -1)
		matchDont := dont.FindAllStringIndex(line, -1)

		fmt.Println(len(match))
		fmt.Println(len(matchDo))
		fmt.Println(len(matchDont))

		var conditions []condition
		conditions = append(conditions, condition{0, enabled})

		for _, m := range matchDo {
			fmt.Println(m)

			//fmt.Println(line[m[0]:m[1]])
			conditions = insert(conditions, condition{m[0], true})
		}

		for _, m := range matchDont {
			fmt.Println(m)

			//fmt.Println(line[m[0]:m[1]])
			conditions = insert(conditions, condition{m[0], false})
		}

		fmt.Println(conditions)

		for _, m := range match {
			// Look for the condition that directly preceeds the position of the current mul
			pos, _ := slices.BinarySearchFunc(conditions, condition{m[0], true}, cmpConditions)
			if conditions[pos-1].enabled {
				var mulString string = line[m[0]:m[1]]
				var op1, op2 int
				fmt.Sscanf(mulString, "mul(%d,%d)", &op1, &op2)
				res += op1 * op2
			}
		}

		enabled = conditions[len(conditions)-1].enabled
	}

	fmt.Println(res)
}
