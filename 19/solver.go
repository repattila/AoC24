package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var targetPatterns []string = make([]string, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	var elemPatterns []string = strings.Split(scanner.Text(), ", ")

	scanner.Scan()

	for scanner.Scan() {
		var targetPattern string
		fmt.Sscanf(scanner.Text(), "%s", &targetPattern)
		targetPatterns = append(targetPatterns, targetPattern)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", elemPatterns)
	fmt.Printf("%v\n", targetPatterns)

	var completeMatchCount int
	for _, tp := range targetPatterns {
		fmt.Println(tp)

		var toMatchStart int

		// List of indexes of the matched element patterns
		var match []int = make([]int, 0)

		for i := 0; true; {
			ep := elemPatterns[i]
			if len(tp)-toMatchStart >= len(ep) && tp[toMatchStart:toMatchStart+len(ep)] == ep {
				match = append(match, i)
				toMatchStart += len(ep)

				if toMatchStart == len(tp) {
					fmt.Println(match)
					completeMatchCount++
					break
				} else {
					i = 0
				}
			} else {
				if i < len(elemPatterns)-1 {
					i++
				} else {
					// Backtrack if possible
					var ok bool
					ok, match = backtrack(match, elemPatterns, &i, &toMatchStart)
					if !ok {
						break
					}
				}
			}
		}
	}

	fmt.Println(completeMatchCount)
}

func backtrack(match []int, elemPatterns []string, elemPatternIndex *int, toMatchStart *int) (bool, []int) {
	if len(match) > 0 {
		// Undo last match

		i := match[len(match)-1]
		*toMatchStart = *toMatchStart - len(elemPatterns[i])
		newMatch := match[:len(match)-1]

		if i < len(elemPatterns)-1 {
			// Set next elem pattern

			i++
			*elemPatternIndex = i
			return true, newMatch
		} else {
			// Backtrack again
			return backtrack(newMatch, elemPatterns, elemPatternIndex, toMatchStart)
		}
	} else {
		return false, nil
	}
}
