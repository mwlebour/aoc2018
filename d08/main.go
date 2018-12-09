package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mwlebour/aoc2018/util"
)

// Tree is just an index into the array
type Tree int

// Node is a node
type Node struct {
	index     int
	nChildren int
	children  []interface{}
	nMetadata int
	metadata  []int
}

// Puzzle is the input puzzle array of ints
type Puzzle []int

func puzzleFromString(s string) Puzzle {
	puzzle := make(Puzzle, 0)
	for _, t := range strings.Fields(s) {
		i, _ := strconv.Atoi(t)
		puzzle = append(puzzle, i)
	}
	return puzzle
}

func buildTree(puzzle Puzzle) []Node {
	index := 0
	childrenToProcess := 1
	allNodes := make([]Node, 0, 10000)
	nodeStack := make([]*Node, 0)
	for len(nodeStack) > 0 || childrenToProcess > 0 {
		if childrenToProcess > 0 {
			node := Node{
				index,
				puzzle[index],
				make([]interface{}, 0),
				puzzle[index+1],
				make([]int, 0),
			}
			index += 2
			allNodes = append(allNodes, node)
			nodeStack = append(nodeStack, &allNodes[len(allNodes)-1])
			childrenToProcess = node.nChildren
		}
		if childrenToProcess == 0 {
			thisNode := nodeStack[len(nodeStack)-1]
			thisNode.metadata = puzzle[index : index+thisNode.nMetadata]
			nodeStack = nodeStack[:len(nodeStack)-1]
			if len(nodeStack) > 0 {
				nextNode := nodeStack[len(nodeStack)-1]
				nextNode.children = append(nextNode.children, *thisNode)
				childrenToProcess = nextNode.nChildren - len(nextNode.children)
			}
			index += thisNode.nMetadata
		}
	}
	return allNodes
}

func metaSum(n Node) int {
	s := 0
	for _, m := range n.metadata {
		s += m
	}
	return s
}

func main1(puzzle Puzzle) {
	nodes := buildTree(puzzle)
	s := 0
	for _, n := range nodes {
		s += metaSum(n)
	}
	fmt.Println(s)
}

func nodeValue(n Node) int {
	if n.nChildren == 0 {
		return metaSum(n)
	}
	s := 0
	for _, m := range n.metadata {
		if m <= len(n.children) {
			s += nodeValue(n.children[m-1].(Node))
		}
	}
	return s
}

func main2(puzzle Puzzle) {
	nodes := buildTree(puzzle)
	fmt.Println(nodeValue(nodes[0]))
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
	fmt.Println("Starting main1")
	tree := puzzleFromString(input[0])
	main1(tree)
	fmt.Println("Starting main2")
	main2(tree)
	fmt.Println("Done!")
}
