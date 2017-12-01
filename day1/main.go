package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	input := readInput("input.txt")

	partOne(input)
	partTwo(input)
}

func partOne(input []int) {
	matches := make(map[int]int)
	for i, v := range input {
		next := input[(i + 1)%len(input)]
		if v == next {
			matches[v]++
		}
	}

	sum := 0
	for k, v := range matches {
		sum += k * v
	}

	fmt.Println(sum)
}

func partTwo(input []int) {
	matches := make(map[int]int)
	for i, v := range input {
		next := input[(i + len(input)/2)%len(input)]
		if v == next {
			matches[v]++
		}
	}

	sum := 0
	for k, v := range matches {
		sum += k * v
	}

	fmt.Println(sum)
}

func readInput(filename string) []int {
	bytes, err := ioutil.ReadFile(filename) //Remember to delete LF char
	if err != nil {
		panic(err)
	}

	input := make([]int, len(bytes))
	for i, b := range bytes {
		input[i] = int(b - 48)
	}

	return input
}
