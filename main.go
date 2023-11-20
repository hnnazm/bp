package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

var Infinity = math.MaxInt

type Graph struct {
	size  int
	nodes []*Node
	edges []*Edge
}

func NewGraph(n int) *Graph {
	return &Graph{
		size:  n,
		nodes: make([]*Node, 0),
		edges: make([]*Edge, 0),
	}
}

func (g *Graph) addNode(name string) (*Node, error) {
	node := &Node{name: name, edges: make(map[string]*Edge)}

	for _, v := range g.nodes {
		if v != nil && v.name == node.name {
			return nil, fmt.Errorf("Node already exists")
		}
	}

	g.nodes = append(g.nodes, node)

	return node, nil
}

func (g *Graph) link(name string, n1 *Node, n2 *Node, duration int) error {
	for _, v := range g.edges {
		if v != nil && v.name == name {
			return fmt.Errorf("Edge %s <-> %s already exists.", n1.name, n2.name)
		}
	}

	n1.edges[name] = &Edge{
		name:     name,
		from:     n1,
		to:       n2,
		duration: duration,
	}

	n2.edges[name] = &Edge{
		name:     name,
		from:     n2,
		to:       n1,
		duration: duration,
	}

	g.edges = append(g.edges, &Edge{
		name:     name,
		from:     n1,
		to:       n2,
		duration: duration,
	})

	return nil
}

type Payload struct {
	weight        int
	currentWeigth int
	origin        *Node
	destination   *Node
}

func NewPayload(name string, origin, destination *Node, weight int) *Payload {
	return &Payload{
		origin:        origin,
		destination:   destination,
		weight:        weight,
		currentWeigth: weight,
	}
}

type Edge struct {
	name     string
	from     *Node
	to       *Node
	duration int
}

type Node struct {
	name  string
	edges map[string]*Edge
}

func NewNode(name string) *Node {
	return &Node{
		name:  name,
		edges: make(map[string]*Edge),
	}
}

func (n *Node) Link(name string, node *Node, duration int) error {
	if n.edges[name] != nil {
		return fmt.Errorf("Edge %s <-> %s already exists.", n.name, node.name)
	}

	n.edges[name] = &Edge{
		name:     name,
		to:       node,
		duration: duration,
	}

	node.edges[name] = &Edge{
		name:     name,
		to:       n,
		duration: duration,
	}

	return nil
}

type Train struct {
	name     string
	capacity int
	current  *Node
	path     map[*Node]int
	parent   map[*Node]*Node
	visited  []*Node
}

func NewTrain(name string, capacity int, origin *Node) *Train {
	return &Train{
		name:     name,
		capacity: capacity,
		current:  origin,
		path:     make(map[*Node]int, 0),
		parent:   make(map[*Node]*Node, 0),
		visited:  make([]*Node, 0),
	}
}

func (t *Train) pick(payload *Payload) {
	for len(t.visited) != len(t.path) {
		nextNode := t.gsu()

		t.visited = append(t.visited, nextNode)
		t.current = nextNode

		for _, edge := range t.current.edges {
			if _, ok := t.path[edge.to]; !ok {
				t.path[edge.to] = edge.duration
			}

			for _, visited := range t.visited[1:] {
				if edge.from == visited && edge.to == visited {

					t.parent[t.current] = edge.to
				}
			}
		}
	}

	for path, duration := range t.path {
		log.Printf("%v: %s", path.name, (time.Duration(duration) * time.Minute))
	}
}

func (t *Train) deliver(payload *Payload) {}
func (t *Train) reset(payload *Payload)   {}

func contains(s []*Node, e *Node) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (t *Train) gsu() *Node {
	var (
		nextNode         *Node
		shortestDuration = Infinity
	)

	for node, duration := range t.path {
		if !contains(t.visited, node) {
			if duration < shortestDuration {
				shortestDuration = duration
				nextNode = node
			}
		}
	}

	return nextNode
}

func (g *Graph) execute(t *Train, p *Payload) error {
	if t.capacity <= p.weight {
		return fmt.Errorf("Capacity exceeded!")
	}

	// Step 0: Setup
	t.path[t.current] = 0
	t.visited = append(t.visited, t.current)
	t.parent[t.current] = nil

	for _, edge := range t.current.edges {
		t.path[edge.to] = edge.duration
	}

	// Step 1: Pickup
	t.pick(p)

	// Step 2: Update payload
	p.currentWeigth = t.capacity - p.weight

	// Step 3: Deliver
	t.deliver(p)

	return nil
}

func main() {
	graph := NewGraph(4)

	n1, err := graph.addNode("A")

	if err != nil {
		log.Fatalf(err.Error())
	}

	n2, err := graph.addNode("B")

	if err != nil {
		log.Fatalf(err.Error())
	}

	n3, err := graph.addNode("C")

	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := graph.link("E1", n1, n2, 30); err != nil {
		log.Fatalf(err.Error())
	}

	if err := graph.link("E2", n2, n3, 10); err != nil {
		log.Fatalf(err.Error())
	}

	p1 := NewPayload("K1", n1, n3, 5)

	q1 := NewTrain("Q1", 6, n2)

	if err := graph.execute(q1, p1); err != nil {
		log.Fatalf(err.Error())
	}
}
