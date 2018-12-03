package main

import (
	"fmt"

	"github.com/mwlebour/aoc2018/util"
)

func main() {

	freqDict := make(map[int]int)
	var firstDouble = 0
	var runningFreq = 0
	var frequencies = util.ListToIntList(util.FileToList("input.out"))

Found:
	for {
		for _, freq := range frequencies {
			runningFreq += freq
			if _, ok := freqDict[runningFreq]; ok {
				firstDouble = runningFreq
				break Found
			} else {
				freqDict[runningFreq] = 1
			}

		}
	}
	fmt.Println(firstDouble)

}
