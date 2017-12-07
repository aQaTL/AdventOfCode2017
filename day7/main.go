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

	bottomName := partOne(input)
	fmt.Printf("%s\n", bottomName)

	partTwo(input, bottomName)
}

type prog struct {
	Name   string
	Weight int
	Aboves []*prog
}

func partTwo(input []string, bottomName string) {
	bot := buildTree(&input, bottomName)
	findUnbalanced(bot)
}

func buildTree(input *[]string, name string) *prog {
	p := &prog{Name: name}
	for _, line := range *input {
		progData := strings.Split(line, " ")

		if progData[0] == name {
			num, _ := strconv.Atoi(progData[1][1: len(progData[1])-1])
			p.Weight = num

			if len(progData) > 2 {
				aboves := make([]*prog, 0)
				for i := 3; i < len(progData); i++ {
					a := strings.TrimSpace(strings.TrimSuffix(progData[i], ","))
					aboves = append(aboves, buildTree(input, a))
				}
				p.Aboves = aboves
			} else {
				return p
			}
		}
	}
	return p
}

func findUnbalanced(p *prog) uint64 {
	total := uint64(p.Weight)

	if len(p.Aboves) != 0 {
		weights := make([]uint64, len(p.Aboves))
		for i, a := range p.Aboves {
			weights[i] = findUnbalanced(a)
			total += weights[i]
		}

		for i := 0; i < len(weights)-1; i++ {
			for j := i + 1; j < len(weights); j++ {
				if weights[i] != weights[j] {
					diff := weights[i] - weights[j]

					fmt.Printf("It should be: %d or %d\n",
						uint64(p.Aboves[i].Weight)-diff,
						uint64(p.Aboves[j].Weight)+diff)
				}
			}
		}
	}
	return total
}

func partOne(input []string) string {
	names := make([]string, 0)
	aboves := make(map[string]interface{})

	for _, line := range input {
		data := strings.Split(line, " ")
		names = append(names, data[0])
		if len(data) > 2 {
			for i := 3; i < len(data); i++ {
				a := strings.TrimSpace(strings.TrimSuffix(data[i], ","))
				aboves[a] = *new(interface{})
			}
		}
	}

	for _, name := range names {
		if _, seen := aboves[name]; !seen {
			return name
		}
	}
	return ""
}
