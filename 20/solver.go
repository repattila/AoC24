package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	r int
	c int
}

type cheatPos struct {
	start pos
	end   pos
}

type route struct {
	cheatUsed bool
	cp        cheatPos
	cost      int
	steps     []pos
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var field [][]int = make([][]int, 0)
	var startPos pos
	var r int
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		var row []int = make([]int, 0)
		for c, val := range line {
			if val == '#' {
				row = append(row, 0)
			} else if val == '.' {
				row = append(row, 1)
			} else if val == 'E' {
				row = append(row, 2)
			} else if val == 'S' {
				row = append(row, 1)
				startPos = pos{r, c}
			}
		}
		field = append(field, row)
		r++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", field)

	var excludedCheats []cheatPos = make([]cheatPos, 0)
	var routes []route = []route{route{cheatUsed: false, cost: 0, steps: []pos{startPos}}}
	noCheatRoute := getRoutesToExit(field, routes, false, excludedCheats)
	noCheatRouteCost := noCheatRoute[0].cost

	fmt.Printf("%v\n", noCheatRouteCost)

	var saving100Routes int
	for true {
		cheatRoutes := getRoutesToExit(field, routes, true, excludedCheats)
		if len(cheatRoutes) == 0 {
			break
		} else {
			for _, chr := range cheatRoutes {
				excludedCheats = append(excludedCheats, chr.cp)

				if chr.cost <= noCheatRouteCost-100 {
					saving100Routes++
					fmt.Printf("%d\n", saving100Routes)
				}
			}
		}
	}

	fmt.Printf("%v\n", saving100Routes)

}

func processNextPos(r route, nextPos pos, usingCheat bool, cp cheatPos, field [][]int, minCostAtPos map[pos]int, updatedRoutes []route) []route {
	var nextCost int
	if usingCheat {
		nextCost = r.cost + 2
	} else {
		nextCost = r.cost + 1
	}

	minCost, ok := minCostAtPos[nextPos]
	if !ok || minCost > nextCost {
		minCostAtPos[nextPos] = nextCost

		newSteps := make([]pos, 0, len(r.steps)+1)
		newSteps = append(newSteps, r.steps...)
		newSteps = append(newSteps, nextPos)
		if usingCheat {
			updatedRoutes = append(updatedRoutes, route{usingCheat, cp, nextCost, newSteps})
		} else {
			updatedRoutes = append(updatedRoutes, route{r.cheatUsed, r.cp, nextCost, newSteps})
		}
	}

	return updatedRoutes
}

func containsCheat(chs []cheatPos, ch cheatPos) bool {
	var res bool = false
	for _, och := range chs {
		if och == ch {
			res = true
			break
		}
	}
	return res
}

func getRoutesToExit(field [][]int, routes []route, cheatEnabled bool, excludedCheats []cheatPos) []route {
	var finishedRoutes []route = make([]route, 0)
	var minCostAtPos map[pos]int = make(map[pos]int)
	for true {
		var updatedRoutes []route = make([]route, 0, len(routes))

		for _, r := range routes {
			currPos := r.steps[len(r.steps)-1]

			if field[currPos.r][currPos.c] == 2 {
				finishedRoutes = append(finishedRoutes, r)
			}

			// up
			nextPos := pos{currPos.r - 1, currPos.c}
			if field[nextPos.r][nextPos.c] != 0 {
				updatedRoutes = processNextPos(r, nextPos, false, cheatPos{}, field, minCostAtPos, updatedRoutes)
			} else if cheatEnabled && !r.cheatUsed {
				if currPos.r-2 >= 0 {
					cp := cheatPos{pos{currPos.r - 1, currPos.c}, pos{currPos.r - 2, currPos.c}}
					if !containsCheat(excludedCheats, cp) {
						nextPos = pos{currPos.r - 2, currPos.c}
						if field[nextPos.r][nextPos.c] != 0 {
							updatedRoutes = processNextPos(r, nextPos, true, cp, field, minCostAtPos, updatedRoutes)
						}
					}
				}
			}

			// right
			nextPos = pos{currPos.r, currPos.c + 1}
			if field[nextPos.r][nextPos.c] != 0 {
				updatedRoutes = processNextPos(r, nextPos, false, cheatPos{}, field, minCostAtPos, updatedRoutes)
			} else if cheatEnabled && !r.cheatUsed {
				if currPos.c+2 < len(field[currPos.r]) {
					cp := cheatPos{pos{currPos.r, currPos.c + 1}, pos{currPos.r, currPos.c + 2}}
					if !containsCheat(excludedCheats, cp) {
						nextPos = pos{currPos.r, currPos.c + 2}
						if field[nextPos.r][nextPos.c] != 0 {
							updatedRoutes = processNextPos(r, nextPos, true, cp, field, minCostAtPos, updatedRoutes)
						}
					}
				}
			}

			// left
			nextPos = pos{currPos.r, currPos.c - 1}
			if field[nextPos.r][nextPos.c] != 0 {
				updatedRoutes = processNextPos(r, nextPos, false, cheatPos{}, field, minCostAtPos, updatedRoutes)
			} else if cheatEnabled && !r.cheatUsed {
				if currPos.c-2 >= 0 {
					cp := cheatPos{pos{currPos.r, currPos.c - 1}, pos{currPos.r, currPos.c - 2}}
					if !containsCheat(excludedCheats, cp) {
						nextPos = pos{currPos.r, currPos.c - 2}
						if field[nextPos.r][nextPos.c] != 0 {
							updatedRoutes = processNextPos(r, nextPos, true, cp, field, minCostAtPos, updatedRoutes)
						}
					}
				}
			}

			// down
			nextPos = pos{currPos.r + 1, currPos.c}
			if field[nextPos.r][nextPos.c] != 0 {
				updatedRoutes = processNextPos(r, nextPos, false, cheatPos{}, field, minCostAtPos, updatedRoutes)
			} else if cheatEnabled && !r.cheatUsed {
				if currPos.r+2 < len(field) {
					cp := cheatPos{pos{currPos.r + 1, currPos.c}, pos{currPos.r + 2, currPos.c}}
					if !containsCheat(excludedCheats, cp) {
						nextPos = pos{currPos.r + 2, currPos.c}
						if field[nextPos.r][nextPos.c] != 0 {
							updatedRoutes = processNextPos(r, nextPos, true, cp, field, minCostAtPos, updatedRoutes)
						}
					}
				}
			}
		}

		if len(updatedRoutes) != 0 {
			routes = updatedRoutes
		} else {
			break
		}
	}

	return finishedRoutes
}
