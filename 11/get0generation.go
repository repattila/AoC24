package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var stones []int = make([]int, 0, 80028872400)
	stones = append(stones, 4)

	start := time.Now()

	fmt.Printf("map[int]int{\n")

	for i := range 50 {
		for j := range len(stones) {
			var s int = stones[j]
			if s == 0 {
				stones[j] = 1
			} else {
				engString := strconv.FormatInt(int64(s), 10)
				if len(engString)%2 == 0 {
					newEngraving, _ := strconv.Atoi(engString[:len(engString)/2])
					stones = append(stones, newEngraving)

					newEngraving, _ = strconv.Atoi(engString[len(engString)/2:])
					stones[j] = newEngraving
				} else {
					stones[j] = s * 2024
				}
			}
		}

		fmt.Printf("%v: %v,\n", i, len(stones))
	}

	fmt.Printf("}\n")

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
