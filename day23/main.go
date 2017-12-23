package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	fmt.Println(partOne(input))
}

func partOne(input []string) int {
	registers := make(map[string]int)

	atoi := func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return registers[s]
		}
		return num
	}

	mulCount := 0

	for i := 0; i < len(input); {
		instr := strings.Split(input[i], " ")
		switch instr[0] {
		case "set":
			registers[instr[1]] = atoi(instr[2])
		case "mul":
			registers[instr[1]] *= atoi(instr[2])
			mulCount++
		case "sub":
			registers[instr[1]] -= atoi(instr[2])
		case "jnz":
			if atoi(instr[1]) != 0 {
				i += atoi(instr[2])
				continue
			}
		}
		i++
	}

	return mulCount
}
