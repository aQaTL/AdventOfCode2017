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

	layers := make([]layer, len(input))
	for i, line := range input {
		words := strings.Split(line, ": ")
		depth, _ := strconv.Atoi(words[0])
		size, _ := strconv.Atoi(words[1])
		layers[i] = layer{uint64(depth), uint64(size)}
	}

	fmt.Printf("%d\n%d\n",
		partOne(layers),
		partTwo(layers))
}

func partTwo(layers []layer) uint64 {
loop:
	for delay := uint64(0); ; delay++ {
		for _, l := range layers {
			if isSpottedByScanner(delay, l) {
				continue loop
			}
		}
		return delay
	}
}

type layer struct {
	depth uint64
	size  uint64
}

func isSpottedByScanner(delay uint64, l layer) bool {
	return (l.depth+delay)%((l.size-1)*2) == 0
}

func partOne(layers []layer) uint64 {
	var severity uint64
	for _, l := range layers {
		if isSpottedByScanner(0, l) {
			severity += l.depth * l.size
		}
	}
	return severity
}
