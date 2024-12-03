package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	var res int
	for scanner.Scan() {
		var line string = scanner.Text()
		match := re.FindAllString(line, -1)

		fmt.Println(len(match))

		for _, m := range match {
			fmt.Println(m)

			var op1, op2 int
			fmt.Sscanf(m, "mul(%d,%d)", &op1, &op2)
			res += op1 * op2
		}
	}

	fmt.Println(res)
}
