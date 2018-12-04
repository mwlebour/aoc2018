package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/mwlebour/aoc2018/util"
)

// SleepMap is
type SleepMap map[int]int

func (m SleepMap) addSleep(s int, e int) {
	for i := s; i < e; i++ {
		m[i]++
	}
}

func (m SleepMap) total() (s int) {
	for _, v := range m {
		s += v
	}
	return
}

func (m SleepMap) bestMinute() (minute int, best int) {
	for i, v := range m {
		if v > best {
			minute, best = i, v
		}
	}
	return
}

func parseGuard(s string) (i int) {
	var t, u string
	fmt.Sscanf(s, "%s %s Guard #%d begins shift", &t, &u, &i)
	return
}

func parseDate(s string) int {
	t, _ := time.Parse("2006-01-02 15:04", s[1:17])
	return t.Minute()
}

func newGuard(s string) bool {
	return strings.Contains(s, "begins shift")
}

const (
	start = iota
	check
	waking
)

func loadGuards(schedule []string) map[int]SleepMap {
	sort.Strings(schedule)
	row := 0
	state := start
	guards := make(map[int]SleepMap)
	currentGuard := 0
	for row < len(schedule) {
		switch state {
		case start:
			currentGuard = parseGuard(schedule[row])
			if _, ok := guards[currentGuard]; !ok {
				guards[currentGuard] = SleepMap{}
			}
			state = check
			row++
		case waking:
			startMinute := parseDate(schedule[row])
			row++
			endMinute := parseDate(schedule[row])
			row++
			guards[currentGuard].addSleep(startMinute, endMinute)
			state = check
		case check:
			state = waking
			if newGuard(schedule[row]) {
				state = start
			}
		}
	}
	return guards
}

func main1(schedule []string) {
	guards := loadGuards(schedule)
	maxSleep, sleepiest := 0, 0
	for i, g := range guards {
		t := g.total()
		if t > maxSleep {
			maxSleep, sleepiest = t, i
		}
	}
	minute, best := guards[sleepiest].bestMinute()
	fmt.Println(maxSleep, sleepiest, minute, best, sleepiest*minute)
}

func main2(schedule []string) {
	guards := loadGuards(schedule)
	maxSleep, sleepiest := 0, 0
	for i, g := range guards {
		_, b := g.bestMinute()
		if b > maxSleep {
			maxSleep, sleepiest = b, i
		}
	}
	minute, best := guards[sleepiest].bestMinute()
	fmt.Println(maxSleep, sleepiest, minute, best, sleepiest*minute)
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
	main1(input)
	main2(input)
}
