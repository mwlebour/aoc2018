package main

import (
	"flag"
	"fmt"
	"strconv"
)

func newScores(s1, s2 int) []int {
	sum := s1 + s2
	nums := make([]int, 0, 2)
	if sum < 10 {
		nums = append(nums, sum)
	} else {
		nums = append(nums, 1)
		nums = append(nums, sum%10)
	}
	return nums
}

func scoresToString(scores []int) string {
	s := ""
	for _, v := range scores {
		s += strconv.Itoa(v)
	}
	return s
}

func main1(serial int) string {
	scores := make([]int, 2, serial+12)
	i, j := 0, 1
	scores[0], scores[1] = 3, 7
	si, sj := 0, 0
	for len(scores) < serial+10 {
		si, sj = scores[i], scores[j]
		scores = append(scores, newScores(si, sj)...)
		i, j = i+1+si, j+1+sj
		i, j = i%len(scores), j%len(scores)
	}
	return scoresToString(scores[serial : serial+10])
}

// https://stackoverflow.com/a/15312097
func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func stringToSlice(s string) []int {
	i := make([]int, 0, len(s))
	for _, r := range s {
		t, _ := strconv.Atoi(string(r))
		i = append(i, t)
	}
	return i
}

func main2(serial string) int {
	comp := stringToSlice(serial)
	scores := make([]int, 2)
	i, j := 0, 1
	scores[0], scores[1] = 3, 7
	si, sj := 0, 0
	compLen := len(comp)
	for {
		si, sj = scores[i], scores[j]
		for _, score := range newScores(si, sj) {
			scores = append(scores, score)
			slice := len(scores) - compLen
			if slice < 0 {
				slice = 0
			}
			if testEq(scores[slice:], comp) {
				return len(scores) - compLen
			}
		}
		i, j = i+1+si, j+1+sj
		i, j = i%len(scores), j%len(scores)
	}
}

func main() {
	flag.Parse()
	fmt.Println("Starting main1")
	fmt.Println(main1(9), "5158916779")
	fmt.Println(main1(5), "0124515891")
	fmt.Println(main1(18), "9251071085")
	fmt.Println(main1(2018), "5941429882")
	fmt.Println(main1(360781))
	fmt.Println("Starting main2")
	fmt.Println(main2("51589"), 9)
	fmt.Println(main2("01245"), 5)
	fmt.Println(main2("92510"), 18)
	fmt.Println(main2("59414"), 2018)
	fmt.Println(main2("360781"))
	fmt.Println("Done!")
}
