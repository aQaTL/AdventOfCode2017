package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n%s\n",
		partOne(strings.Split(string(input), ",")),
		partTwo(input))
}

func partTwo(input []byte) string {
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
		res := byte(list[i*16])
		for j := i*16 + 1; j < i*16+16; j++ {
			res ^= byte(list[j])
		}
		hash[i] = res
	}

	return hex.EncodeToString(hash)
}

func partOne(input []string) int {
	lengths := make([]int, len(input))
	for i, s := range input {
		num, _ := strconv.Atoi(s)
		lengths[i] = num
	}

	listLen := 256
	list := genList(listLen)
	pos, skipSize := 0, 0

	for _, length := range lengths {
		for idx := 0; idx < length/2; idx++ {
			i := (pos + idx) % listLen
			j := (pos + length - 1 - idx) % listLen
			list[i], list[j] = list[j], list[i]
		}
		pos = (pos + length + skipSize) % listLen
		skipSize++
	}

	return list[0] * list[1]
}

func genList(listSize int) []int {
	list := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = i
	}
	return list
}
