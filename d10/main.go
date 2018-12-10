package main

import (
	"flag"
	"fmt"

	"github.com/mwlebour/aoc2018/util"
)

// Particle particles
type Particle struct {
	x  int
	y  int
	vx int
	vy int
}

func (p Particle) move() Particle {
	return Particle{p.x + p.vx, p.y + p.vy, p.vx, p.vy}
}

// Particles particles
type Particles []Particle

func (particles Particles) bounds() (int, int, int, int) {
	var left, right, up, down int
	for i, p := range particles {
		if i == 0 {
			left = p.x
			right = p.x
			up = p.y
			down = p.y
			continue
		}
		if p.y < up {
			up = p.y
		}
		if p.y > down {
			down = p.y
		}
		if p.x < left {
			left = p.x
		}
		if p.x > right {
			right = p.x
		}
	}
	return left, right, up, down
}

func (particles Particles) area() int {
	left, right, up, down := particles.bounds()
	return (down - up) * (right - left)
}

type pos struct {
	x int
	y int
}

func (particles Particles) show() {
	left, right, up, down := particles.bounds()
	picture := make(map[pos]struct{})
	for _, p := range particles {
		picture[pos{p.x, p.y}] = struct{}{}
	}
	for y := up; y <= down; y++ {
		for x := left; x <= right; x++ {
			if _, ok := picture[pos{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func particlesFromInput(input []string) Particles {
	particles := make(Particles, len(input), len(input))
	for i, line := range input {
		var x, y, vx, vy int
		fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &x, &y, &vx, &vy)
		particles[i] = Particle{x, y, vx, vy}
	}
	return particles
}

func main1(particles Particles) {
	lastArea := 9999999999999
	i := 0
	for {
		newParticles := make(Particles, len(particles), len(particles))
		for i, p := range particles {
			newParticles[i] = p.move()
		}
		thisArea := newParticles.area()
		if thisArea < lastArea {
			lastArea = thisArea
			particles = newParticles
		} else {
			// let's see if we can assume puzzles were generated and then
			// reverse velocitied such that when we get the smallest area,
			// we should print
			break
		}
		i++
	}
	particles.show()
	fmt.Println(i)
}

func main2(particles Particles) {
}

var full bool

func init() {
	flag.BoolVar(&full, "full", false, "run full solution")
}

func main() {
	flag.Parse()
	input := util.FileToList("unit.out")
	if full {
		input = util.FileToList("input.out")
	}
	fmt.Println("Starting main1")
	particles := particlesFromInput(input)
	main1(particles)
	fmt.Println("Starting main2")
	main2(particles)
	fmt.Println("Done!")
}
