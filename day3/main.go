package main

import (
	"fmt"
	"math"
)

func main() {
	input := 312051

	fmt.Printf("%d\n%d\n", partOne(input), partTwo(input))
}

func partTwo(input int) int {
	p := point{0, 0,}
	borders := rectangle{}

	memory := make(map[point]int)
	memory[p] = 1

	currVal := 0

	found := func() (found bool) {
		surr := surround(p)
		currVal =
			memory[surr.a] +
				memory[surr.b] +
				memory[surr.c] +
				memory[surr.d] +
				memory[point{p.x, p.y + 1}] +
				memory[point{p.x, p.y - 1}] +
				memory[point{p.x - 1, p.y}] +
				memory[point{p.x + 1, p.y}]

		memory[p] = currVal
		return currVal > input
	}

	for {
		p.x++
		borders = nextRect(borders)

		if found() {
			return currVal
		}

		for p != borders.d {
			p.y++
			if found() {
				return currVal
			}
		}

		for p != borders.a {
			p.x--
			if found() {
				return currVal
			}
		}

		for p != borders.b {
			p.y--
			if found() {
				return currVal
			}
		}

		for p != borders.c {
			p.x++
			if found() {
				return currVal
			}
		}
	}
}

func surround(p point) rectangle {
	return rectangle{
		a: point{p.x - 1, p.y + 1},
		b: point{p.x - 1, p.y - 1},
		c: point{p.x + 1, p.y - 1},
		d: point{p.x + 1, p.y + 1},
	}
}

func partOne(input int) int {
	p := bruteForcePoint(input)
	distance := distance(point{0, 0}, p)
	return distance
}

type point struct {
	x, y int
}

type rectangle struct {
	a, b, c, d point
}

func bruteForcePoint(pointVal int) point {
	p := point{0, 0,}
	currVal := 1
	rect := rectangle{}

	for {
		p.x++
		currVal++
		rect = nextRect(rect)

		for p != rect.d {
			p.y++
			currVal++
			if currVal == pointVal {
				return p
			}
		}

		for p != rect.a {
			p.x--
			currVal++
			if currVal == pointVal {
				return p
			}
		}

		for p != rect.b {
			p.y--
			currVal++
			if currVal == pointVal {
				return p
			}
		}

		for p != rect.c {
			p.x++
			currVal++
			if currVal == pointVal {
				return p
			}
		}
	}
}

func nextRect(origin rectangle) rectangle {
	return rectangle{
		a: point{origin.a.x - 1, origin.a.y + 1},
		b: point{origin.b.x - 1, origin.b.y - 1},
		c: point{origin.c.x + 1, origin.c.y - 1},
		d: point{origin.d.x + 1, origin.d.y + 1},
	}
}

func distance(a, b point) int {
	res := math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
	return int(res)
}
