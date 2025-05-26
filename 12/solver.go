package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type pos struct {
	row int
	col int
}

type HSide []pos
type VSide []pos

func (s HSide) Len() int {
	return len(s)
}

// Implement the Less method required by sort.Interface
func (s HSide) Less(i, j int) bool {
	if s[i].row == s[j].row {
		return s[i].col < s[j].col
	} else {
		return s[i].row < s[j].row
	}
}

// Implement the Swap method required by sort.Interface
func (s HSide) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s VSide) Len() int {
	return len(s)
}

// Implement the Less method required by sort.Interface
func (s VSide) Less(i, j int) bool {
	if s[i].col == s[j].col {
		return s[i].row < s[j].row
	} else {
		return s[i].col < s[j].col
	}
}

// Implement the Swap method required by sort.Interface
func (s VSide) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
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
	var res1 int
	for _, region := range regions {
		var regionId rune = field[region[0].row][region[0].col]

		// Fence on the upper side of the pos
		var horizontalSides HSide = make([]pos, 0)
		// Fence on the left side of the pos
		var verticalSides VSide = make([]pos, 0)
		var fenceNum int

		for _, cell := range region {
			r := cell.row
			c := cell.col

			if r == 0 || field[r-1][c] != regionId {
				fenceNum += 1

				horizontalSides = append(horizontalSides, pos{r, c})
			}
			if r == len(field)-1 || field[r+1][c] != regionId {
				fenceNum += 1

				// This needs to be shifted by len(field) so upper and lower sides are not on the same row (test4)
				horizontalSides = append(horizontalSides, pos{r + 1 + len(field), c})
			}
			if c == 0 || field[r][c-1] != regionId {
				fenceNum += 1

				verticalSides = append(verticalSides, pos{r, c})
			}
			if c == len(field[r])-1 || field[r][c+1] != regionId {
				fenceNum += 1

				// This needs to be shifted by len(field[r]) so left and right sides are not on the same row (test4)
				verticalSides = append(verticalSides, pos{r, c + 1 + len(field[r])})
			}
		}

		res += fenceNum * len(region)

		sort.Sort(horizontalSides)
		sort.Sort(verticalSides)

		fmt.Printf("HS: %v\n", horizontalSides)

		var hSideNum int = 0
		var currRow int = -1
		var currCol int = -1
		for _, p := range horizontalSides {
			if currRow != -1 {
				if p.row != currRow {
					hSideNum += 1
				} else if p.col != (currCol + 1) {
					hSideNum += 1
				}
			} else {
				hSideNum += 1
			}

			currRow = p.row
			currCol = p.col
		}

		fmt.Printf("HSNum: %v\n", hSideNum)

		fmt.Printf("VS: %v\n", verticalSides)

		var vSideNum int = 0
		currRow = -1
		currCol = -1
		for _, p := range verticalSides {
			if currCol != -1 {
				if p.col != currCol {
					vSideNum += 1
				} else if p.row != (currRow + 1) {
					vSideNum += 1
				}
			} else {
				vSideNum += 1
			}

			currRow = p.row
			currCol = p.col
		}

		fmt.Printf("VSNum: %v\n", vSideNum)

		res1 += (hSideNum + vSideNum) * len(region)
	}

	//fmt.Printf("%v\n", field)
	//fmt.Printf("%v\n", used)
	//fmt.Printf("%v\n", regions)
	fmt.Printf("%v\n", res)
	fmt.Printf("%v\n", res1)
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
