package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
)

const GridSize = 128

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data = append(data, '-')
	hashes := [GridSize]string{}
	for i := range hashes {
		input := make([]byte, len(data))
		copy(input, data)
		input = append(input, []byte(strconv.FormatInt(int64(i), 10))...)
		hashes[i] = knotHash(input)
	}
	grid := &[GridSize][GridSize]bool{}
	squares := 0
	for i, hash := range hashes {
		for j, _ := range hash {
			num, _ := strconv.ParseInt(string(hash[j]), 16, 8)
			for k := 0; k < 4; k++ {
				bit := num&(1<<uint(3-k)) != 0
				if bit {
					squares++
				}
				grid[i][j*4+k] = bit
			}
		}
	}
	fmt.Println("Part 1:", squares)

	visited := &[GridSize][GridSize]bool{}
	groups := 0
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if grid[i][j] && !visited[i][j] {
				dfs(grid, visited, i, j)
				groups++
			}
		}
	}
	fmt.Println("Part 2:", groups)
}

type point struct {
	x, y int
}

var cellNeighbours = [...]point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func dfs(grid, visited *[GridSize][GridSize]bool, i, j int) {
	visited[i][j] = true
	for k := 0; k < len(cellNeighbours); k++ {
		row := i + cellNeighbours[k].y
		col := j + cellNeighbours[k].x
		if (row >= 0) && (row < GridSize) &&
			(col >= 0) && (col < GridSize) &&
			(grid[row][col] && !visited[row][col]) {
			dfs(grid, visited, row, col)
		}
	}
}

func knotHash(input []byte) string {
	input = append(input, []byte{17, 31, 73, 47, 23}...)
	listLen := 256
	list := genList(listLen)
	pos, skipSize := 0, 0

	for i := 0; i < 64; i++ {
		for _, length := range input {
			for idx := 0; idx < (int(length) / 2); idx++ {
				i := (pos + idx) % listLen
				j := (pos + int(length) - 1 - idx) % listLen
				list[i], list[j] = list[j], list[i]
			}
			pos = (pos + int(length) + skipSize) % listLen
			skipSize++
		}
	}

	hash := make([]byte, 16)
	for i := 0; i < 16; i++ {
		hash[i] = byte(list[i*16])
		for j := i*16 + 1; j < i*16+16; j++ {
			hash[i] ^= byte(list[j])
		}
	}

	return hex.EncodeToString(hash)
}

func genList(listSize int) []int {
	list := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = i
	}
	return list
}
