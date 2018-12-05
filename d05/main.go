package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"

	"github.com/mwlebour/aoc2018/util"
)

func react(polymer string) int {
	i := len(polymer) - 1
	for {
		if i == 0 {
			break
		}
		if i >= len(polymer) {
			i = len(polymer) - 1
		}
		up1 := unicode.ToUpper(rune(polymer[i-1]))
		up2 := unicode.ToUpper(rune(polymer[i]))
		if up1 == up2 && polymer[i] != polymer[i-1] {
			polymer = polymer[:i-1] + polymer[i+1:]
		} else {
			i--
		}
	}
	return len(polymer)
}

func main1(polymer string) {
	fmt.Println(react(polymer))
}

func getUnique(polymer string) []byte {
	uniqMap := make(map[byte]struct{})
	for _, v := range polymer {
		v = unicode.ToUpper(rune(v))
		uniqMap[byte(v)] = struct{}{}
	}
	// https://stackoverflow.com/a/27848197
	keys := make([]byte, len(uniqMap))
	i := 0
	for k := range uniqMap {
		keys[i] = k
		i++
	}
	return keys
}

func removeUnits(b byte, polymer string) string {
	t := strings.Replace(polymer, string(unicode.ToUpper(rune(b))), "", -1)
	t = strings.Replace(t, string(unicode.ToLower(rune(b))), "", -1)
	return t
}

func main2(polymer string) {
	uniqueUnits := getUnique(polymer)
	unitMap := make(map[string]int)
	for _, u := range uniqueUnits {
		updatedPolymer := removeUnits(u, polymer)
		unitMap[string(u)] = react(updatedPolymer)
	}
	m := react(polymer)
	for _, v := range unitMap {
		if v < m {
			m = v
		}
	}
	fmt.Println(m)
}

var full bool

func init() {
	flag.BoolVar(&full, "full", false, "run full solution")
}

func main() {
	flag.Parse()
	input := "dabAcCaCBAcCcaDA"
	if full {
		input = util.FileToList("input.out")[0]
	}
	main1(input)
	main2(input)
}
