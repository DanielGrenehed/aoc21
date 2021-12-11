package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func byteToInt(b byte) int {
	val, err := strconv.Atoi(string(b))
	if err != nil {
		fmt.Println("Error converting", b, "to int")
		return 10
	}
	return val
}

type Lowpointer interface {
	IsValid(string) bool
}

type Lowpoint struct {
	x   int
	y   int
	val byte
}

func (l *Lowpoint) IsValid(newline string) bool {
	if len(newline) == 0 || len(newline) < l.x {
		return true
	}
	val := newline[l.x]
	if val >= l.val {
		return true
	}
	return false
}

type Node struct {
	next *Node
	data *Lowpoint
}

type LowpointLister interface {
	Add(*Lowpoint)
	AddList(*LowpointList)
	Pop() *Lowpoint
	IsEmpty() bool
	Filter(string)
	Delete(*Lowpoint)
	Size() int
}

type LowpointList struct {
	head *Node
}

func (l *LowpointList) Add(data *Lowpoint) {
	if l.head == nil {
		l.head = &Node{nil, data}
	} else {
		c := l.head
		for ; c.next != nil; c = c.next {
		}
		c.next = &Node{nil, data}
	}
}

func (l *LowpointList) AddList(sec *LowpointList) {
	if l.head == nil {
		l.head = sec.head
	} else {
		c := l.head
		for ; c.next != nil; c = c.next {
		}
		c.next = sec.head
	}
}

func (l *LowpointList) Pop() *Lowpoint {
	if l.IsEmpty() {
		return nil

	}
	res := l.head.data
	l.head = l.head.next
	return res
}

func (l *LowpointList) IsEmpty() bool {
	return l.head == nil
}

func (l *LowpointList) Delete(d *Lowpoint) {
	if l.head == nil {
		return
	}
	if l.head.data == d {
		l.head = l.head.next
		return
	}
	c := l.head
	prev := l.head
	for ; c != nil; c = c.next {
		if c.data == d {
			prev.next = c.next
		}
	}
}

func (l *LowpointList) Filter(line string) {
	c := l.head
	for ; c != nil; c = c.next {
		if !c.data.IsValid(line) {
			l.Delete(c.data)
		}
	}
}

func (l *LowpointList) Size() int {
	if l.head == nil {
		return 0
	}
	sum := 1
	for c := l.head; c != nil; c = c.next {
		sum++
	}
	return sum
}

func CreateLowList() *LowpointList {
	return &LowpointList{nil}
}

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

func GetBasin(l Lowpoint, hmap []string) []Position {
	var covered []Position
	start := Position{l.x, l.y}
	covered = append(covered, start)
	q := CreateQueue()
	q.Enqueue(&start)
	maxx := len(hmap[0])
	maxy := len(hmap)

	for !q.IsEmpty() {
		pos := q.Dequeue()
		/* if hmap[pos.y][pos.x] != '9' {
			if !contains(covered, *pos) {
				covered = append(covered, *pos)
				adjecent := getAdjecent(pos, maxx, maxy)
				for i := 0; i < len(adjecent); i++ {
					q.Enqueue(&adjecent[i])
				}
				fmt.Print("QueueSize:", q.Size())
			}
		} */
		adjecent := getAdjecent(pos, maxx, maxy)
		for i := 0; i < len(adjecent); i++ {
			apos := adjecent[i]
			if hmap[apos.y][apos.x] != '9' && !contains(covered, apos) {
				//fmt.Println("Adding position {", apos.x, apos.y, "}", string(hmap[apos.y][apos.x]))
				covered = append(covered, apos)
				q.Enqueue(&apos)
			}
		}
	}

	return covered
}

func printMap(hmap []string, covered []Position) {
	for y := 0; y < len(hmap); y++ {
		line := hmap[y]
		var out string
		if y < 10 {
			out = " "
		}
		out += strconv.Itoa(y) + ": "
		for x := 0; x < len(line); x++ {
			char := line[x]
			if contains(covered, Position{x, y}) {
				out += "\033[32m" + string(char) + "\033[0m"
			} else {
				if char == '9' {
					out += "\033[31m" + string(char) + "\033[0m"
				} else {
					out += string(char)
				}

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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := CreateLowList()

	var heigtmap []string
	prevl := ""
	tlist := CreateLowList()
	vert := 0
	for scanner.Scan() {
		line := scanner.Text()
		heigtmap = append(heigtmap, line)
		// test and empty lowpoint list
		tlist.Filter(line)
		list.AddList(tlist)
		tlist = CreateLowList()
		// find possible new lowpoints
		left := ":"[0]
		for i := 0; i < len(line); i++ {
			center := line[i]
			right := ":"[0]
			if i != len(line)-1 {
				right = line[i+1]
			}
			if center < left && center < right {
				low := &Lowpoint{i, vert, center}
				if low.IsValid(prevl) {
					tlist.Add(low)
				}
			}
			left = center
		}
		prevl = line
		vert++
	}
	tlist.Filter("")
	list.AddList(tlist)
	fmt.Println("Lowpoints found:", list.Size())
	t := CreateTop()
	var covered []Position

	for !list.IsEmpty() {
		low := list.Pop()
		line := []rune(heigtmap[low.y])
		line[low.x] = ' '
		heigtmap[low.y] = string(line)
		cov := GetBasin(*low, heigtmap)
		covered = append(covered, cov...)
		t.Push(len(cov))
	}
	fmt.Println("MapSize:", len(heigtmap[0]), len(heigtmap))
	printMap(heigtmap, covered)
	// find basins for items in list on the heigtmap

	return t.Multiply()
}

func main() {

	fmt.Println(sumLowpoints(os.Args[1]))
}
