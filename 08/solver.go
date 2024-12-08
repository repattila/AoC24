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

func (p *pos) isValid(maxRow int, maxCol int) bool {
	return p.row >= 0 && p.row < maxRow && p.col >= 0 && p.col < maxCol
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var antennas [][]rune = make([][]rune, 0)
	var antinodes [][]bool = make([][]bool, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := []rune(scanner.Text())
		antennas = append(antennas, line)
		antinodes = append(antinodes, make([]bool, len(line)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", antennas)

	var processedFreqs map[rune]bool = make(map[rune]bool)
	for _, row := range antennas {
		for _, f := range row {
			if f != '.' {
				_, ok := processedFreqs[f]
				if !ok {
					processedFreqs[f] = true

					sameFreqAs := getSameFreqAs(f, antennas)

					fmt.Printf("%c: %v\n", f, sameFreqAs)

					for i, a1 := range sameFreqAs {
						for j := i + 1; j < len(sameFreqAs); j++ {
							antisPos := getAntinodes(a1, sameFreqAs[j], len(antinodes), len(antinodes[0]))
							for _, antiPos := range antisPos {
								antinodes[antiPos.row][antiPos.col] = true
							}
						}
					}
				}
			}
		}
	}

	var res int
	for _, row := range antinodes {
		for _, an := range row {
			if an {
				fmt.Print("#")
				res++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println(res)
}

func getSameFreqAs(freq rune, antennas [][]rune) []pos {
	var res []pos = make([]pos, 0)
	for r, row := range antennas {
		for c, f := range row {
			if f == freq {
				res = append(res, pos{r, c})
			}
		}
	}
	return res
}

func getAntinodes1(a1 pos, a2 pos) (pos, pos) {
	var res1 pos
	var res2 pos

	rDiff := abs(a1.row - a2.row)
	cDiff := abs(a1.col - a2.col)

	if a1.row == a2.row {
		if a1.col < a2.col {
			res1 = pos{a1.row, a1.col - cDiff}
			res2 = pos{a1.row, a2.col + cDiff}
		} else {
			res1 = pos{a2.row, a2.col - cDiff}
			res2 = pos{a2.row, a1.col + cDiff}
		}
	} else if a1.row < a2.row {
		if a1.col == a2.col {
			res1 = pos{a1.row - rDiff, a1.col}
			res2 = pos{a2.row + rDiff, a1.col}
		} else if a1.col < a2.col {
			res1 = pos{a1.row - rDiff, a1.col - cDiff}
			res2 = pos{a2.row + rDiff, a2.col + cDiff}
		} else {
			res1 = pos{a1.row - rDiff, a1.col + cDiff}
			res2 = pos{a2.row + rDiff, a2.col - cDiff}
		}
	} else {
		if a1.col == a2.col {
			res1 = pos{a2.row - rDiff, a2.col}
			res2 = pos{a1.row + rDiff, a2.col}
		} else if a1.col < a2.col {
			res1 = pos{a2.row - rDiff, a2.col + cDiff}
			res2 = pos{a1.row + rDiff, a1.col - cDiff}
		} else {
			res1 = pos{a2.row - rDiff, a2.col - cDiff}
			res2 = pos{a2.row + rDiff, a1.col + cDiff}
		}
	}

	return res1, res2
}

func getAntinodes(a1 pos, a2 pos, maxRow int, maxCol int) []pos {
	var res []pos = make([]pos, 0)

	rDiff := abs(a1.row - a2.row)
	cDiff := abs(a1.col - a2.col)

	if a1.row == a2.row {
		for anti := a1; anti.isValid(maxRow, maxCol); {
			res = append(res, anti)

			anti = pos{anti.row, anti.col + cDiff}
		}
		for anti := a1; anti.isValid(maxRow, maxCol); {
			res = append(res, anti)

			anti = pos{anti.row, anti.col - cDiff}
		}
	} else if a1.col == a2.col {
		for anti := a1; anti.isValid(maxRow, maxCol); {
			res = append(res, anti)

			anti = pos{anti.row + rDiff, anti.col}
		}
		for anti := a1; anti.isValid(maxRow, maxCol); {
			res = append(res, anti)

			anti = pos{anti.row - rDiff, anti.col}
		}
	} else if a1.row < a2.row {
		if a1.col < a2.col {
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row + rDiff, anti.col + cDiff}
			}
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row - rDiff, anti.col - cDiff}
			}
		} else {
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row - rDiff, anti.col + cDiff}
			}
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row + rDiff, anti.col - cDiff}
			}
		}
	} else {
		if a1.col < a2.col {
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row - rDiff, anti.col + cDiff}
			}
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row + rDiff, anti.col - cDiff}
			}
		} else {
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row + rDiff, anti.col + cDiff}
			}
			for anti := a1; anti.isValid(maxRow, maxCol); {
				res = append(res, anti)

				anti = pos{anti.row - rDiff, anti.col - cDiff}
			}
		}
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}
