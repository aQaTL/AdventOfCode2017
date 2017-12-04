package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"sort"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n%d\n", partOne(bytes), partTwo(bytes))
}

type byteSlice []byte

func (s byteSlice) Len() int           { return len(s) }
func (s byteSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byteSlice) Less(i, j int) bool { return s[i] < s[j] }

func partTwo(input []byte) int {
	count := 0
	for _, passphrase := range strings.Split(string(input), "\n") {
		words := strings.Split(passphrase, " ")
		valid := true

		for i := 0; i < len(words)-1; i++ {
			for j := i + 1; j < len(words); j++ {
				if len(words[i]) == len(words[j]) {
					b1 := byteSlice(words[i])
					sort.Sort(b1)
					b2 := byteSlice(words[j])
					sort.Sort(b2)

					if string(b1) == string(b2) {
						valid = false
						break
					}
				}
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func partOne(input []byte) int {
	count := 0
	for _, passphrase := range strings.Split(string(input), "\n") {
		words := make(map[string]bool)

		valid := true
		for _, word := range strings.Split(passphrase, " ") {
			if words[word] {
				valid = false
				break
			} else {
				words[word] = true
			}
		}
		if valid {
			count++
		}
	}
	return count
}
