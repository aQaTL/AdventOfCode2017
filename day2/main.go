package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	partOne(bytes)
	partTwo(bytes)
}

func partOne(input []byte) {
	sum := 0
	for _, row := range strings.Split(string(input), "\n") {
		min, max := (1<<31)-1, -(1 << 31)
		for _, cell := range strings.Split(row, "\x09") {
			num, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		sum += max - min
	}
	fmt.Println(sum)
}

func partTwo(input []byte) {
	sum := 0
	for _, row := range strings.Split(string(input), "\n") {
		nums, err := stringSliceToInt(strings.Split(row, "\x09"))
		if err != nil {
			panic(err)
		}

		sort.Ints(nums)

		for i := len(nums) - 1; i > 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if nums[i]%nums[j] == 0 {
					sum += nums[i] / nums[j]
				}
			}
		}
	}
	fmt.Println(sum)
}

func stringSliceToInt(slice []string) ([]int, error) {
	nums := make([]int, len(slice))
	for i, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}
