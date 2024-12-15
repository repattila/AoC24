package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var warehouse [][]int = make([][]int, 0)
	var moves []int = make([]int, 0)
	var robotR, robotC, l int
	var readWarehouse bool = true
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if readWarehouse {
			if len(line) == 0 {
				readWarehouse = false
				continue
			}

			var row []int = make([]int, 0)
			for c, val := range line {
				if val == '#' {
					row = append(row, 0)
				} else if val == '.' {
					row = append(row, 1)
				} else if val == 'O' {
					row = append(row, 2)
				} else if val == '@' {
					row = append(row, 1)
					robotR = l
					robotC = c
				}
			}
			warehouse = append(warehouse, row)
			l++
		} else {
			for _, val := range line {
				if val == '<' {
					moves = append(moves, 3)
				} else if val == 'v' {
					moves = append(moves, 2)
				} else if val == '^' {
					moves = append(moves, 0)
				} else if val == '>' {
					moves = append(moves, 1)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", warehouse)
	fmt.Printf("%v\n", moves)
	fmt.Printf("%d %d\n", robotR, robotC)

	for _, m := range moves {
		var nextRobotR, nextRobotC int = getNextPos(robotR, robotC, m)
		next := warehouse[nextRobotR][nextRobotC]
		if next == 0 {
			continue
		} else if next == 1 {
			robotR = nextRobotR
			robotC = nextRobotC
		} else {
			if canBoxMove(nextRobotR, nextRobotC, m, warehouse) {
				robotR = nextRobotR
				robotC = nextRobotC
			}
		}
	}

	fmt.Printf("%v\n", warehouse)

	var res int
	for r, row := range warehouse {
		for c, val := range row {
			if val == 2 {
				res += 100*r + c
			}
		}
	}

	fmt.Println(res)
}

func getNextPos(r int, c int, m int) (int, int) {
	var nextR, nextC int
	switch m {
	case 0:
		nextR = r - 1
		nextC = c
	case 1:
		nextR = r
		nextC = c + 1
	case 2:
		nextR = r + 1
		nextC = c
	case 3:
		nextR = r
		nextC = c - 1
	}
	return nextR, nextC
}

func canBoxMove(r int, c int, dir int, warehouse [][]int) bool {
	var nextR, nextC int = getNextPos(r, c, dir)

	next := warehouse[nextR][nextC]
	if next == 0 {
		return false
	} else if next == 1 {
		warehouse[nextR][nextC] = 2
		warehouse[r][c] = 1
		return true
	} else {
		if canBoxMove(nextR, nextC, dir, warehouse) {
			warehouse[nextR][nextC] = 2
			warehouse[r][c] = 1
			return true
		} else {
			return false
		}
	}
}
