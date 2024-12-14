package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

func (r *robot) move(mapWidth int, mapHeight int) {
	r.x = (mapWidth + r.x + r.vx) % mapWidth
	r.y = (mapHeight + r.y + r.vy) % mapHeight
}

func main() {
	file, err := os.Open("input.txt")
	//file, err := os.Open("example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var robots []robot = make([]robot, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var x, y, vx, vy int
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)

		robots = append(robots, robot{x, y, vx, vy})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", robots)
	fmt.Println()

	var w int = 101
	var h int = 103
	//var w int = 11
	//var h int = 7

	for range 100 {
		for r, _ := range robots {
			robots[r].move(w, h)
		}
	}

	fmt.Printf("%v\n", robots)

	var halfW int = w / 2
	var halfH int = h / 2
	var q1Count, q2Count, q3Count, q4Count int
	for _, r := range robots {
		if r.x < halfW && r.y < halfH {
			fmt.Printf("1: %v\n", r)
			q1Count++
		} else if r.x < halfW && r.y > halfH {
			q2Count++
			fmt.Printf("2: %v\n", r)
		} else if r.x > halfW && r.y > halfH {
			q3Count++
			fmt.Printf("3: %v\n", r)
		} else if r.x > halfW && r.y < halfH {
			q4Count++
			fmt.Printf("4: %v\n", r)
		}
	}

	fmt.Println(q1Count * q2Count * q3Count * q4Count)

}
