package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(bytes), "\n")

	registers := make(map[string]int)

	maxEver := 0

	for _, line := range input {
		instr := strings.Split(line, " ")

		regName := instr[0]
		mode := instr[1]
		val, _ := strconv.Atoi(instr[2])

		regInCondition := instr[4]
		conditionVal, _ := strconv.Atoi(instr[6])

		if mode == "dec" {
			val *= -1
		}

		switch instr[5] {
		case "==":
			if registers[regInCondition] == conditionVal {
				registers[regName] += val
			}
		case "!=":
			if registers[regInCondition] != conditionVal {
				registers[regName] += val
			}
		case ">":
			if registers[regInCondition] > conditionVal {
				registers[regName] += val
			}
		case ">=":
			if registers[regInCondition] >= conditionVal {
				registers[regName] += val
			}
		case "<":
			if registers[regInCondition] < conditionVal {
				registers[regName] += val
			}
		case "<=":
			if registers[regInCondition] <= conditionVal {
				registers[regName] += val
			}
		default:
			fmt.Println("unimplemented instruction: ", instr[5])
		}

		if registers[regName] > maxEver {
			maxEver = registers[regName]
		}
	}

	max := -(1 << 31)
	for _, v := range registers {
		if v > max {
			max = v
		}
	}

	fmt.Println(max) //PartOne
	fmt.Println(maxEver) //PartTwo
}
