package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	row int
	col int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var field [][]rune = make([][]rune, 0)
	var used [][]bool = make([][]bool, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line string = scanner.Text()

		field = append(field, []rune(line))
		used = append(used, make([]bool, len(line)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var regions [][]pos = make([][]pos, 0)
	for r, row := range field {
		for c, regionId := range row {
			if !used[r][c] {
				var elems []pos = make([]pos, 0)
				elems = traceRegion(r, c, regionId, field, used, elems)
				regions = append(regions, elems)
			}
		}
	}

	var res int
	for _, region := range regions {
		var regionId rune = field[region[0].row][region[0].col]

		var fenceNum int
		for _, cell := range region {
			r := cell.row
			c := cell.col

			if r == 0 || field[r-1][c] != regionId {
				fenceNum += 1
			}
			if r == len(field)-1 || field[r+1][c] != regionId {
				fenceNum += 1
			}
			if c == 0 || field[r][c-1] != regionId {
				fenceNum += 1
			}
			if c == len(field[r])-1 || field[r][c+1] != regionId {
				fenceNum += 1
			}
		}

		res += fenceNum * len(region)
	}

	fmt.Printf("%v\n", field)
	fmt.Printf("%v\n", used)
	fmt.Printf("%v\n", regions)
	fmt.Printf("%v\n", res)
}

func traceRegion(r int, c int, regionId rune, field [][]rune, used [][]bool, elems []pos) []pos {
	used[r][c] = true
	elems = append(elems, pos{r, c})

	if r > 0 && !used[r-1][c] && field[r-1][c] == regionId {
		elems = traceRegion(r-1, c, regionId, field, used, elems)
	}
	if c > 0 && !used[r][c-1] && field[r][c-1] == regionId {
		elems = traceRegion(r, c-1, regionId, field, used, elems)
	}
	if r < len(field)-1 && !used[r+1][c] && field[r+1][c] == regionId {
		elems = traceRegion(r+1, c, regionId, field, used, elems)
	}
	if c < len(field[r])-1 && !used[r][c+1] && field[r][c+1] == regionId {
		elems = traceRegion(r, c+1, regionId, field, used, elems)
	}

	return elems
}
