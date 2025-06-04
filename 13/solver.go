package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		m.px = x + 10000000000000
		m.py = y + 10000000000000

		machines = append(machines, m)

		// Read empty line
		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", machines)

	var res int
machineLoop:
	for n, m := range machines {
		fmt.Printf("Machine: %v\n", n)

		var aRate float64 = float64(m.ax) / float64(m.ay)
		var bRate float64 = float64(m.bx) / float64(m.by)
		var pRate float64 = float64(m.px) / float64(m.py)

		if almostEqual(aRate, bRate) && !almostEqual(aRate, pRate) {
			fmt.Printf("Unsolvable\n")
			continue machineLoop
		} else {
			// Solving the following equation system:
			// i * m.ax + j * m.bx = m.px
			// i * m.ay + j * m.by = m.py

			var d int = m.ax*m.by - m.bx*m.ay
			fmt.Printf("d: %v\n", d)
			if d != 0 {
				var di int = m.px*m.by - m.bx*m.py
				var dj int = m.ax*m.py - m.px*m.ay

				// We only care about solutions that give whole numbers for button presses
				if di%d == 0 && dj%d == 0 {
					var i int = di / d
					var j int = dj / d

					// We only care about solutions where the button presses are positive
					if i >= 0 && j >= 0 {
						fmt.Printf("Press A: %v, press B: %v\n\n", i, j)

						res += i*3 + j
					} else {
						fmt.Printf("Solution has negative value\n\n")
					}
				} else {
					fmt.Printf("Solution has fractions\n\n")
				}
			} else {
				var multB int = m.px / m.bx

				if m.px%m.bx == 0 {
					if m.by*multB == m.py {
						res += multB
						continue machineLoop
					}
				}

				for i := multB; i >= 0; i-- {
					if (m.px-m.bx*i)%m.ax == 0 {
						var multA int = (m.px - m.bx*i) / m.ax
						if m.by*i+m.ay*multA == m.py {
							fmt.Printf("Press A: %v, press B: %v\n\n", multA, i)

							res += i + 3*multA
							continue machineLoop
						}
					}
				}
			}
		}
	}

	fmt.Printf("%v\n", res)
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
