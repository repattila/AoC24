package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gate struct {
	in1       string
	in2       string
	out       string
	op        string
	evaluated bool
}

func (g *gate) evaluate(setWires map[string]bool, setOutputs map[string]bool) {
	in1Val, ok1 := setWires[g.in1]
	if ok1 {
		in2Val, ok2 := setWires[g.in2]
		if ok2 {
			switch g.op {
			case "AND":
				if g.out[0] == 'z' {
					setOutputs[g.out] = in1Val && in2Val
				} else {
					setWires[g.out] = in1Val && in2Val
				}
			case "OR":
				if g.out[0] == 'z' {
					setOutputs[g.out] = in1Val || in2Val
				} else {
					setWires[g.out] = in1Val || in2Val
				}
			case "XOR":
				if g.out[0] == 'z' {
					setOutputs[g.out] = in1Val != in2Val
				} else {
					setWires[g.out] = in1Val != in2Val
				}
			}

			g.evaluated = true
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var setWires map[string]bool = make(map[string]bool)
	var setOutputs map[string]bool = make(map[string]bool)
	var outputCount int
	var gates []gate = make([]gate, 0)
	var readWireVals bool = true

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line string = scanner.Text()
		var splitLine []string

		if len(line) == 0 {
			readWireVals = false
			continue
		} else {
			splitLine = strings.Split(line, " ")
		}

		if readWireVals {
			setWires[splitLine[0][:len(splitLine[0])-1]] = splitLine[1] == "1"
		} else {
			gates = append(gates, gate{splitLine[0], splitLine[2], splitLine[4], splitLine[1], false})
			if splitLine[4][0] == 'z' {
				outputCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", setWires)
	fmt.Printf("%v\n", gates)
	fmt.Printf("%v\n", outputCount)

	for len(setOutputs) < outputCount {
		for _, gate := range gates {
			if !gate.evaluated {
				gate.evaluate(setWires, setOutputs)
			}
		}
	}

	fmt.Printf("%v\n", setOutputs)

	var res string
	for i := outputCount - 1; i >= 0; i-- {
		var is string
		if i < 10 {
			is = "0" + strconv.Itoa(i)
		} else {
			is = strconv.Itoa(i)
		}

		val, _ := setOutputs["z"+is]
		if val {
			res += "1"
		} else {
			res += "0"
		}
	}

	fmt.Printf("%s\n", res)

	i, err := strconv.ParseInt(res, 2, 64)
	fmt.Printf("%d\n", i)
}
