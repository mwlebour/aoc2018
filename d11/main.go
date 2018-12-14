package main

import (
	"flag"
	"fmt"
)

func hundo(i int) int {
	return (i % 1000) / 100
}

func cellValue(serial, x, y int) (v int) {
	return hundo(((x+10)*y+serial)*(x+10)) - 5
}

type pos struct {
	x, y int
}

type varpos struct {
	x, y, w int
}

func cellValues(serial int) map[pos]int {
	cells := make(map[pos]int)
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			cells[pos{i, j}] = cellValue(serial, i, j)
		}
	}
	return cells
}

func main1(serial int) (pos, int) {
	cells := cellValues(serial)
	grids := make(map[pos]int)
	for i := 1; i <= 297; i++ {
		for j := 1; j <= 297; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					grids[pos{i, j}] += cells[pos{i + k, j + l}]
				}
			}
		}
	}
	vmax := -10000000
	pmax := pos{}
	for p, v := range grids {
		if v > vmax {
			vmax = v
			pmax = p
		}
	}
	return pmax, vmax
}

func buildGrid(cells map[pos]int, w int) map[varpos]int {
	grids := make(map[varpos]int)
	for i := 1; i <= 300-w; i++ {
		for j := 1; j <= 300-w; j++ {
			this := varpos{i, j, w}
			// if w > 1 {
			// 	grids[this] = grids[varpos{i, j, w - 1}]
			// }
			// for k := 0; k < w-1; k++ {
			// 	grids[this] += cells[pos{i + w - 1, j + k}]
			// 	grids[this] += cells[pos{i + k, j + w - 1}]
			// }
			// grids[this] += cells[pos{i + w - 1, j + w - 1}]
			for k := 0; k < w; k++ {
				for l := 0; l < w; l++ {
					grids[this] += cells[pos{i + k, j + l}]
				}
			}
		}
	}
	return grids
}

type maxmax struct {
	x, y, w, v int
}

func main2(serial int) {
	cells := cellValues(serial)
	c := make(chan maxmax, 6)
	for w := 1; w <= 300; w++ {
		go func(w int) {
			grids := buildGrid(cells, w)
			vmax := -10000000
			pmax := varpos{}
			for p, v := range grids {
				if v > vmax {
					vmax = v
					pmax = p
				}
			}
			c <- maxmax{pmax.x, pmax.y, pmax.w, vmax}
		}(w)
	}
	for m := range c {
		fmt.Println(m)
	}
}

func main() {
	flag.Parse()
	fmt.Println("Starting main1")
	//fmt.Println(4, cellValue(8, 3, 5))
	//fmt.Println(-5, cellValue(57, 122, 79))
	//fmt.Println(0, cellValue(39, 217, 196))
	//fmt.Println(4, cellValue(71, 101, 153))
	p, v := main1(18)
	fmt.Println(33, 45, 29, p, v)
	p, v = main1(42)
	fmt.Println(21, 61, 30, p, v)
	p, v = main1(2694)
	fmt.Println(p, v)
	fmt.Println("Starting main2")
	// vp, v = main2(18)
	// fmt.Println(vp, v)
	// vp, v = main2(42)
	// fmt.Println(vp, v)
	main2(2694)
	fmt.Println("Done!")
}
