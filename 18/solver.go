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
	head pos
	len  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var field [][]int = make([][]int, 0)
	for range 71 {
		field = append(field, make([]int, 71))
	}

	var b int
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if b == 1024 {
			break
		}

		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		field[y][x] = 1

		b++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var routes []route = make([]route, 0)
	routes = append(routes, route{pos{0, 0}, 0})

	fmt.Printf("%v\n", field)
	fmt.Printf("%v\n", routes)

	var minStepsAtPos map[pos]int = make(map[pos]int)
	for true {
		var updatedRoutes []route = make([]route, 0, len(routes))

		for _, r := range routes {
			currPos := r.head

			if currPos.r == 70 && currPos.c == 70 {
				continue
			}

			// up
			if currPos.r > 0 {
				nextPos := pos{currPos.r - 1, currPos.c}
				updatedRoutes = processNextPos(nextPos, r.len+1, field, minStepsAtPos, updatedRoutes)
			}

			// right
			if currPos.c < len(field[currPos.r])-1 {
				nextPos := pos{currPos.r, currPos.c + 1}
				updatedRoutes = processNextPos(nextPos, r.len+1, field, minStepsAtPos, updatedRoutes)
			}

			// left
			if currPos.c > 0 {
				nextPos := pos{currPos.r, currPos.c - 1}
				updatedRoutes = processNextPos(nextPos, r.len+1, field, minStepsAtPos, updatedRoutes)
			}

			// down
			if currPos.r < len(field)-1 {
				nextPos := pos{currPos.r + 1, currPos.c}
				updatedRoutes = processNextPos(nextPos, r.len+1, field, minStepsAtPos, updatedRoutes)
			}
		}

		fmt.Printf("%v\n", updatedRoutes)

		if len(updatedRoutes) != 0 {
			routes = updatedRoutes
		} else {
			break
		}
	}

	fmt.Println(minStepsAtPos[pos{70, 70}])
}

func processNextPos(nextPos pos, nextLen int, field [][]int, minStepsAtPos map[pos]int, updatedRoutes []route) []route {
	if field[nextPos.r][nextPos.c] != 1 {
		minSteps, ok := minStepsAtPos[nextPos]
		if !ok || minSteps > nextLen {
			minStepsAtPos[nextPos] = nextLen
			return append(updatedRoutes, route{nextPos, nextLen})
		}
	}

	return updatedRoutes
}
