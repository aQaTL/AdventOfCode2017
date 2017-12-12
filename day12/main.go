package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(
		strings.
			NewReplacer(",", "", "<-> ", "").
			Replace(string(bytes)),
		"\n",
	)

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input []string) int {
	progs := make(map[string]interface{})
	buildGroup("0", input, progs)
	return len(progs)
}

func partTwo(input []string) int {
	progs := make(map[string]interface{})
	groups := 0
	
	for _, line := range input {
		pipes := strings.Split(line, " ")
		if _, seen := progs[pipes[0]]; !seen {
			buildGroup(pipes[0], input, progs)
			groups++
		}
	}
	
	return groups
}

func buildGroup(root string, input []string, progs map[string]interface{}) {
	for _, line := range input {
		program := strings.Split(line, " ")
		if program[0] == root {
			if _, seen := progs[root]; seen {
				return
			}
			progs[root] = *new(interface{})
			
			for _, pipe := range program[1:] {
				buildGroup(pipe, input, progs)
			}
			
			return
		}
	}
}
