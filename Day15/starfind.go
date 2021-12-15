package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Position struct {
	x int
	y int
}

type PQNode struct {
	weight int
	value  Position
}

type PriorityQueue struct {
	queue []PQNode
}

func (pq *PriorityQueue) Push(w int, d Position) {
	pq.queue = append(pq.queue, PQNode{w, d})
}

func (pq *PriorityQueue) Pop() PQNode {
	least := pq.queue[0]
	index := 0
	for i, e := range pq.queue {
		if e.weight < least.weight {
			least = e
			index = i
		}
	}
	pq.queue = append(pq.queue[:index], pq.queue[index+1:]...)

	return least
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.queue) == 0
}

func (pq *PriorityQueue) Contains(p Position) bool {
	for _, e := range pq.queue {
		if p.Equal(&e.value) {
			return true
		}
	}
	return false
}

func (pq *PriorityQueue) Size() int {
	return len(pq.queue)
}

type Score struct {
	g int
	f int
}

// path

func (t *Position) heuristic(p Position) int {
	return ((t.x - p.x) + (t.y - p.y))
}

func (p *Position) getV(field []string) int {
	v, e := strconv.Atoi(string(field[p.y][p.x]))
	if e == nil {
		return v
	}
	fmt.Println("Could not convert to integer")
	return 0
}

func (p *Position) Equal(p2 *Position) bool {
	if p.x == p2.x && p.y == p2.y {
		return true
	}
	return false
}

func AStar(field []string, target Position) []Position {
	// set points - score best path to point
	//PathCovered := make(map([

	came_from := make(map[Position]Position)
	start := Position{0, 0}
	open := PriorityQueue{}
	open.Push(0, start)

	Scores := make(map[Position]Score)
	Scores[start] = Score{0, target.heuristic(start)}

	for !open.IsEmpty() {
		current := open.Pop().value
		//fmt.Println(d.value)

		if target.Equal(&current) { // (reconstruct_path)
			c := target
			var path []Position
			for ; c != start; c = came_from[c] {
				path = append(path, c)
			}
			return append(path, start)
		}

		var neighbours []Position
		if current.x+1 < len(field[0]) {
			neighbours = append(neighbours, Position{current.x + 1, current.y})
		}
		if current.y+1 < len(field) {
			neighbours = append(neighbours, Position{current.x, current.y + 1})
		}
		if current.x-1 >= 0 {
			neighbours = append(neighbours, Position{current.x - 1, current.y})
		}
		if current.y-1 >= 0 {
			neighbours = append(neighbours, Position{current.x, current.y - 1})
		}

		for _, n := range neighbours {
			tgscore := Scores[current].g + n.getV(field)
			//fmt.Println(tgscore)
			//fmt.Println(open.Size())
			pscore, f := Scores[n]
			if !f || pscore.g > tgscore {
				came_from[n] = current
				fscore := tgscore + target.heuristic(n)
				Scores[n] = Score{tgscore, fscore}
				if !open.Contains(n) {
					open.Push(fscore, n)
				}
			}
		}

	}
	return []Position{}
}

func readMap(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var out []string
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}

func ScorePath(field []string, path []Position) int {

	for y := 0; y < len(field); y++ {
		out := ""
		for x := 0; x < len(field[y]); x++ {
			found := false
			for _, e := range path {
				if e.Equal(&Position{x, y}) {
					found = true
					break
				}
			}
			if found {
				out += "\033[31m" + string(field[y][x]) + "\033[0m"
			} else {
				out += string(field[y][x])
			}
		}
		fmt.Println(out)
	}
	start := Position{0, 0}
	score := 0
	for _, p := range path {
		if !p.Equal(&start) {
			score += p.getV(field)
			//fmt.Println(p, "Score", score)
		}
	}
	return score
}

func main() {
	field := readMap(os.Args[1])
	target := Position{len(field[0]) - 1, len(field) - 1}
	path := AStar(field, target)

	fmt.Println(ScorePath(field, path))
}
