package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	partOne, partTwo := bothParts(&bytes)
	fmt.Printf("%d\n%d\n", partOne, partTwo)
}

func bothParts(input *[]byte) (partOne, partTwo int) {
	score := 0
	scoreInc := 0
	garbage := false
	garbageCount := 0

	for i := 0; i < len(*input); i++ {
		switch (*input)[i] {
		case '{':
			if !garbage {
				scoreInc++
			} else {
				garbageCount++
			}
		case '}':
			if !garbage {
				score += scoreInc
				scoreInc--
			} else {
				garbageCount++
			}
		case '!':
			i++
		case '<':
			if garbage {
				garbageCount++
			}
			garbage = true
		case '>':
			garbage = false
		default:
			if garbage {
				garbageCount++
			}
		}
	}
	return score, garbageCount
}
