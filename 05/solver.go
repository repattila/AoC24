package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	updates := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	var readingRules bool = true
	var rules map[string]bool = make(map[string]bool)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line string
		fmt.Sscanf(scanner.Text(), "%s", &line)

		if line == "" {
			readingRules = false
		} else if readingRules {
			rules[line] = true
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var res1 int
updatesLoop:
	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			if _, ok := rules[update[i]+"|"+update[i+1]]; !ok {
				continue updatesLoop
			}
		}

		intVal, _ := strconv.Atoi(update[len(update)/2])
		res1 += intVal
	}

	var res2 int
	for _, update := range updates {
		for i := 0; i < len(update)-1; i++ {
			if _, ok := rules[update[i]+"|"+update[i+1]]; !ok {
				slices.SortFunc(update, func(a, b string) int {
					if _, ok := rules[a+"|"+b]; ok {
						return -1
					} else {
						return 1
					}
				})
				intVal, _ := strconv.Atoi(update[len(update)/2])
				res2 += intVal
			}
		}
	}

	fmt.Printf("%d\n", res1)
	fmt.Printf("%d\n", res2)
}
