package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	diag := make(grid, len(lines))
	for i := range diag {
		diag[i] = []rune(lines[i])
	}

	p := &point{findEntry(diag), 0, down}
	letters := make([]rune, 0)
	steps := 1

	for {
		p.Move()
		if c, ok := diag.GetByPoint(*p); !ok {
			break
		} else if unicode.IsLetter(c) {
			steps++
			letters = append(letters, c)
		} else if c == '+' {
			steps++
			if p.direction == up || p.direction == down {
				if c, ok := diag.Get(p.y, p.x+1); ok && (c == '-' || unicode.IsLetter(c)) {
					p.direction = right
				} else {
					p.direction = left
				}
			} else {
				if c, ok := diag.Get(p.y+1, p.x); ok && (c == '|' || unicode.IsLetter(c)) {
					p.direction = down
				} else {
					p.direction = up
				}
			}
		} else if c == ' ' {
			break
		} else {
			steps++
		}
	}
	fmt.Println("Part 1:", string(letters))
	fmt.Println("Part 2:", steps)
}

type grid [][]rune

func (g grid) GetByPoint(p point) (rune, bool) {
	return g.Get(p.y, p.x)
}

func (g grid) Get(y, x int) (rune, bool) {
	if y >= len(g) || x >= len(g[y]) {
		return 0, false
	}
	return g[y][x], true
}

type direction int

const (
	up direction = iota
	down
	right
	left
)

type point struct {
	x, y int
	direction
}

func (p *point) Move() {
	switch p.direction {
	case up:
		p.y--
	case down:
		p.y++
	case right:
		p.x++
	case left:
		p.x--
	}
}

func findEntry(diag [][]rune) int {
	for i, s := range diag[0] {
		if s == '|' {
			return i
		}
	}
	return -1
}
