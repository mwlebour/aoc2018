package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// FileToList takes a path to a file and returns a list
// of strings, one per line
func FileToList(fn string) []string {

	lst := make([]string, 0)

	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		lst = append(lst, t)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lst
}

// ListToIntList takes a list of strings
// and converts them all to integers returning
// an equivalently sized list of integers
func ListToIntList(lst []string) []int {

	newLst := make([]int, len(lst))

	for i, v := range lst {
		t, _ := strconv.Atoi(v)
		newLst[i] = t

	}
	return newLst
}
