package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words [][]rune = make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line string = scanner.Text()
		words = append(words, []rune(line))
	}

	var res int = 0
	for r := 0; r < len(words); r++ {
		row := words[r]
		for c := 0; c < len(row); c++ {
			if row[c] == 'X' {
				res += countXMAS(r, c, words)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func countXMAS(r int, c int, words [][]rune) int {
	var res int = 0

	row := words[r]
	// right
	if c <= len(row)-4 {
		if string(row[c:c+4]) == "XMAS" {
			res++
		}
	}

	// left
	if c >= 3 {
		if string(row[c-3:c+1]) == "SAMX" {
			res++
		}
	}

	// down
	if r <= len(words)-4 {
		word := make([]rune, 0)
		for _, row := range words[r : r+4] {
			word = append(word, row[c])
		}
		if string(word) == "XMAS" {
			res++
		}
	}

	// up
	if r >= 3 {
		word := make([]rune, 0)
		for _, row := range words[r-3 : r+1] {
			word = append(word, row[c])
		}
		if string(word) == "SAMX" {
			res++
		}
	}

	// up right
	if r >= 3 && c <= len(words[r])-4 {
		word := make([]rune, 0)
		var currCol int = c + 3
		for _, row := range words[r-3 : r+1] {
			word = append(word, row[currCol])
			currCol--
		}
		if string(word) == "SAMX" {
			res++
		}
	}

	// down rigth
	if r <= len(words)-4 && c <= len(words[r])-4 {
		word := make([]rune, 0)
		var currCol int = c
		for _, row := range words[r : r+4] {
			word = append(word, row[currCol])
			currCol++
		}
		if string(word) == "XMAS" {
			res++
		}
	}

	// up left
	if r >= 3 && c >= 3 {
		word := make([]rune, 0)
		var currCol int = c - 3
		for _, row := range words[r-3 : r+1] {
			word = append(word, row[currCol])
			currCol++
		}
		if string(word) == "SAMX" {
			res++
		}
	}

	// down left
	if r <= len(words)-4 && c >= 3 {
		word := make([]rune, 0)
		var currCol int = c
		for _, row := range words[r : r+4] {
			word = append(word, row[currCol])
			currCol--
		}
		if string(word) == "XMAS" {
			res++
		}
	}

	return res
}
