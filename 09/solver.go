package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type blocks struct {
	free  bool
	id    int
	start int
	len   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var files []blocks = make([]blocks, 0)
	var freeSpace []blocks = make([]blocks, 0)
	var fileId int
	var fileSystemSize int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()

	var currPos int
	for i, r := range scanner.Text() {
		blockCount, _ := strconv.Atoi(string(r))

		if i%2 == 0 {
			files = append(files, blocks{false, fileId, currPos, blockCount})

			fileId++
			fileSystemSize += blockCount
		} else {
			if blockCount > 0 {
				freeSpace = append(freeSpace, blocks{true, -1, currPos, blockCount})

				fileSystemSize += blockCount
			}
		}

		currPos += blockCount
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", files)
	fmt.Printf("%v\n", freeSpace)
	fmt.Printf("%v\n", fileSystemSize)

	for i := len(files) - 1; i >= 0; i-- {
		currFile := files[i]
		for j := 0; j < len(freeSpace); j++ {
			currFree := freeSpace[j]
			if currFree.start < currFile.start {
				if currFree.free {
					if currFree.len >= currFile.len {
						// Move file to beginning of free space
						files[i].start = currFree.start

						if currFree.len > currFile.len {
							// Decrease free space size
							freeSpace[j].start = currFree.start + currFile.len
							freeSpace[j].len = currFree.len - currFile.len
						} else if currFree.len == currFile.len {
							// Free space is fully used
							freeSpace[j].free = false
						}

						// Add free space where the file was originally
						freeSpace = append(freeSpace, blocks{true, -1, currFile.start, currFile.len})
						break
					}
				}
			} else {
				break
			}
		}
	}

	fmt.Printf("%v\n", files)
	fmt.Printf("%v\n", freeSpace)

	var res int
posLoop:
	for pos := 0; pos < fileSystemSize; {
		fmt.Println(pos)

		for _, f := range files {
			if f.start == pos {
				for i := 0; i < f.len; i++ {
					res += f.id * (pos + i)
				}

				pos += f.len
				continue posLoop
			}
		}

		for _, f := range freeSpace {
			if f.free && f.start == pos {
				pos += f.len
				continue posLoop
			}
		}
	}

	fmt.Println(res)
}
