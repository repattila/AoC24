package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var secretNumber int
		fmt.Sscanf(scanner.Text(), "%d", &secretNumber)

		fmt.Printf("%d:", secretNumber)

		for range 2000 {
			secretNumber = calcNextSecret(secretNumber)
		}

		fmt.Printf("%d\n", secretNumber)

		res += secretNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func calcNextSecret(secret int) int {
	secret = (secret ^ (secret * 64)) % 16777216
	secret = (secret ^ (secret / 32)) % 16777216
	secret = (secret ^ (secret * 2048)) % 16777216

	return secret
}
