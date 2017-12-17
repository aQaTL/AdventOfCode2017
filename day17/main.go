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
	steps, _ := strconv.Atoi(strings.TrimSpace(string(data)))

	buf := make([]int, 1)
	idx := 0
	for i := 0; i < 2017; i++ {
		idx = (idx + steps) % len(buf)
		insert(&buf, idx, i+1)
		idx = (idx + 1) % len(buf)
	}
	fmt.Println("Part 1:", buf[idx])

	spins := 50000000
	buf = make([]int, spins)
	idx = 0
	for i := 0; i < spins; i++ {
		idx = (idx + steps) % (i + 1)
		buf[idx] = i + 1
		idx = (idx + 1) % (i + 2)
	}
	fmt.Println("Part 2:", buf[0])
}

func insert(buf *[]int, idx, val int) {
	*buf = append(*buf, 0)
	copy((*buf)[idx+1:], (*buf)[idx:])
	(*buf)[idx] = val
}
