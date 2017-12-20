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
	dataStr := strings.
		NewReplacer("p=<", "", "v=<", "", "a=<", "", ">", "").
		Replace(string(data))
	lines := strings.Split(dataStr, "\n")

	particles := make([]particle, len(lines))
	for i := range particles {
		vecs := strings.Split(lines[i], ", ")
		points := [3]point{}
		for i, vec := range vecs {
			p := point{}
			fmt.Sscanf(vec, "%d,%d,%d", &p.x, &p.y, &p.z)
			points[i] = p
		}
		particles[i] = particle{points[0], points[1], points[2], i}
	}
	particles2 := make([]particle, len(particles))
	copy(particles2, particles)

	fmt.Println("Part 1:", partOne(particles))
	fmt.Println("Part 2", partTwo(particles2))
}

func partOne(particles []particle) int {
	zero := particle{}
	for i := 0; i < 100000; i++ {
		for i := range particles {
			process(&particles[i])
		}
	}
	least := (1 << 63) - 1
	leastP := particle{}
	for i, part := range particles {
		dst := distance(&part, &zero)
		if int(dst) < least {
			least = int(dst)
			leastP = particles[i]
		}
	}
	return leastP.idx
}

func partTwo(particles []particle) int {
	for i := 0; i < 100000; i++ {
		for i := range particles {
			process(&particles[i])
		}
		idxs := make([]int, len(particles))
		for i := 0; i < len(particles)-1; i++ {
			for j := i + 1; j < len(particles); j++ {
				if particles[i].p == particles[j].p {
					idxs[i] = 1
					idxs[j] = 1
				}
			}
		}
		deleted := 0
		for idx, val := range idxs {
			if val == 1 {
				particles = append(particles[:idx-deleted], particles[idx-deleted+1:]...)
				deleted++
			}
		}
	}
	return len(particles)
}

func process(p *particle) {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z

	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func distance(a, b *particle) float64 {
	return math.Abs(float64(a.p.x-b.p.x)) +
		math.Abs(float64(a.p.y-b.p.y)) +
		math.Abs(float64(a.p.z-b.p.z))
}

type particle struct {
	p, v, a point
	idx     int
}

type point struct {
	x, y, z int
}
