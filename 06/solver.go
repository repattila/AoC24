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
	var visited map[pos]bool = make(map[pos]bool)
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

	fmt.Println(lab)
	fmt.Println(guardPos)
	fmt.Println(followGuardPath(guardPos, guardDir, visited, lab))
}

func followGuardPath(p pos, dir int, visited map[pos]bool, lab [][]bool) int {
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

	if nextPos.row < 0 || nextPos.col < 0 || nextPos.row >= len(lab) || nextPos.col >= len(lab[0]) {
		return len(visited) + 1
	} else if !lab[nextPos.row][nextPos.col] {
		visited[p] = true
		return followGuardPath(nextPos, dir, visited, lab)
	} else {
		return followGuardPath(p, (dir+1)%4, visited, lab)
	}

}
