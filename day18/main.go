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

	fmt.Println("Part 1:", partOne(input))
	partTwo(input)
}

func partTwo(input []string) {
	sent, rcv := make(chan int, 1000000), make(chan int, 1000000) //I hope it will be enough
	process := func(send chan<- int, rcv <-chan int, progID int) {
		registers := make(map[string]int)
		registers["p"] = progID

		valuesSent := 0

		atoi := func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				return registers[s]
			}
			return num
		}
		for i := 0; i < len(input); {
			instr := strings.Split(input[i], " ")
			switch instr[0] {
			case "set":
				registers[instr[1]] = atoi(instr[2])
			case "mul":
				registers[instr[1]] *= atoi(instr[2])
			case "add":
				registers[instr[1]] += atoi(instr[2])
			case "mod":
				registers[instr[1]] %= atoi(instr[2])
			case "rcv":
				registers[instr[1]] = <-rcv
			case "jgz":
				if atoi(instr[1]) > 0 {
					i += atoi(instr[2])
					continue
				}
			case "snd":
				send <- atoi(instr[1])
				valuesSent++
				if progID == 1 {
					fmt.Printf("\rPart 2: %d", valuesSent)
				}
			}
			i++
		}
	}
	go process(sent, rcv, 0)
	process(rcv, sent, 1)
}

func partOne(input []string) int {
	registers := make(map[string]int)
	lastPlayed := 0

	atoi := func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return registers[s]
		}
		return num
	}

	for i := 0; i < len(input); {
		instr := strings.Split(input[i], " ")
		switch instr[0] {
		case "set":
			registers[instr[1]] = atoi(instr[2])
		case "mul":
			registers[instr[1]] *= atoi(instr[2])
		case "add":
			registers[instr[1]] += atoi(instr[2])
		case "mod":
			registers[instr[1]] %= atoi(instr[2])
		case "rcv":
			if val := registers[instr[1]]; val != 0 {
				return lastPlayed
			}
		case "jgz":
			if atoi(instr[1]) > 0 {
				i += atoi(instr[2])
				continue
			}
		case "snd":
			lastPlayed = registers[instr[1]]
		}
		i++
	}
	return lastPlayed
}
