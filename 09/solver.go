package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fileSystem []int = make([]int, 0)
	var fileId int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()

	for i, r := range scanner.Text() {
		blockCount, _ := strconv.Atoi(string(r))

		if i%2 == 0 {
			for range blockCount {
				fileSystem = append(fileSystem, fileId)
			}

			fileId++
		} else {
			for range blockCount {
				fileSystem = append(fileSystem, -1)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", fileSystem)

	var moveFrom int = len(fileSystem) - 1
	for i := 0; i < len(fileSystem); i++ {
		if fileSystem[i] == -1 {
			if moveFrom < i {
				moveFrom = len(fileSystem) - 1
			}

			for fileSystem[moveFrom] == -1 && moveFrom > 0 {
				moveFrom--
			}
			if moveFrom > i {
				fileSystem[i] = fileSystem[moveFrom]
				fileSystem[moveFrom] = -1
			}
		}
	}

	fmt.Printf("%v\n", fileSystem)

	var res int
	for i, id := range fileSystem {
		if id != -1 {
			res += id * i
		}
	}

	fmt.Println(res)
}
