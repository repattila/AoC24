package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type pos struct {
	r int
	c int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var field [][]int = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var row []int = make([]int, 0)
		for _, r := range scanner.Text() {
			i, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			row = append(row, i)
		}
		field = append(field, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", field)

	var scoreSum int
	var ratingSum int
	for r, row := range field {
		for c, fieldVal := range row {
			if fieldVal == 0 {
				fmt.Printf("%d, %d\n", r, c)

				var nines map[pos]bool = make(map[pos]bool)
				followRoute(0, r, c, field, nines)
				rating := getRating(0, r, c, field)

				fmt.Printf("%d\n", nines)
				fmt.Printf("%d\n", rating)

				scoreSum += len(nines)
				ratingSum += rating
			}
		}
	}

	fmt.Println(scoreSum)
	fmt.Println(ratingSum)
}

func followRoute(currVal int, r int, c int, field [][]int, nines map[pos]bool) {
	if currVal == 9 {
		nines[pos{r, c}] = true
	} else {
		var nextVal = currVal + 1
		// up
		if r > 0 && field[r-1][c] == nextVal {
			followRoute(nextVal, r-1, c, field, nines)
		}
		// down
		if r < len(field)-1 && field[r+1][c] == nextVal {
			followRoute(nextVal, r+1, c, field, nines)
		}
		// left
		if c > 0 && field[r][c-1] == nextVal {
			followRoute(nextVal, r, c-1, field, nines)
		}
		// right
		if c < len(field[r])-1 && field[r][c+1] == nextVal {
			followRoute(nextVal, r, c+1, field, nines)
		}
	}
}

func getRating(currVal int, r int, c int, field [][]int) int {
	if currVal == 9 {
		return 1
	} else {
		var res int
		var nextVal = currVal + 1
		// up
		if r > 0 && field[r-1][c] == nextVal {
			res += getRating(nextVal, r-1, c, field)
		}
		// down
		if r < len(field)-1 && field[r+1][c] == nextVal {
			res += getRating(nextVal, r+1, c, field)
		}
		// left
		if c > 0 && field[r][c-1] == nextVal {
			res += getRating(nextVal, r, c-1, field)
		}
		// right
		if c < len(field[r])-1 && field[r][c+1] == nextVal {
			res += getRating(nextVal, r, c+1, field)
		}
		return res
	}
}
