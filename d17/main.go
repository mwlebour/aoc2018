package main

import (
	"flag"
	"fmt"

	"github.com/mwlebour/aoc2018/util"
)

type position struct {
	x, y int
}

type square struct {
	position
	element rune
}

type squareMap map[position]square

type clayMap struct {
	squares        squareMap
	xx, xy, mx, my int
}

func (c clayMap) elementAt(i, j int) rune {
	if v, ok := c.squares[position{i, j}]; ok {
		return v.element
	}
	if j > c.xy {
		return '#'
	}
	return '.'
}

func (c clayMap) String() string {
	var s string
	for j := c.my - 1; j <= c.xy+1; j++ {
		for i := c.mx - 1; i <= c.xx+1; i++ {
			s += string(c.elementAt(i, j))
		}
		s += "\n"
	}
	return s[:len(s)-1]
}

func parse(lines []string) clayMap {
	squares := make(squareMap)
	for _, line := range lines {
		var xory1, xory2 string
		var vein, start, stop int
		fmt.Sscanf(line, "%s = %d, %s = %d..%d", &xory1, &vein, &xory2, &start, &stop)
		if xory1 == "x" {
			for j := start; j <= stop; j++ {
				p := position{vein, j}
				squares[p] = square{p, '#'}
			}
		} else {
			for i := start; i <= stop; i++ {
				p := position{i, vein}
				squares[p] = square{p, '#'}
			}
		}
	}
	i := 0
	var xx, xy, mx, my int
	for p := range squares {
		if i == 0 {
			xx, xy, mx, my = p.x, p.y, p.x, p.y
			i = 1
			continue
		}
		if p.x > xx {
			xx = p.x
		}
		if p.y > xy {
			xy = p.y
		}
		if p.x < mx {
			mx = p.x
		}
		if p.y < my {
			my = p.y
		}
	}
	return clayMap{squares, xx, xy, mx, my}
}

func fall(p position, m clayMap) (position, clayMap) {
	for {
		if m.elementAt(p.x, p.y+1) != '.' && m.elementAt(p.x, p.y+1) != '|' {
			return p, m
		}
		if m.elementAt(p.x, p.y) == '.' {
			m.squares[p] = square{p, '|'}
		}
		p = position{p.x, p.y + 1}
	}
}

func rollLeft(p position, m clayMap) position {
	for {
		if m.elementAt(p.x-1, p.y) != '.' || m.elementAt(p.x, p.y+1) == '.' {
			return p
		}
		p = position{p.x - 1, p.y}
	}
}

func rollRight(p position, m clayMap) position {
	for {
		if m.elementAt(p.x+1, p.y) != '.' || m.elementAt(p.x, p.y+1) == '.' {
			return p
		}
		p = position{p.x + 1, p.y}
	}
}

func findNewSpring(c clayMap) position {
	for j := c.my - 1; j <= c.xy+1; j++ {
		for i := c.mx - 1; i <= c.xx+1; i++ {
			if c.elementAt(i, j) == '~' && c.elementAt(i, j+1) == '.' {
				return position{i, j}
			}
		}
	}
	return position{}
}

func fillWater(springPos position, theMap clayMap) (position, clayMap) {
	// TODO: do not yet have a good algo for determining when to change
	// the fountain position:
	// ......|.......
	// ...~..|.....#.
	// .#..#~~~~...#.
	// .#..#~~#......
	// .#..#~~#......
	// .#~~~~~#......
	// .#~~~~~#......
	// .#######......
	// ..............
	// ..............
	// ....#.....#...
	var depth position
	depth, theMap = fall(springPos, theMap)
	if theMap.elementAt(depth.x-1, depth.y) == '.' {
		depth = rollLeft(depth, theMap)
	} else if theMap.elementAt(depth.x+1, depth.y) == '.' {
		depth = rollRight(depth, theMap)
	}
	if depth == springPos {
		springPos = findNewSpring(theMap)
	} else {
		theMap.squares[depth] = square{depth, '~'}
	}
	return springPos, theMap
}

func main1(theMap clayMap) {
	springPos := position{500, 0}
	for springPos != (position{}) {
		springPos, theMap = fillWater(springPos, theMap)
		fmt.Println(theMap)
	}
	fmt.Println(theMap)
}

func main2(theMap clayMap) {
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
	theMap := parse(input)
	fmt.Println("Starting main1")
	main1(theMap)
	fmt.Println("Starting main2")
	main2(theMap)
	fmt.Println("Done!")
}
