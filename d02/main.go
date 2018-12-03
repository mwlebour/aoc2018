package main

import (
	"fmt"
	"math"

	"github.com/mwlebour/aoc2018/util"
)

func frequencyAnalysis(id string) map[rune]int {

	m := make(map[rune]int)
	for _, v := range id {
		m[v]++
	}
	return m
}

func count2count3(m map[rune]int) (bool, bool) {
	m2 := make(map[int]int)
	for _, v := range m {
		m2[v]++
	}
	is2 := false
	is3 := false
	if v, ok := m2[2]; ok && v > 0 {
		is2 = true
	}
	if v, ok := m2[3]; ok && v > 0 {
		is3 = true
	}
	return is2, is3
}

func equalby1(lhs map[rune]int, rhs map[rune]int) bool {
	if math.Abs(float64(len(rhs)-len(lhs))) > 1 {
		return false
	}
	diffs := 0
	for k, lhsv := range lhs {
		if rhsv, ok := rhs[k]; ok {
			diffs += int(math.Abs(float64(lhsv - rhsv)))
		} else {
			diffs += lhsv
		}
		if diffs > 1 {
			return false
		}
	}
	return true
}

func main1(ids []string) {
	running2 := 0
	running3 := 0
	for _, id := range ids {
		m := frequencyAnalysis(id)
		is2, is3 := count2count3(m)
		if is2 {
			running2++
		}
		if is3 {
			running3++
		}
	}
	fmt.Println(running2 * running3)
}

func main2(ids []string) {
	freqAnalysisMap := make(map[string]map[rune]int)
	for _, id := range ids {
		freqAnalysisMap[id] = frequencyAnalysis(id)
	}
	for lhsid, lhsAnalysis := range freqAnalysisMap {
		for rhsid, rhsAnalysis := range freqAnalysisMap {
			if lhsid == rhsid {
				continue
			}
			if equalby1(lhsAnalysis, rhsAnalysis) {
				fmt.Println(lhsid, rhsid)
			}
		}
	}
	fmt.Println(equalby1(frequencyAnalysis("fghij"), frequencyAnalysis("fguij")))
}

func main() {
	var ids = util.FileToList("input.out")
	main1(ids)
	main2(ids)
}
