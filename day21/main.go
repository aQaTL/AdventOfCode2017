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

	enhs := make([]enh, len(lines))

	for i, line := range lines {
		io := strings.Split(line, " => ")

		in0 := strings.Split(io[0], "/")
		size := len(in0) - 1

		in := make(rule, len(in0))

		for i, row := range in0 {
			in[i] = []rune(row)
		}

		out1 := strings.Split(io[1], "/")
		out := make(rule, len(out1))

		for i, row := range out1 {
			out[i] = []rune(row)
		}

		in90r := NewRule(len(in0))
		in90l := NewRule(len(in0))
		in180 := NewRule(len(in0))
		flipH := NewRule(len(in0))
		flipV := NewRule(len(in0))
		for i := range in {
			for j := range in[i] {
				in90r[j][size-i] = in[i][j]
				in90l[size-j][i] = in[i][j]
				in180[size-i][size-j] = in[i][j]
				flipH[size-i][j] = in[i][j]
				flipV[i][size-j] = in[i][j]
			}
		}

		transpose := NewRule(len(in0))
		diagonal := NewRule(len(in0))
		for i := range in90r {
			for j := range in90r {
				transpose[i][size-j] = in90r[i][j]
				diagonal[size-i][j] = in90r[i][j]
			}
		}

		enhs[i] = enh{
			in,
			out,
			in90r,
			in90l,
			in180,
			flipH,
			flipV,
			transpose,
			diagonal,
		}
	}

	r := rule{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}

	for iter := 0; iter < 18; iter++ {
		newSize := 0
		if r.Stride()%2 == 0 {
			newSize = 2
		} else if r.Stride()%3 == 0 {
			newSize = 3
		} else {
			fmt.Println("alarm!")
		}

		rules := make([]rule, 0)

		for groupI := 0; groupI < r.Stride()/newSize; groupI++ {
			for groupJ := 0; groupJ < r.Stride()/newSize; groupJ++ {
				group := NewRule(newSize)
				for i := 0; i < newSize; i++ {
					for j := 0; j < newSize; j++ {
						group[i][j] = r[groupI*newSize+i][groupJ*newSize+j]
					}
				}
				for _, e := range enhs {
					if group.Equals(e.in) ||
						group.Equals(e.in90r) ||
						group.Equals(e.in90l) ||
						group.Equals(e.in180) ||
						group.Equals(e.inFlipH) ||
						group.Equals(e.inFlipV) ||
						group.Equals(e.transpose) ||
						group.Equals(e.diagonal) {
						rules = append(rules, e.out)
					}
				}
			}
		}

		newSize++
		groupsSqrt := int(math.Sqrt(float64(len(rules))))

		r = NewRule(newSize * groupsSqrt)

		for groupI := 0; groupI < groupsSqrt; groupI++ {
			for groupJ := 0; groupJ < groupsSqrt; groupJ++ {
				rule := rules[groupI*groupsSqrt+groupJ]
				for i := 0; i < rule.Stride(); i++ {
					for j := 0; j < rule.Stride(); j++ {
						r[groupI*newSize+i][groupJ*newSize+j] = rule[i][j]
					}
				}
			}
		}

		if iter == 5 {
			fmt.Println("Part 1:", calcOn(r))
		}
	}

	fmt.Println("Part 2:", calcOn(r))
}

func calcOn(r rule) int {
	on := 0
	for i := 0; i < r.Stride(); i++ {
		for j := 0; j < r.Stride(); j++ {
			if r[i][j] == '#' {
				on++
			}
		}
	}
	return on
}

type rule [][]rune

func (r rule) Stride() int {
	return len(r)
}

func (r rule) Equals(x rule) bool {
	if r.Stride() != x.Stride() {
		return false
	}
	for i := 0; i < r.Stride(); i++ {
		for j := 0; j < r.Stride(); j++ {
			if r[i][j] != x[i][j] {
				return false
			}
		}
	}
	return true
}

func NewRule(size int) rule {
	rule := make(rule, size)
	for i := range rule {
		rule[i] = make([]rune, size)
	}
	return rule
}

type enh struct {
	in, out                                                    rule
	in90r, in90l, in180, inFlipH, inFlipV, transpose, diagonal rule
}
