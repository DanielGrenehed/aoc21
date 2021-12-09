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

type QNode struct {
	next *QNode
	data *Position
}

type Queuer interface {
	IsEmpty() bool
	Enqueue(*Position)
	Dequeue() *Position
	Size() int
}

type Queue struct {
	start *QNode
}

func (q *Queue) IsEmpty() bool {
	return q.start == nil
}

func (q *Queue) Size() int {
	if q.start == nil {
		return 0
	}
	res := 0
	for c := q.start; c != nil; c = c.next {
		res++
	}
	return res
}

func (q *Queue) Enqueue(p *Position) {
	if q.IsEmpty() {
		q.start = &QNode{nil, p}
	} else {
		c := q.start
		for ; c.next != nil; c = c.next {
		}
		c.next = &QNode{nil, p}
	}
}

func (q *Queue) Dequeue() *Position {
	if q.IsEmpty() {
		return nil
	} else {
		res := q.start.data
		q.start = q.start.next
		return res
	}
}

func CreateQueue() *Queue {
	return &Queue{nil}
}

func contains(arr []Position, p Position) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i].x == p.x && arr[i].y == p.y {
			return true
		}
	}
	return false
}

func getAdjecent(p *Position, maxx int, maxy int) []Position {
	var out []Position
	if p.x-1 >= 0 {
		out = append(out, Position{p.x - 1, p.y})
	}
	if p.y-1 >= 0 {
		out = append(out, Position{p.x, p.y - 1})
	}
	if p.x+1 < maxx {
		out = append(out, Position{p.x + 1, p.y})
	}
	if p.y+1 < maxy {
		out = append(out, Position{p.x, p.y + 1})
	}
	return out
}

func GetBasin(p Position, hmap []string) int {
	if hmap[p.y][p.x] == '9' || hmap[p.y][p.x] == ' ' {
		return 0
	}
	var covered []Position
	covered = append(covered, p)
	line := []rune(hmap[p.y])
	line[p.x] = ' '
	hmap[p.y] = string(line)
	q := CreateQueue()
	q.Enqueue(&p)
	maxx := len(hmap[0])
	maxy := len(hmap)

	for !q.IsEmpty() {
		pos := q.Dequeue()

		adjecent := getAdjecent(pos, maxx, maxy)
		for i := 0; i < len(adjecent); i++ {
			apos := adjecent[i]
			if hmap[apos.y][apos.x] != '9' && hmap[apos.y][apos.x] != ' ' && !contains(covered, apos) {
				covered = append(covered, apos)
				q.Enqueue(&apos)
				line := []rune(hmap[apos.y])
				line[apos.x] = ' '
				hmap[apos.y] = string(line)
			}
		}
	}

	return len(covered)
}

func printMap(hmap []string) {
	for y := 0; y < len(hmap); y++ {
		line := hmap[y]
		var out string
		if y < 10 {
			out = " "
		}
		out += strconv.Itoa(y) + ": "
		for x := 0; x < len(line); x++ {
			char := line[x]

			if char == '9' {
				out += "\033[31m" + string(char) + "\033[0m"
			} else {
				out += string(char)
			}

		}
		fmt.Println(out)
	}
}

type Top3er interface {
	Push(int)
	Multiply() int
	SmallestPos() int
}

type Top3 struct {
	values [3]int
}

func (t *Top3) Push(d int) {
	i := t.SmallestPos()
	if d > t.values[i] {
		t.values[i] = d
	}
}

func (t *Top3) SmallestPos() int {
	pos := 0
	for i := 0; i < 3; i++ {
		if t.values[i] < t.values[pos] {
			pos = i
		}
	}
	return pos
}

func (t *Top3) Multiply() int {
	sum := 1
	for i := 0; i < 3; i++ {
		sum *= t.values[i]
	}
	return sum
}

func CreateTop() Top3 {
	t := Top3{}
	t.values[0] = 0
	t.values[1] = 1
	t.values[2] = 2
	return t
}

// find lowpoint in line, compare to previous line
// if lowpoint is not succeded in next line, add to sum

func sumLowpoints(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var heigtmap []string

	for scanner.Scan() {
		line := scanner.Text()
		heigtmap = append(heigtmap, line)
	}

	t := CreateTop()
	for y := 0; y < len(heigtmap); y++ {
		for x := 0; x < len(heigtmap[0]); x++ {
			t.Push(GetBasin(Position{x, y}, heigtmap))
		}
	}
	fmt.Println("MapSize:", len(heigtmap[0]), len(heigtmap))
	printMap(heigtmap)
	// find basins for items in list on the heigtmap

	return t.Multiply()
}

func main() {

	fmt.Println(sumLowpoints(os.Args[1]))
}
