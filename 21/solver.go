package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type pos struct {
	row int
	col int
}

func main() {
	file, err := os.Open("example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var codeVals []string = make([]string, 0)
	var codes [][]pos = make([][]pos, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var codeVal string = scanner.Text()
		var code []pos = make([]pos, 0)

		for _, r := range codeVal {
			var p pos

			switch r {
			case 'A':
				p = pos{3, 2}
			case '0':
				p = pos{3, 1}
			case '1':
				p = pos{2, 0}
			case '2':
				p = pos{2, 1}
			case '3':
				p = pos{2, 2}
			case '4':
				p = pos{1, 0}
			case '5':
				p = pos{1, 1}
			case '6':
				p = pos{1, 2}
			case '7':
				p = pos{0, 0}
			case '8':
				p = pos{0, 1}
			case '9':
				p = pos{0, 2}
			}

			code = append(code, p)
		}

		codeVals = append(codeVals, codeVal)
		codes = append(codes, code)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", codeVals)
	fmt.Printf("%v\n", codes)

	var res int
	for i, code := range codes {
		var inst []rune = make([]rune, 0)
		// A
		var currPos pos = pos{3, 2}
		var colDone bool = false
		for _, p := range code {
			rDiff := currPos.row - p.row
			cDiff := currPos.col - p.col

			if rDiff > 0 {
				for range rDiff {
					inst = append(inst, '^')
				}
			} else if rDiff < 0 {
				if currPos.col == 0 {
					if cDiff < 0 {
						for range -1 * cDiff {
							inst = append(inst, '>')
						}
						colDone = true
					}
				}

				for range -1 * rDiff {
					inst = append(inst, 'v')
				}
			}

			if !colDone {
				if cDiff < 0 {
					for range -1 * cDiff {
						inst = append(inst, '>')
					}
				} else if cDiff > 0 {
					for range cDiff {
						inst = append(inst, '<')
					}
				}
			}

			inst = append(inst, 'A')
			currPos = p
			colDone = false
		}

		for _, r := range inst {
			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")

		var instAsPos []pos = getInstAsPos(inst)
		var inst2ndLevel []rune = getInst2ndLevel(instAsPos)

		for _, r := range inst2ndLevel {
			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")

		instAsPos = getInstAsPos(inst2ndLevel)
		var inst3rdLevel []rune = getInst2ndLevel(instAsPos)

		for _, r := range inst3rdLevel {
			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")

		val, _ := strconv.Atoi(codeVals[i][:len(codeVals[i])-1])
		fmt.Printf("%d * %d\n", val, len(inst3rdLevel))

		res += val * len(inst3rdLevel)
	}

	fmt.Println(res)
}

func getInstAsPos(inst []rune) []pos {
	var res []pos = make([]pos, 0)
	for _, r := range inst {
		switch r {
		case 'A':
			res = append(res, pos{0, 2})
		case '^':
			res = append(res, pos{0, 1})
		case '<':
			res = append(res, pos{1, 0})
		case 'v':
			res = append(res, pos{1, 1})
		case '>':
			res = append(res, pos{1, 2})
		}
	}
	return res
}

func getInst2ndLevel(inst []pos) []rune {
	var res []rune = make([]rune, 0)
	// A
	var currPos pos = pos{0, 2}
	var colDone bool = false
	for _, p := range inst {
		rDiff := currPos.row - p.row
		cDiff := currPos.col - p.col

		if rDiff < 0 {
			for range -1 * rDiff {
				res = append(res, 'v')
			}
		} else if rDiff > 0 {
			if currPos.col == 0 {
				if cDiff < 0 {
					for range -1 * cDiff {
						res = append(res, '>')
						colDone = true
					}
				}
			}

			for range rDiff {
				res = append(res, '^')
			}
		}

		if !colDone {
			if cDiff < 0 {
				for range -1 * cDiff {
					res = append(res, '>')
				}
			} else if cDiff > 0 {
				for range cDiff {
					res = append(res, '<')
				}
			}
		}

		res = append(res, 'A')
		currPos = p
		colDone = false
	}
	return res
}
