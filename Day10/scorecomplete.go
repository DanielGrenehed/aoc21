package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func Closing(c byte) int {
	switch c {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}

func Opening(c byte) int {
	switch c {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	}
	return 0
}

type Node struct {
	next *Node
	data int
}

type Stacker interface {
	Push(int)
	Pop() int
	IsEmpty() bool
	Size() int
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(d int) {
	node := &Node{s.top, d}
	s.top = node
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	data := s.top.data
	s.top = s.top.next
	return data
}

func (s *Stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	count := 0
	for c := s.top; c != nil; c = c.next {
		count++
	}
	return count
}

func scoreFixedLine(line string) int {
	opening := &Stack{nil}
	for i := 0; i < len(line); i++ {
		if Opening(line[i]) > 0 {
			opening.Push(i)
		} else if Closing(line[i]) > 0 {
			start := opening.Pop()
			if Opening(line[start]) != Closing(line[i]) {
				return -1
			}
		} else {
			fmt.Println("Unknown symbol:", string(line[i]), "in", line)
		}
	}

	score := 0
	for !opening.IsEmpty() {
		score = (score * 5) + Opening(line[opening.Pop()])
	}

	return score
}

func scoreFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var list []int
	for scanner.Scan() {
		score := scoreFixedLine(scanner.Text())
		if score > 0 {
			list = append(list, score)
		}
	}
	sort.Ints(list)
	//fmt.Println(list)
	return list[(len(list) / 2)]
}

func main() {
	fmt.Println(scoreFile(os.Args[1]))
}
