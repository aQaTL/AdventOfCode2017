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

	memory := intSlice(strings.Split(string(bytes), "\t"))

	newMem := partOne(memory)
	partOne(newMem) //PartTwo
}

func partOne(memory []int) []int {
	configurations := make(map[string]interface{})
	cycles := 0

	for {
		memConfig := memoryToString(memory)
		if _, seen := configurations[memConfig]; !seen {
			configurations[memConfig] = *new(interface{})
		} else {
			break
		}

		highestIdx := biggestBankIndex(memory)
		blocks := memory[highestIdx]
		memory[highestIdx] = 0

		for i := (highestIdx + 1) % len(memory); blocks != 0; i = (i + 1) % len(memory) {
			memory[i]++
			blocks--
		}
		cycles++
	}

	fmt.Println(cycles)

	return memory
}

func memoryToString(slice []int) string {
	stringSlice := make([]string, len(slice))
	for i, s := range slice {
		stringSlice[i] = strconv.Itoa(s)
	}
	return strings.Join(stringSlice, " ")
}

func biggestBankIndex(memory []int) int {
	idx, max := 0, -(1 << 31)
	for i, mem := range memory {
		if mem > max {
			max = mem
			idx = i
		}
	}
	return idx
}

func intSlice(slice []string) []int {
	intSlice := make([]int, len(slice))
	for i, s := range slice {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intSlice[i] = num
	}
	return intSlice
}
