package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var initValuesTemplate = `Begin in state %s
Perform a diagnostic checksum after %d steps. `

var stateTemplate = `In state %s
  If the current value is 0:
    - Write the value %d.
    - Move one slot to the %s
    - Continue with state %s
  If the current value is 1:
    - Write the value %d.
    - Move one slot to the %s
    - Continue with state %s `

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n\n")

	curState, steps := "", 0
	_, err = fmt.Sscanf(
		input[0],
		initValuesTemplate,
		&curState,
		&steps)
	if err != nil {
		panic(err)
	}

	curState = curState[:len(curState)-1]

	states := make(map[string]stateInstr)

	for i := 1; i < len(input); i++ {
		stateName := ""

		if0, if1 := instr{}, instr{}
		if0MoveBy, if1MoveBy := "", ""

		_, err := fmt.Sscanf(
			input[i],
			stateTemplate,
			&stateName,
			&if0.write,
			&if0MoveBy,
			&if0.continueWith,
			&if1.write,
			&if1MoveBy,
			&if1.continueWith)
		if err != nil {
			panic(err)
		}

		stateName = stateName[:len(stateName)-1]
		if0.continueWith = if0.continueWith[:len(if0.continueWith)-1]
		if1.continueWith = if1.continueWith[:len(if1.continueWith)-1]

		if if0MoveBy == "right." {
			if0.moveBy = 1
		} else {
			if0.moveBy = -1
		}
		if if1MoveBy == "right." {
			if1.moveBy = 1
		} else {
			if1.moveBy = -1
		}

		states[stateName] = stateInstr{if0, if1}
	}

	tape := make(map[int]int, 0)
	pos := 0

	for i := 0; i < steps; i++ {
		instr := states[curState]
		if tape[pos] == 1 {
			tape[pos] = instr.if1.write
			pos += instr.if1.moveBy
			curState = instr.if1.continueWith
		} else {
			tape[pos] = instr.if0.write
			pos += instr.if0.moveBy
			curState = instr.if0.continueWith
		}
	}

	count := 0
	for _, v := range tape {
		if v == 1 {
			count++
		}
	}
	fmt.Println(count)
}

type stateInstr struct {
	if0, if1 instr
}

type instr struct {
	write        int
	moveBy       int
	continueWith string
}
