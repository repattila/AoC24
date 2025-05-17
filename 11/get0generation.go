package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	lenGeneratedFrom0 := map[int]int{
		0:  1,
		1:  1,
		2:  2,
		3:  4,
		4:  4,
		5:  7,
		6:  14,
		7:  16,
		8:  20,
		9:  39,
		10: 62,
		11: 81,
		12: 110,
		13: 200,
		14: 328,
		15: 418,
		16: 667,
		17: 1059,
		18: 1546,
		19: 2377,
		20: 3572,
		21: 5602,
		22: 8268,
		23: 12343,
		24: 19778,
		25: 29165,
		26: 43726,
		27: 67724,
		28: 102131,
		29: 156451,
		30: 234511,
		31: 357632,
		32: 549949,
		33: 819967,
		34: 1258125,
		35: 1916299,
		36: 2886408,
		37: 4414216,
		38: 6669768,
		39: 10174278,
	}

	var stones []int = make([]int, 0, 80028872400)
	stones = append(stones, 0)

	start := time.Now()

	fmt.Printf("map[string]int{\n")

	var stepNum int = 40
	var lenFrom0s int = 0

	for i := range stepNum {
		for j := range len(stones) {
			var s int = stones[j]
			if s == -1 {
				// skip
			} else if s == 0 {
				generatedLen, ok := lenGeneratedFrom0[stepNum-1-i]

				if ok {
					lenFrom0s += generatedLen
					stones[j] = -1
				} else {
					stones[j] = 1
				}
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

		lenStones := 0
		for _, s := range stones {
			if s != -1 {
				lenStones += 1
			}
		}

		fmt.Printf("%v: %v,\n", i, lenStones+lenFrom0s)
	}

	fmt.Printf("}\n")

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
