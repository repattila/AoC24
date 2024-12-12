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

	for i := range 25 {
		var newStones []int = make([]int, 0)
		for j, s := range stones {
			if s == 0 {
				stones[j] = 1
			} else {
				engString := strconv.FormatInt(int64(s), 10)
				if len(engString)%2 == 0 {
					newEngraving, _ := strconv.Atoi(engString[:len(engString)/2])
					newStones = append(newStones, newEngraving)

					newEngraving, _ = strconv.Atoi(engString[len(engString)/2:])
					stones[j] = newEngraving
				} else {
					stones[j] = s * 2024
				}
			}
		}

		stones = append(stones, newStones...)
		fmt.Println(i)
	}

	fmt.Printf("%v\n", len(stones))
}
