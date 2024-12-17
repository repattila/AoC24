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

type route struct {
	dir   int
	cost  int
	steps []pos
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var routes []route = make([]route, 0)

	var field [][]int = make([][]int, 0)
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
				routes = append(routes, route{1, 0, []pos{pos{r, c}}})
			}
		}
		field = append(field, row)
		r++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", field)
	fmt.Printf("%v\n", routes)

	var finishedRoutes []route = make([]route, 0)
	for true {
		var updatedRoutes []route = make([]route, 0, len(routes))

		for _, r := range routes {
			var currPos pos = r.steps[len(r.steps)-1]
			if field[currPos.r][currPos.c] == 2 {
				finishedRoutes = append(finishedRoutes, r)
				continue
			}

			// up
			nextPos := pos{currPos.r - 1, currPos.c}
			if field[nextPos.r][nextPos.c] != 0 && !stepVisited(nextPos, r) {
				switch r.dir {
				// up
				case 0:
					updatedRoutes = append(updatedRoutes, addToRoute(0, r.cost+1, nextPos, r))
				// right, left
				case 1, 3:
					updatedRoutes = append(updatedRoutes, addToRoute(0, r.cost+1001, nextPos, r))
				}
			}
			// right
			nextPos = pos{currPos.r, currPos.c + 1}
			if field[nextPos.r][nextPos.c] != 0 && !stepVisited(nextPos, r) {
				switch r.dir {
				// right
				case 1:
					updatedRoutes = append(updatedRoutes, addToRoute(1, r.cost+1, nextPos, r))
				// up, down
				case 0, 2:
					updatedRoutes = append(updatedRoutes, addToRoute(1, r.cost+1001, nextPos, r))
				}
			}
			// left
			nextPos = pos{currPos.r, currPos.c - 1}
			if field[nextPos.r][nextPos.c] != 0 && !stepVisited(nextPos, r) {
				switch r.dir {
				// left
				case 3:
					updatedRoutes = append(updatedRoutes, addToRoute(3, r.cost+1, nextPos, r))
				// up, down
				case 0, 2:
					updatedRoutes = append(updatedRoutes, addToRoute(3, r.cost+1001, nextPos, r))
				}
			}
			// down
			nextPos = pos{currPos.r + 1, currPos.c}
			if field[nextPos.r][nextPos.c] != 0 && !stepVisited(nextPos, r) {
				switch r.dir {
				// left
				case 2:
					updatedRoutes = append(updatedRoutes, addToRoute(2, r.cost+1, nextPos, r))
				// up, down
				case 1, 3:
					updatedRoutes = append(updatedRoutes, addToRoute(2, r.cost+1001, nextPos, r))
				}
			}
		}

		fmt.Printf("%v\n", updatedRoutes)
		fmt.Printf("%v\n", finishedRoutes)

		if len(updatedRoutes) != 0 {
			routes = updatedRoutes
		} else {
			break
		}
	}

	fmt.Printf("%v\n", finishedRoutes)
}

func stepVisited(step pos, r route) bool {
	var hasStep bool = false
	for _, s := range r.steps {
		if step == s {
			hasStep = true
			break
		}
	}
	return hasStep
}

func addToRoute(dir int, cost int, step pos, r route) route {
	steps := make([]pos, 0, len(r.steps))
	steps = append(steps, r.steps...)
	steps = append(steps, step)
	return route{dir, cost, steps}
}
