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
		fmt.Printf("%v\n", code)

		// A
		var currPos pos = pos{3, 2}
		var routeOptionsByCodePos [][]route = getRouteOptions(currPos, code)

		printRouteOptions(routeOptionsByCodePos)

		var inst1stLevel [][]rune = getInstructions(routeOptionsByCodePos)

		fmt.Printf("%v\n", inst1stLevel)

		inst1stLevel = getValidInst1stLevel(inst1stLevel)

		fmt.Printf("%v\n", inst1stLevel)
		fmt.Printf("\n")

		var inst2ndLevel [][]rune = make([][]rune, 0)
		for i, inst := range inst1stLevel {
			fmt.Printf("2nd level, instruction: %d\n", i)
			var instAsPos []pos = getInstAsPos2ndLevel(inst)
			fmt.Printf("%v\n", instAsPos)

			// A
			var currPos pos = pos{0, 2}
			var routeOptionsByInst [][]route = getRouteOptions(currPos, instAsPos)

			printRouteOptions(routeOptionsByInst)

			var inst2ndLevelBy1stLevelInst [][]rune = getInstructions(routeOptionsByInst)

			fmt.Printf("%v\n", inst2ndLevelBy1stLevelInst)

			inst2ndLevelBy1stLevelInst = getValidInst2ndLevel(inst2ndLevelBy1stLevelInst)

			fmt.Printf("%v\n", inst2ndLevelBy1stLevelInst)

			inst2ndLevel = append(inst2ndLevel, inst2ndLevelBy1stLevelInst...)
			fmt.Printf("\n")
		}
		fmt.Printf("\n")

		var inst3rdLevel [][]rune = make([][]rune, 0)
		for i, inst := range inst2ndLevel {
			fmt.Printf("3rd level, instruction: %d\n", i)
			var instAsPos []pos = getInstAsPos2ndLevel(inst)
			fmt.Printf("%v\n", instAsPos)

			// A
			var currPos pos = pos{0, 2}
			var routeOptionsByInst [][]route = getRouteOptions(currPos, instAsPos)

			printRouteOptions(routeOptionsByInst)

			var inst3rdLevelBy2ndLevelInst [][]rune = getInstructions(routeOptionsByInst)

			fmt.Printf("%v\n", inst3rdLevelBy2ndLevelInst)

			inst3rdLevelBy2ndLevelInst = getValidInst2ndLevel(inst3rdLevelBy2ndLevelInst)

			fmt.Printf("%v\n", inst3rdLevelBy2ndLevelInst)

			inst3rdLevel = append(inst3rdLevel, inst3rdLevelBy2ndLevelInst...)
			fmt.Printf("\n")
		}

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

// Returns the possible sequences of characters that need to be pressed on the keypad
// The return value is a list of lists, representing the options of character sequences
// Uses the previously determined possible routes between the characters of the code
func getInstructions(routeOptionsByCodePos [][]route) [][]rune {
	var instructions [][]rune = make([][]rune, 1)
	instructions[0] = make([]rune, 0)

	// Generate all possible sequences of routes on the keypad between the characters of the code
	// After appending the steps of a route between two codepoints, always add the 'A' character
	for _, routeOptions := range routeOptionsByCodePos {
		var newInstructions [][]rune = make([][]rune, 0)
		for _, inst := range instructions {
			for _, route := range routeOptions {
				newInstructions = append(newInstructions, append(append(inst, route.steps...), 'A'))
			}
		}
		instructions = newInstructions
	}

	return instructions
}

func getRouteOptions(currPos pos, code []pos) [][]route {
	var routeOptionsByCodePos [][]route = make([][]route, 0, len(code))
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

		routeOptionsByCodePos = append(routeOptionsByCodePos, getRoutes(verticalMove, horizontalMove, []route{route{cDiff, rDiff, make([]rune, 0)}}))

		currPos = p
	}
	return routeOptionsByCodePos
}

func getInstAsPos2ndLevel(inst []rune) []pos {
	var res []pos = make([]pos, 0)
	for _, r := range inst {
		// , ^, A
		// <, v, >
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

func getValidInst1stLevel(instructions [][]rune) [][]rune {
	var validInst [][]rune = make([][]rune, 0)
instLoop:
	for _, inst := range instructions {
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
			case 'A':
			}

			if currPos.row == 3 && currPos.col == 0 {
				continue instLoop
			}
		}

		validInst = append(validInst, inst)
	}
	return validInst
}

func getValidInst2ndLevel(instructions [][]rune) [][]rune {
	var validInst [][]rune = make([][]rune, 0)
instLoop:
	for _, inst := range instructions {
		var currPos pos = pos{0, 2}
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
			case 'A':
			}

			if currPos.row == 0 && currPos.col == 0 {
				continue instLoop
			}
		}

		validInst = append(validInst, inst)
	}
	return validInst
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

func (r route) String() string {
	var result string
	for _, s := range r.steps {
		result = result + fmt.Sprintf("%c", s)
	}
	return result
}

func printRouteOptions(routeOptions [][]route) {
	for _, ro := range routeOptions {
		for _, r := range ro {
			fmt.Printf("%v ", r)
		}
		fmt.Printf("\n")
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
