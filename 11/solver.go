package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stone struct {
	engraving int
	nextStone *stone
}

func (s *stone) needSplit() bool {
	engString := strconv.FormatInt(int64(s.engraving), 10)
	return len(engString)%2 == 0
}

func (s *stone) split() {
	currEngraving := strconv.FormatInt(int64(s.engraving), 10)
	newEngraving, _ := strconv.Atoi(currEngraving[:len(currEngraving)/2])
	s.engraving = newEngraving

	newEngraving, _ = strconv.Atoi(currEngraving[len(currEngraving)/2:])
	newStone := stone{newEngraving, s.nextStone}

	s.nextStone = &newStone
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	splitLine := strings.Split(scanner.Text(), " ")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var firstStone *stone
	var currStone *stone
	for _, s := range splitLine {
		i, _ := strconv.Atoi(s)
		newStone := stone{i, nil}
		if currStone != nil {
			currStone.nextStone = &newStone
		}
		currStone = &newStone

		if firstStone == nil {
			firstStone = currStone
		}
	}

	for i := range 75 {
		for s := firstStone; s != nil; {
			if s.engraving == 0 {
				s.engraving = 1
				s = s.nextStone
			} else if s.needSplit() {
				s.split()
				s = s.nextStone.nextStone
			} else {
				s.engraving = s.engraving * 2024
				s = s.nextStone
			}
		}

		fmt.Println(i)
	}

	var res int
	for s := firstStone; s != nil; s = s.nextStone {
		res++
	}

	fmt.Printf("%v\n", res)
}
