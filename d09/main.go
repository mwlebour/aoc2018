package main

import (
	"flag"
	"fmt"
)

func nextPlayer(c int, n int) int {
	if c+1 == n {
		return 0
	}
	return c + 1
}

func place(board []int, index int, value int) ([]int, int) {
	index = clock(index, len(board), 2)
	if index == 0 {
		board = append(board, value)
		index = len(board) - 1
	} else {
		board = append(board[:index], append([]int{value}, board[index:]...)...)
	}
	return board, index
}

func counter(i int, size int, shift int) int {
	if i-shift < 0 {
		return i - shift + size
	}
	return i - shift
}

func clock(i int, size int, shift int) int {
	if i+shift >= size {
		return i + shift - size
	}
	return i + shift
}

func max(l []int) int {
	m := 0
	for i, v := range l {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func main1(nPlayers int, nMarbles int) {
	board := make([]int, nMarbles, nMarbles)
	currentPlayer := 0
	scores := make([]int, nPlayers, nPlayers)
	index := 0
	for i := 1; i <= nMarbles; i++ {
		if i%23 == 0 {
			scores[currentPlayer] += i
			index = counter(index, len(board), 7)
			scores[currentPlayer] += board[index]
			board = append(board[:index], board[(index+1):]...)
		} else {
			board, index = place(board, index, i)
		}
		currentPlayer = nextPlayer(currentPlayer, nPlayers)
		if i%10000 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(max(scores))
}

// Marble marbles
type Marble struct {
	value int
	prev  *Marble
	next  *Marble
}

// Board boards
type Board struct {
	current *Marble
}

func (b *Board) shift(n int) {
	if n > 0 {
		b.current = b.current.next
		b.shift(n - 1)
	}
	if n < 0 {
		b.current = b.current.prev
		b.shift(n + 1)
	}
}

func (b *Board) remove() {
	b.current.prev.next = b.current.next
	b.current.next.prev = b.current.prev
	b.current = b.current.next
	// TODO: does this leak? who manages memory aaaaaaah?
}

func (b *Board) place(value int) {
	m := &Marble{value, b.current.prev, b.current}
	b.current.prev.next = m
	b.current.prev = m
	b.current = m
}

func main2(nPlayers int, nMarbles int) {
	m0 := Marble{0, nil, nil}
	m0.prev = &m0
	m0.next = &m0
	board := &Board{&m0}
	scores := make([]int, nPlayers, nPlayers)
	for i := 1; i <= nMarbles; i++ {
		if i%23 == 0 {
			scores[i%nPlayers] += i
			board.shift(-7)
			scores[i%nPlayers] += board.current.value
			board.remove()
		} else {
			board.shift(2)
			board.place(i)
		}
	}
	fmt.Println(max(scores))
}

func main() {
	flag.Parse()
	fmt.Println("Starting main1")
	main1(9, 25)
	main1(10, 1618)
	main1(13, 7999)
	main1(17, 1104)
	// main1(471, 72026)
	// main1(471, 7202600)
	fmt.Println("Starting main2")
	main2(9, 25)
	main2(10, 1618)
	main2(13, 7999)
	main2(17, 1104)
	main2(471, 72026)
	main2(471, 7202600)
	fmt.Println("Done!")
}
