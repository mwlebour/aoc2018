package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mwlebour/aoc2018/util"
)

// Claim represents an elven claim on chimney-squeeze cloth
type Claim struct {
	id     int
	xoff   int
	yoff   int
	width  int
	height int
}

func parseClaim(s string) Claim {
	fields := strings.Fields(s)
	id, _ := strconv.Atoi(strings.Trim(fields[0], "#"))
	t := strings.Split(strings.Trim(fields[2], ":"), ",")
	xoff, _ := strconv.Atoi(t[0])
	yoff, _ := strconv.Atoi(t[1])
	t = strings.Split(fields[3], "x")
	width, _ := strconv.Atoi(t[0])
	height, _ := strconv.Atoi(t[1])
	return Claim{id, xoff, yoff, width, height}
}

type square struct {
	x int
	y int
}

func allSquares(c Claim) []square {
	m := make([]square, 0)
	for x := c.xoff; x < c.xoff+c.width; x++ {
		for y := c.yoff; y < c.yoff+c.height; y++ {
			m = append(m, square{x, y})
		}
	}
	return m
}

func buildClaims(stringlyClaimList []string) []Claim {
	claimList := make([]Claim, 0)
	for _, s := range stringlyClaimList {
		claimList = append(claimList, parseClaim(s))
	}
	return claimList
}

func buildClaimMap(claimList []Claim) map[square]int {
	claimMap := make(map[square]int)
	for _, claim := range claimList {
		for _, v := range allSquares(claim) {
			claimMap[v]++
		}
	}
	return claimMap
}

func main1(stringlyClaimList []string) {

	claimList := buildClaims(stringlyClaimList)
	claimMap := buildClaimMap(claimList)

	i := 0
	for _, v := range claimMap {
		if v > 1 {
			i++
		}
	}
	fmt.Println(i)
}

func main2(stringlyClaimList []string) {

	claimList := buildClaims(stringlyClaimList)
	claimMap := buildClaimMap(claimList)
Out:
	for _, claim := range claimList {
		for _, cell := range allSquares(claim) {
			if claimMap[cell] > 1 {
				continue Out
			}
		}
		fmt.Println(claim)
		// break
	}
}

var full bool

func init() {
	flag.BoolVar(&full, "full", false, "run full solution")
}

func main() {
	flag.Parse()
	var input = util.FileToList("unit.out")
	if full {
		input = util.FileToList("input.out")
	}
	main1(input)
	main2(input)
}
