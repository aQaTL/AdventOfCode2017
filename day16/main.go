package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type moveType byte

const (
	spin moveType = iota
	swapIdx
	swapVal
)

type move struct {
	moveType moveType
	a, b     int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	moves := parseInput(data)

	length := 16
	initialState := genAlphabet(length)
	initialStateStr := string(initialState)

	part2 := make([]rune, length)
	copy(part2, initialState)

	partOne := dance(moves, initialState)
	fmt.Println("Part 1:", string(partOne))

	steps := 1
	for ; steps < 1000000000; steps++ {
		part2 = dance(moves, part2)
		if string(part2) == initialStateStr {
			break
		}
	}
	for i := 0; i < 1000000000%steps; i++ {
		part2 = dance(moves, part2)
	}

	fmt.Println("Part 2:", string(part2))
}

func dance(moves []move, progs []rune) []rune {
	for _, move := range moves {
		switch move.moveType {
		case spin:
			prog := progs[:len(progs)-move.a]
			progs = progs[len(progs)-move.a:]
			progs = append(progs, prog...)
		case swapIdx:
			progs[move.a], progs[move.b] = progs[move.b], progs[move.a]
		case swapVal:
			a, b := 0, 0
			for i, r := range progs {
				if int(r) == move.a {
					a = i
				}
				if int(r) == move.b {
					b = i
				}
			}
			progs[a], progs[b] = progs[b], progs[a]
		}
	}
	return progs
}

func parseInput(data []byte) []move {
	moves := make([]move, 0)
	for _, s := range strings.Split(string(data), ",") {
		switch s[0] {
		case 's':
			x, _ := strconv.Atoi(s[1:])
			moves = append(moves, move{spin, x, 0})
		case 'x':
			s := strings.Split(s, "/")
			a, _ := strconv.Atoi(s[0][1:])
			b, _ := strconv.Atoi(s[1])
			moves = append(moves, move{swapIdx, a, b})
		case 'p':
			a, b := s[1], s[3]
			moves = append(moves, move{swapVal, int(a), int(b)})
		default:
			fmt.Println("Unknown move:", s[0])
		}
	}
	return moves
}

func genAlphabet(upTo int) []rune {
	initialState := make([]rune, upTo)
	for i := range initialState {
		initialState[i] = rune(i + 97)
	}
	return initialState
}
