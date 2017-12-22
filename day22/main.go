package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	sideSize := len(lines)
	grid, grid2 := NewSquare(sideSize), NewSquare(sideSize)
	for i := range lines {
		for j, c := range lines[i] {
			if c == '#' {
				grid[i][j] = infected
				grid2[i][j] = infected
			} else if c == '.' {
				grid[i][j] = clean
				grid2[i][j] = clean
			} else {
				panic("Alarm! Malformed input!")
			}
		}
	}

	fmt.Println("Part 1:", partOne(grid, sideSize))
	fmt.Println("Part 2:", partTwo(grid2, sideSize))
}

func partOne(grid Square, sideSize int) int {
	bursts, infectedFlags := 10000, 0

	center := int(math.Ceil(float64(sideSize / 2)))
	p := Point{center, center, up}

	for i := 0; i < bursts; i++ {
		if grid[p.i][p.j] == infected {
			p.direction = (p.direction + 1) % 4
			grid[p.i][p.j] = clean
		} else {
			if tmp := p.direction - 1; tmp < 0 {
				p.direction = 3
			} else {
				p.direction = tmp
			}
			grid[p.i][p.j] = infected
			infectedFlags++
		}

		p.Move()
		grid = extend(grid, &sideSize, &p)
	}
	return infectedFlags
}

func partTwo(grid Square, sideSize int) int {
	bursts, infectedFlags := 10000000, 0

	center := int(math.Ceil(float64(sideSize / 2)))
	p := Point{center, center, up}

	for i := 0; i < bursts; i++ {
		if f := grid[p.i][p.j]; f == infected {
			p.direction = (p.direction + 1) % 4
			grid[p.i][p.j] = flagged
		} else if f == clean {
			if tmp := p.direction - 1; tmp < 0 {
				p.direction = 3
			} else {
				p.direction = tmp
			}
			grid[p.i][p.j] = weakened
		} else if f == flagged {
			p.direction = (p.direction + 2) % 4
			grid[p.i][p.j] = clean
		} else if f == weakened {
			grid[p.i][p.j] = infected
			infectedFlags++
		}

		p.Move()
		grid = extend(grid, &sideSize, &p)
	}
	return infectedFlags
}

type direction int8

const (
	up direction = iota
	right
	down
	left
)

type Point struct {
	i, j int
	direction
}

func (p *Point) Move() {
	switch p.direction {
	case up:
		p.i--
	case down:
		p.i++
	case right:
		p.j++
	case left:
		p.j--
	}
}

type Flag int8

const (
	clean Flag = iota
	infected
	weakened
	flagged
)

type Square [][]Flag

func NewSquare(sideSize int) Square {
	s := make(Square, sideSize)
	for i := range s {
		s[i] = make([]Flag, sideSize)
	}
	return s
}

func extend(grid Square, sideSize *int, p *Point) Square {
	if p.i >= *sideSize || p.j >= *sideSize {
		newLine := make([]Flag, *sideSize+1)
		grid = append(grid, newLine)

		for i := range grid {
			grid[i] = append(grid[i], clean)
		}

		*sideSize++
	}
	if p.i < 0 || p.j < 0 {
		for i := range grid {
			grid[i] = append([]Flag{clean}, grid[i]...)
		}

		newLine := make([]Flag, *sideSize+1)
		grid = append(Square{newLine}, grid...)

		*sideSize++
		p.i++
		p.j++
	}

	return grid
}
