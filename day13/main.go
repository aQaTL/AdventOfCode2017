package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	severity := partOne(input)
	fmt.Println(severity)

	delay := partTwo(input)
	fmt.Println(delay)
}

func partTwo(input []string) int {
	gotCaught := true
	delay := 0
	for gotCaught {
		gotCaught = false
		for _, line := range input {
			depth, layerRange := 0, 0
			fmt.Sscanf(line, "%d: %d", &depth, &layerRange)

			pos := calcScannerPos(depth+delay, layerRange)
			if pos == 0 {
				gotCaught = true
			}
		}
		if !gotCaught {
			break
		}
		delay++
	}
	return delay
}

func partOne(input []string) int {
	severity := 0
	for _, line := range input {
		depth, layerRange := 0, 0
		fmt.Sscanf(line, "%d: %d", &depth, &layerRange)

		if calcScannerPos(depth, layerRange) == 0 {
			severity += depth * layerRange
		}
	}
	return severity
}

func calcScannerPos(depth, layerRange int) int {
	forward := 1
	pos := 0
	for i := 0; i < depth; i++ {
		pos += forward
		if pos == layerRange-1 || pos == 0 {
			forward *= -1
		}
	}
	return pos
}
