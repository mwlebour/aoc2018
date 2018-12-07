package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mwlebour/aoc2018/util"
)

type coordinate struct {
	x int
	y int
}

func translate(coords []string) []coordinate {
	newCoords := make([]coordinate, 0, len(coords))
	for _, s := range coords {
		t := strings.Split(s, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(t[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(t[1]))
		newCoords = append(newCoords, coordinate{x, y})
	}
	return newCoords
}

func findBounds(coords []coordinate) []int {
	bounds := make([]int, 4, 4)
	// x lower
	bounds[0] = 999999
	// x higher
	bounds[1] = -999999
	// y lower
	bounds[2] = 999999
	// y higher
	bounds[3] = -999999

	for _, coord := range coords {
		if coord.x < bounds[0] {
			bounds[0] = coord.x
		}
		if coord.x > bounds[1] {
			bounds[1] = coord.x
		}
		if coord.y < bounds[2] {
			bounds[2] = coord.y
		}
		if coord.y > bounds[3] {
			bounds[3] = coord.y
		}
	}
	return bounds
}

func buildMap(bounds []int) [][]int {
	a := make([][]int, bounds[1]-bounds[0]+1)
	ydist := bounds[3] - bounds[2] + 1
	for i := range a {
		a[i] = make([]int, ydist)
	}
	return a
}

func buildCoordMap(coords []coordinate) map[int]map[int]int {
	coordMap := make(map[int]map[int]int)
	i := 0
	for _, coord := range coords {
		if _, ok := coordMap[coord.x]; !ok {
			coordMap[coord.x] = make(map[int]int)
		}
		coordMap[coord.x][coord.y] = i
		i++
	}
	return coordMap
}

func manhatten(a coordinate, b coordinate) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func findNearest(coords []coordinate, x int, y int) (coordinate, error) {
	// map of distance to n coords
	minDistMap := make(map[int][]coordinate)
	for _, coord := range coords {
		d := manhatten(coordinate{x, y}, coord)
		if _, ok := minDistMap[d]; !ok {
			minDistMap[d] = make([]coordinate, 0)
		}
		minDistMap[d] = append(minDistMap[d], coord)
	}
	m := 999999
	for k := range minDistMap {
		if k < m {
			m = k
		}
	}
	if len(minDistMap[m]) > 1 {
		return coordinate{}, errors.New("no single match")
	}
	return minDistMap[m][0], nil
}

func bounded(coord coordinate, bounds []int) bool {
	return coord.x != bounds[0] && coord.x != bounds[1] &&
		coord.y != bounds[2] && coord.y != bounds[3]
}

func main1(coords []coordinate) {
	// unfortunately, not a full solution, `bounded` needs work
	coordMap := buildCoordMap(coords)
	bounds := findBounds(coords)
	boundMap := buildMap(bounds)
	sumMap := make(map[int]int)
	for x, ymap := range boundMap {
		for y := range ymap {
			nearestCoord, err := findNearest(coords, x+bounds[0], y+bounds[2])
			if err == nil {
				coordN := coordMap[nearestCoord.x][nearestCoord.y]
				ymap[y] = coordN
				if bounded(nearestCoord, bounds) {
					sumMap[coordN]++
				}
			} else {
				ymap[y] = 0
			}
			//fmt.Printf("%2d", ymap[y])
		}
		//fmt.Print("\n")
	}
	for k, v := range sumMap {
		fmt.Println(k, v)
	}
}

func totalManhatten(coords []coordinate, x int, y int) int {
	total := 0
	for _, coord := range coords {
		total += manhatten(coordinate{x, y}, coord)
	}
	return total
}

func main2(coords []coordinate) {
	bounds := findBounds(coords)
	boundMap := buildMap(bounds)
	islandSize := 0
	for x, ymap := range boundMap {
		for y := range ymap {
			if totalManhatten(coords, x+bounds[0], y+bounds[2]) < 10000 {
				islandSize++
			}
		}
	}
	fmt.Println(islandSize)
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
	coords := translate(input)
	main1(coords)
	main2(coords)
}
