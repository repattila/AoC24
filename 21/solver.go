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

	var codeVals []string = make([]string, 0)
	var codes [][]pos = make([][]pos, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var codeVal string = scanner.Text()
		var code []pos = make([]pos, 0)

		for _, r := range codeVal {
			var p pos

			//7, 8, 9
			//4, 5, 6
			//1, 2, 3
			// , 0, A
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
	fmt.Printf("\n")

	var res int
	for _, code := range codes {
		var routeOptions [][]route = make([][]route, 0, len(code))
		// A
		var currPos pos = pos{3, 2}
		for _, p := range code {
			rDiff := currPos.row - p.row
			cDiff := currPos.col - p.col

			var verticalMove rune
			var horizontalMove rune
			if rDiff > 0 {
				verticalMove = '^'
			} else {
				verticalMove = 'v'
				rDiff = -1 * rDiff
			}

			if cDiff < 0 {
				horizontalMove = '>'
				cDiff = -1 * cDiff
			} else {
				horizontalMove = '<'
			}

			routeOptions = append(routeOptions, getRoutes(verticalMove, horizontalMove, []route{route{cDiff, rDiff, make([]rune, 0)}}))

			currPos = p
		}

		fmt.Printf("%v\n", routeOptions)

		var inst1stLevel [][]rune = make([][]rune, 1)
		inst1stLevel[0] = make([]rune, 0)

		for _, routes := range routeOptions {
			var newInst1stLevel [][]rune = make([][]rune, 0)
			for _, inst := range inst1stLevel {
				for _, route := range routes {
					newInst1stLevel = append(newInst1stLevel, append(inst, route.steps...))
				}
			}
			inst1stLevel = newInst1stLevel
		}

		fmt.Printf("%v\n", inst1stLevel)

		var validInst [][]rune = make([][]rune, 0)
	instLoop:
		for _, inst := range inst1stLevel {
			var currPos pos = pos{3, 2}
			for _, step := range inst {
				switch step {
				case '<':
					currPos.col -= 1
				case '>':
					currPos.col += 1
				case '^':
					currPos.row -= 1
				case 'v':
					currPos.row += 1
				}

				if currPos.row == 3 && currPos.col == 0 {
					continue instLoop
				}
			}

			validInst = append(validInst, inst)
		}

		inst1stLevel = validInst

		fmt.Printf("%v\n", inst1stLevel)
		fmt.Printf("\n")

		/*
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
		*/
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

type route struct {
	horizontalStepsRemaining int
	verticalStepsRemaining   int
	steps                    []rune
}

func (r *route) addStep(step rune, isVertical bool) {
	newSteps := make([]rune, 0, len(r.steps)+1)
	newSteps = append(newSteps, r.steps...)
	newSteps = append(newSteps, step)
	r.steps = newSteps

	if isVertical {
		r.verticalStepsRemaining -= 1
	} else {
		r.horizontalStepsRemaining -= 1
	}
}

func (r *route) copyStepsAndAdd(from route, step rune, isVertical bool) {
	newSteps := make([]rune, 0, len(from.steps)+1)
	newSteps = append(newSteps, from.steps...)
	newSteps = append(newSteps, step)
	r.steps = newSteps

	if isVertical {
		r.verticalStepsRemaining = from.verticalStepsRemaining - 1
		r.horizontalStepsRemaining = from.horizontalStepsRemaining
	} else {
		r.verticalStepsRemaining = from.verticalStepsRemaining
		r.horizontalStepsRemaining = from.horizontalStepsRemaining - 1
	}
}

func getRoutes(verticalMove rune, horizontalMove rune, routes []route) []route {
	var newRoutes []route = make([]route, 0)
	var added bool = false

	for _, r := range routes {
		if r.horizontalStepsRemaining == 0 && r.verticalStepsRemaining == 0 {
			newRoutes = append(newRoutes, r)
		} else if r.horizontalStepsRemaining > 0 && r.verticalStepsRemaining > 0 {
			var newRoute route = route{}
			newRoute.copyStepsAndAdd(r, horizontalMove, false)
			newRoutes = append(newRoutes, newRoute)
			r.addStep(verticalMove, true)
			newRoutes = append(newRoutes, r)

			added = true
		} else if r.horizontalStepsRemaining > 0 {
			r.addStep(horizontalMove, false)
			newRoutes = append(newRoutes, r)

			added = true
		} else {
			r.addStep(verticalMove, true)
			newRoutes = append(newRoutes, r)

			added = true
		}
	}

	if !added {
		return routes
	} else {
		return getRoutes(verticalMove, horizontalMove, newRoutes)
	}
}
