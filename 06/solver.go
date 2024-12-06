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

	var lab [][]bool = make([][]bool, 0)
	var visited map[pos]int = make(map[pos]int)
	var guardPos pos
	var guardDir int
	var r int = 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line string
		fmt.Sscanf(scanner.Text(), "%s", &line)

		var obstacleRow []bool = make([]bool, 0)
		for c, ch := range line {
			if ch == '.' {
				obstacleRow = append(obstacleRow, false)
			} else if ch == '#' {
				obstacleRow = append(obstacleRow, true)
			} else {
				obstacleRow = append(obstacleRow, false)

				guardPos = pos{r, c}
				switch ch {
				case '^':
					guardDir = 0
				case '>':
					guardDir = 1
				case 'v':
					guardDir = 2
				case '<':
					guardDir = 3
				}
			}
		}
		lab = append(lab, obstacleRow)

		r++
	}

	fmt.Println(followGuardPath(guardPos, guardDir, visited, lab))

	var loopCount int = 0
	i := 0
	for k := range visited {
		if i != 0 {
			var visitedLC map[pos]int = make(map[pos]int)
			lab[k.row][k.col] = true
			if hasLoop(guardPos, guardDir, visitedLC, lab) {
				loopCount++
			}
			lab[k.row][k.col] = false
		}

		i++
	}

	fmt.Println(loopCount)
}

func getNextPos(p pos, dir int) pos {
	var nextPos pos
	switch dir {
	case 0:
		nextPos = pos{p.row - 1, p.col}
	case 1:
		nextPos = pos{p.row, p.col + 1}
	case 2:
		nextPos = pos{p.row + 1, p.col}
	case 3:
		nextPos = pos{p.row, p.col - 1}
	}
	return nextPos
}

func followGuardPath(p pos, dir int, visited map[pos]int, lab [][]bool) int {
	var nextPos pos = getNextPos(p, dir)

	if nextPos.row < 0 || nextPos.col < 0 || nextPos.row >= len(lab) || nextPos.col >= len(lab[0]) {
		return len(visited) + 1
	} else if !lab[nextPos.row][nextPos.col] {
		visited[p] = 1
		return followGuardPath(nextPos, dir, visited, lab)
	} else {
		return followGuardPath(p, (dir+1)%4, visited, lab)
	}
}

func hasLoop(p pos, dir int, visited map[pos]int, lab [][]bool) bool {
	var nextPos pos = getNextPos(p, dir)

	if nextPos.row < 0 || nextPos.col < 0 || nextPos.row >= len(lab) || nextPos.col >= len(lab[0]) {
		return false
	} else if !lab[nextPos.row][nextPos.col] {
		d, ok := visited[p]
		if ok && d == dir {
			return true
		} else {
			visited[p] = dir
			return hasLoop(nextPos, dir, visited, lab)
		}
	} else {
		return hasLoop(p, (dir+1)%4, visited, lab)
	}
}
