package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type cube struct {
	x, y, z float64
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	start, p := cube{}, cube{}
	maxDistanceEver := 0.0

	for _, step := range strings.Split(string(bytes), ",") {
		switch step {
		case "n":
			p.y++
			p.z--
		case "s":
			p.y--
			p.z++
		case "nw":
			p.y++
			p.x--
		case "sw":
			p.x--
			p.z++
		case "ne":
			p.x++
			p.z--
		case "se":
			p.x++
			p.y--
		default:
			fmt.Println("unknown direction:", step)
		}
		if dst := distance(start, p); dst > maxDistanceEver {
			maxDistanceEver = dst
		}
	}

	fmt.Println("Part one:", distance(start, p))
	fmt.Println("Part two:", maxDistanceEver)
}

func distance(a, b cube) float64 {
	return (math.Abs(a.x - b.x) + math.Abs(a.y - b.y) + math.Abs(a.z - b.z)) / 2
}
