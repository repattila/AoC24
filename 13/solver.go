package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type machine struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var machines []machine = make([]machine, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var m machine = machine{}
		var x, y int

		line := scanner.Text()
		line = line[10:]
		fmt.Sscanf(line, "X%d, Y%d", &x, &y)
		m.ax = x
		m.ay = y

		scanner.Scan()
		line = scanner.Text()
		line = line[10:]
		fmt.Sscanf(line, "X%d, Y%d", &x, &y)
		m.bx = x
		m.by = y

		scanner.Scan()
		line = scanner.Text()
		line = line[7:]
		fmt.Sscanf(line, "X=%d, Y=%d", &x, &y)
		m.px = x
		m.py = y

		machines = append(machines, m)

		// Read empty line
		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", machines)

	var res int
	for _, m := range machines {
		var minCost int = 500
		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				if i*m.ax+j*m.bx == m.px && i*m.ay+j*m.by == m.py {
					minCost = min(minCost, i*3+j*1)
				}
			}
		}
		if minCost != 500 {
			res += minCost
		}
	}

	fmt.Printf("%v\n", res)
}
