package main

import (
	"flag"
	"fmt"

	"github.com/mwlebour/aoc2018/util"
)

// Edge represents a line in the directed graph
type Edge struct {
	begin rune
	end   rune
}

// Edges are a list of Edge
type Edges []Edge

// Vertex is a vertex in the directed graph
type Vertex struct {
	id       rune
	incoming Edges
	outgoing Edges
	endTime  int
}

func (v Vertex) String() string {
	// return fmt.Sprintf("[ %v ] -> %s -> [ %v ]", v.incoming, string(v.id), v.outgoing)
	return fmt.Sprintf("%s (%d)", string(v.id), len(v.incoming))
}

// Vertices are a list of Vertex
type Vertices []Vertex

// VertexMap is a map from rune to vertex
type VertexMap map[rune]Vertex

func findAll(vertices VertexMap) Vertices {

	candidates := make(Vertices, 0)
	for _, v := range vertices {
		if len(v.incoming) == 0 {
			candidates = append(candidates, v)
		}
	}
	return candidates
}

func findFirst(vertices VertexMap) *Vertex {
	return pickCandidate(findAll(vertices))
}

func pickCandidate(candidates Vertices) *Vertex {
	m := rune(9999999)
	var next Vertex
	for _, c := range candidates {
		if c.id < m {
			next = c
			m = c.id
		}
	}
	return &next
}

func removeIncoming(c *Vertex, n *Vertex) *Vertex {
	edges := c.incoming
	for i, edge := range edges {
		if edge.begin == n.id {
			edges[i] = edges[len(edges)-1]
			return &Vertex{c.id, edges[:len(edges)-1], c.outgoing, c.endTime}
		}
	}
	return nil
}

func main1(edges Edges, vertexMap VertexMap) {
	var next *Vertex
	for len(vertexMap) > 0 {
		next = findFirst(vertexMap)
		fmt.Print(string(next.id))
		for _, v := range next.outgoing {
			t := vertexMap[v.end]
			vertexMap[v.end] = *removeIncoming(&t, next)
		}
		delete(vertexMap, next.id)
	}
	fmt.Println()
}

// Workers gonna work work work
type Workers map[rune]Vertex

func (workers Workers) String() string {
	s := ""
	for i, v := range workers {
		s = fmt.Sprintf("%s|%d %d", s, i, v.endTime)
	}
	return s
}

// Check needs to disable golang warnings
func (workers Workers) Check(currentTime int) Vertices {
	vertices := make(Vertices, 0)
	for _, v := range workers {
		if currentTime > v.endTime {
			vertices = append(vertices, v)
		}
	}
	return vertices
}

func main2(edges Edges, vertexMap VertexMap) {
	workers := make(Workers)
	baseTime := 60
	maxWorkers := 5
	time := 0
	done := make([]rune, 0, len(vertexMap))
	for len(vertexMap) > 0 || len(workers) > 0 {
		finishedVertices := workers.Check(time)
		for _, v := range finishedVertices {
			for _, e := range v.outgoing {
				t := vertexMap[e.end]
				vertexMap[e.end] = *removeIncoming(&t, &v)
			}
			done = append(done, v.id)
			delete(workers, v.id)
		}
		possibleNexts := findAll(vertexMap)
		for len(possibleNexts) > 0 && len(workers) < maxWorkers {
			nextID := possibleNexts[len(possibleNexts)-1].id
			v := vertexMap[nextID]
			delete(vertexMap, nextID)
			v.endTime = time + baseTime + int(nextID) - int('A')
			workers[nextID] = v
			possibleNexts = possibleNexts[:len(possibleNexts)-1]
		}
		fmt.Printf("Starting %d with %d workers, %d finished, %d done\n",
			time, len(workers), len(finishedVertices), len(done))
		time++
	}
}

var full bool

func init() {
	flag.BoolVar(&full, "full", false, "run full solution")
}

func buildEdges(lines []string) Edges {
	edges := make(Edges, 0, len(lines))
	for _, line := range lines {
		var b, e rune
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &b, &e)
		edges = append(edges, Edge{b, e})
	}
	return edges
}

func buildVertices(edges Edges) VertexMap {
	vertexMap := make(VertexMap)
	for _, e := range edges {
		if _, ok := vertexMap[e.begin]; !ok {
			vertexMap[e.begin] = Vertex{e.begin, make(Edges, 0), make(Edges, 0), 0}
		}
		if _, ok := vertexMap[e.end]; !ok {
			vertexMap[e.end] = Vertex{e.end, make(Edges, 0), make(Edges, 0), 0}
		}
		t := vertexMap[e.begin]
		vertexMap[e.begin] = Vertex{t.id, t.incoming, append(t.outgoing, e), 0}
		t = vertexMap[e.end]
		vertexMap[e.end] = Vertex{t.id, append(t.incoming, e), t.outgoing, 0}
	}
	return vertexMap
}

func main() {
	flag.Parse()
	input := util.FileToList("unit.out")
	if full {
		input = util.FileToList("input.out")
	}
	edges := buildEdges(input)
	vertices := buildVertices(edges)
	fmt.Println("Starting main1")
	main1(edges, vertices)
	edges = buildEdges(input)
	vertices = buildVertices(edges)
	fmt.Println("Starting main2")
	main2(edges, vertices)
	fmt.Println("Done!")
}
