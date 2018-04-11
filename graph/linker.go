package graph

import (
	"fmt"
	"log"
	"sync"
)

type linkerGraph struct {
	mu sync.RWMutex

	graphSize int
	nodeCount int

	node   map[string]Node
	inEdge map[string][]string
	outD   map[string]int

	queue    chan Node
	finishCh chan struct{}
}

func newLinkerGraph(num int) Graph {
	return &linkerGraph{
		graphSize: num,
		node:      make(map[string]Node, num),
		inEdge:    make(map[string][]string),
		outD:      make(map[string]int),
		queue:     make(chan Node, num),
		finishCh:  make(chan struct{}),
	}
}

func (g *linkerGraph) AddNode(n Node) error {
	id := n.ID()

	g.mu.RLock()
	if g.nodeCount >= g.graphSize {
		g.mu.RUnlock()
		return fmt.Errorf("%s can not add to node, for graph is full", id)
	}
	if g.nodeIsExist(id) {
		g.mu.RUnlock()
		return fmt.Errorf("%s has already existed.", id)
	}
	g.mu.RUnlock()

	g.mu.Lock()
	g.node[id] = n
	g.nodeCount++
	g.mu.Unlock()
	return nil
}

func (g *linkerGraph) AddEdge(src, dst string) error {
	g.mu.RLock()
	if !g.nodeIsExist(src) {
		g.mu.RUnlock()
		return fmt.Errorf("%s is not exist", src)
	}
	if !g.nodeIsExist(dst) {
		g.mu.RUnlock()
		return fmt.Errorf("%s is not exist", dst)
	}

	g.mu.RUnlock()
	g.mu.Lock()
	if g.nodeHasInDegreeEdge(dst) {
		g.inEdge[dst] = append(g.inEdge[dst], src)
	} else {
		g.inEdge[dst] = []string{src}
	}
	g.mu.Unlock()
	return nil
}

func (g *linkerGraph) nodeIsExist(id string) bool {
	_, ok := g.node[id]
	return ok
}

func (g *linkerGraph) nodeHasInDegreeEdge(id string) bool {
	_, ok := g.inEdge[id]
	return ok
}

func (g *linkerGraph) Edges() interface{} {
	return g.inEdge
}

func (g *linkerGraph) OutD() interface{} {
	return g.outD
}

func (g *linkerGraph) DFS() {
	g.calculateOutD()
	g.dfs("")
}

func (g *linkerGraph) dfs(root string) {
	subnodes := g.findOutDzeroNodes(root)
	if len(subnodes) == 0 {
		return
	}

	for _, id := range subnodes {
		g.node[id].Run()
		g.regreshOutD(id, false)
		g.dfs(id)
	}
}

func (g *linkerGraph) BFS() {
	g.calculateOutD()
	g.bfs()
}

func (g *linkerGraph) bfs() {
	subnodes := g.findOutDzeroNodes("")

	for _, id := range subnodes {
		g.queue <- g.node[id]
	}
Loop:
	for {
		select {
		case <-g.finishCh:
			// log.Println("graph already finished")
			break Loop
		case n := <-g.queue:
			n.Run()
			g.regreshOutD(n.ID(), true)
		}
	}
}

func (g *linkerGraph) calculateOutD() {
	for id, edges := range g.inEdge {
		if _, ok := g.outD[id]; !ok {
			g.outD[id] = 0
		}
		for _, out := range edges {
			g.outD[out] += 1
		}
	}
}

func (g *linkerGraph) findOutDzeroNodes(root string) []string {
	nodes := []string{}
	g.mu.RLock()
	if root == "" {
		for id, out := range g.outD {
			if out == 0 {
				nodes = append(nodes, id)
			}
		}
		g.mu.RUnlock()
		return nodes
	}

	for _, id := range g.inEdge[root] {
		if g.outD[id] == 0 {
			nodes = append(nodes, id)
		}
	}
	g.mu.RUnlock()
	return nodes
}

func (g *linkerGraph) regreshOutD(id string, bfs bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, i := range g.inEdge[id] {
		if _, ok := g.outD[i]; ok {
			g.outD[i] -= 1
			if bfs && g.outD[i] == 0 {
				g.queue <- g.node[i]
			}
		}
	}
	delete(g.outD, id)
	if bfs && len(g.outD) == 0 {
		g.finish()
	}
}

func (g *linkerGraph) finish() {
	select {
	case <-g.finishCh:
		log.Printf("finishCh already closed.")
	default:
		close(g.finishCh)
	}
}
