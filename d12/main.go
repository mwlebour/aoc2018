package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mwlebour/aoc2018/util"
)

type pot bool

func (p pot) String() string {
	if p == true {
		return string('#')
	}
	return string('.')
}

type state struct {
	l2, l1, c, r1, r2 pot
}

func (s state) String() string {
	return fmt.Sprintf("%s%s%s%s%s", s.l2, s.l1, s.c, s.r1, s.r2)
}

type hall []pot

func (h hall) StartStop() (int, int) {
	var start, stop int
	for i := 0; i < len(h); i++ {
		if h[i] {
			start = i
			break
		}
	}
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] {
			stop = i
			break
		}
	}
	return start, stop
}
func (h hall) String() string {
	var s string
	started := false
	for _, p := range h {
		if p {
			started = true
		}
		if started {
			s += p.String()
		}
	}
	return s
}

func (h hall) Sum(start int) int {
	s := 0
	for i, v := range h {
		if v {
			s += (i - start)
		}
	}
	return s
}

type stateMap map[state]pot

func (m stateMap) String() string {
	r := make([]string, 0, len(m))
	for s, p := range m {
		r = append(r, fmt.Sprintf("%s => %s", s, p))
	}
	return strings.Join(r, "\n")
}

func parse(lines []string) (hall, stateMap) {
	hall := stringToHall(strings.Fields(lines[0])[2])
	m := make(stateMap)
	for _, l := range lines[2:] {
		k, v := stringToStatePot(l)
		m[k] = v
	}
	return hall, m
}

func stringToHall(s string) hall {
	h := make(hall, 0)
	for _, r := range s {
		h = append(h, runeToPot(r))
	}
	return h
}

func runeToPot(r rune) pot {
	if r == '#' {
		return true
	}
	return false
}

func hallToState(h hall) state {
	return state{h[0], h[1], h[2], h[3], h[4]}
}

func stringToStatePot(s string) (state, pot) {
	var l, r string
	fmt.Sscanf(s, "%s => %s", &l, &r)
	h := stringToHall(l)
	return hallToState(h), runeToPot(rune(r[0]))
}

func generate(h hall, m stateMap) hall {
	newH := make(hall, len(h), len(h))
	for i := 2; i <= len(h)-3; i++ {
		newH[i] = m[hallToState(h[i-2:i+3])]
	}
	return newH
}

func main1(h hall, m stateMap) {
	generations := 3000
	h = append(make(hall, generations*2, generations*2), h...)
	h = append(h, make(hall, generations*2, generations*2)...)
	for i := 1; i <= generations; i++ {
		h = generate(h, m)
		if i%1000 == 0 {
			fmt.Println(i, h.Sum(generations*2))
		}
	}
	fmt.Println(h.Sum(generations * 2))
}

func main2(h hall, m stateMap) {
	n := 20000
	step := 1000
	start := 384
	inc := 38000
	for i := 0; i <= n/step; i++ {
		fmt.Println(i*step, inc*i+start)
	}
	n = 50000000000
	fmt.Println(n/step*inc + start)
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
	h, m := parse(input)
	main1(h, m)
	fmt.Println("Starting main2")
	main2(h, m)
	fmt.Println("Done!")
}
