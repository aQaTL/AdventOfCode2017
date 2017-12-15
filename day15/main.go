package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var factorA, factorB = 16807, 48271
var mask = 0xffff

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.Replace(string(data), "\n", " ", -1), " ")
	a, _ := strconv.Atoi(input[4])
	b, _ := strconv.Atoi(input[9])

	matches := partOne(a, b)
	fmt.Println(matches)

	matches = partTwo(a, b)
	fmt.Println(matches)
}

func partTwo(a, b int) int {
	matches := 0
	pairs := 0
	aQueue, bQueue := list.New(), list.New()
	for pairs < 5000000 {
		a = gen(a, factorA)
		b = gen(b, factorB)
		if a%4 == 0 {
			aQueue.PushBack(a)
		}
		if b%8 == 0 {
			bQueue.PushBack(b)
		}
		if aQueue.Len() > 0 && bQueue.Len() > 0 {
			elA := aQueue.Front()
			elB := bQueue.Front()
			if (elA.Value.(int)&mask)^(elB.Value.(int)&mask) == 0 {
				matches++
			}
			aQueue.Remove(elA)
			bQueue.Remove(elB)
			pairs++
		}
	}
	return matches
}

func partOne(a, b int) int {
	matches := 0
	for i := 0; i < 40000000; i++ {
		a = gen(a, factorA)
		b = gen(b, factorB)
		if (a&mask)^(b&mask) == 0 {
			matches++
		}
	}
	return matches
}

func gen(prev, factor int) int {
	return prev * factor % 2147483647
}
