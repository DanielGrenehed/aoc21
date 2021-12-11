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
	pos int
	val int
}

func (l *Lowpoint) IsValid(newline string) bool {
	if len(newline) == 0 || len(newline) < l.pos {
		return true
	}
	val := byteToInt(newline[l.pos])
	if val >= l.val {
		return true
	}
	return false
}

type Node struct {
	next *Node
	data *Lowpoint
}

type Lister interface {
	Add(*Lowpoint)
	Pop() *Lowpoint
	IsEmpty() bool
	Sum(string) int
}

type List struct {
	head *Node
}

func (l *List) Add(data *Lowpoint) {
	if l.head == nil {
		l.head = &Node{nil, data}
	} else {
		c := l.head
		for ; c.next != nil; c = c.next {
		}
		c.next = &Node{nil, data}
	}
}

func (l *List) Pop() *Lowpoint {
	if l.IsEmpty() {
		return nil

	}
	res := l.head.data
	l.head = l.head.next
	return res
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) Sum(line string) int {
	sum := 0
	for !l.IsEmpty() {
		low := l.Pop()
		if low == nil {
			continue
		}

		if low.IsValid(line) {
			sum += low.val + 1
		}

	}

	return sum
}

func CreateList() *List {
	return &List{nil}
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
	list := CreateList()

	prevl := ""
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		// test and empty lowpoint list
		sum += list.Sum(line)

		// find possible new lowpoints
		prevv := 10
		for i := 0; i < len(line); i++ {
			c := byteToInt(line[i])
			n := 10
			if i != len(line)-1 {
				n = byteToInt(line[i+1])
			}
			if c < prevv && c < n {
				low := &Lowpoint{i, c}
				if low.IsValid(prevl) {
					list.Add(low)
				}

			}
			prevv = c

		}

		prevl = line

	}

	sum += list.Sum("")
	return sum
}

func main() {

	fmt.Println(sumLowpoints(os.Args[1]))
}
