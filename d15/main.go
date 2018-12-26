package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mwlebour/aoc2018/util"
)

func main1(f floor) int {
	fmt.Println(f)
	return 0
}

func main2() {
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
	pairs := parse(input)
	for _, pair := range pairs {
		fmt.Println(main1(pair.f), pair.a)
	}
	fmt.Println("Starting main2")
	main2()
	fmt.Println("Done!")
}

type position struct {
	x, y int
}
type positionList []position
type positionMap map[position]struct{}
type cell struct {
	position
	wall bool
}
type cellList []cell
type cellMap map[position]cell
type floor struct {
	cells   cellMap
	nx, ny  int
	goblins positionMap
	elves   positionMap
}

func (f floor) String() string {
	var s string
	for j := 0; j < f.ny; j++ {
		for i := 0; i < f.nx; i++ {
			p := position{i, j}
			if f.cells[p].wall {
				s += "#"
			} else if _, ok := f.goblins[p]; ok {
				s += "G"
			} else if _, ok := f.elves[p]; ok {
				s += "E"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s[:len(s)-1]
}

func cellMapFromCells(cells cellList) cellMap {
	myMap := make(cellMap)
	for _, c := range cells {
		myMap[c.position] = c
	}
	return myMap
}

func positionMapFromPositions(positions positionList) positionMap {
	myMap := make(positionMap)
	for _, p := range positions {
		myMap[p] = struct{}{}
	}
	return myMap
}

func cellSort(a, b cell) bool {
	if a.y < b.y {
		return true
	}
	return a.x < b.x
}

type floorAnswerPair struct {
	f floor
	a int
}

func lineToInt(line string) int {
	i, _ := strconv.Atoi(strings.Trim(line, ""))
	return i
}

func parse(lines []string) []floorAnswerPair {
	pairs := make([]floorAnswerPair, 0)
	theseCells := make(cellList, 0)
	goblins := make(positionList, 0)
	elves := make(positionList, 0)
	var ny, nx int
	for _, line := range lines {
		if line[0] != '#' {
			pairs = append(pairs,
				floorAnswerPair{
					floor{
						cellMapFromCells(theseCells),
						nx + 1,
						ny,
						positionMapFromPositions(goblins),
						positionMapFromPositions(elves),
					},
					lineToInt(line),
				},
			)
			theseCells = make(cellList, 0)
			goblins = make(positionList, 0)
			elves = make(positionList, 0)
			ny = 0
			nx = 0
			continue
		}
		for x, r := range line {
			theseCells = append(theseCells, cell{position{x, ny}, r == '#'})
			if r == 'G' {
				goblins = append(goblins, position{x, ny})
			}
			if r == 'E' {
				elves = append(elves, position{x, ny})
			}
			if x > nx {
				nx = x
			}
		}
		ny++

	}
	return pairs
}
