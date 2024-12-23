package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var connections []string = make([]string, 0)
	var connMap map[string][]string = make(map[string][]string)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var conn, from, to string
		fmt.Sscanf(scanner.Text(), "%s", &conn)
		splitConn := strings.Split(conn, "-")
		from = splitConn[0]
		to = splitConn[1]

		connections = append(connections, conn)

		targets, ok := connMap[from]
		if ok {
			connMap[from] = append(targets, to)
		} else {
			connMap[from] = make([]string, 0)
		}

		targets, ok = connMap[to]
		if ok {
			connMap[to] = append(targets, from)
		} else {
			connMap[to] = make([]string, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", connections)
	fmt.Printf("%v\n", connMap)

	var groups map[string]bool = make(map[string]bool)
	for _, conn := range connections {
		var from, to string
		splitConn := strings.Split(conn, "-")
		from = splitConn[0]
		to = splitConn[1]

		fromConns, _ := connMap[from]
		toConns, _ := connMap[to]
		intersection := SimpleIntersect(fromConns, toConns)

		for _, m := range intersection {
			groupSlice := []string{from, to, m}
			sort.Strings(groupSlice)
			group := groupSlice[0] + groupSlice[1] + groupSlice[2]
			groups[group] = true
		}
	}

	fmt.Printf("%v\n", groups)

	var filteredGroups []string = make([]string, 0)
	for g, _ := range groups {
		if g[0] == 't' || g[2] == 't' || g[4] == 't' {
			filteredGroups = append(filteredGroups, g)
		}
	}

	fmt.Printf("%v\n", filteredGroups)
	fmt.Printf("%d\n", len(filteredGroups))
}

func SimpleIntersect[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if contains(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func contains[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
