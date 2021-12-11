package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func scoreClosing(c byte) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
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

func findClosing(line string) int {
	i := 0
	opening := &Stack{nil}
	for ; i < len(line); i++ {
		if Opening(line[i]) > 0 {
			opening.Push(i)
		} else if Closing(line[i]) > 0 {
			//fmt.Print("Stack size:", opening.Size())
			start := opening.Pop()
			if Opening(line[start]) != Closing(line[i]) {
				fmt.Println("Expected", string(line[start]), ", but found", string(line[i]), "in", line)
				return i
			}
		} else {
			fmt.Println("Unknown symbol:", string(line[i]), "in", line)
		}
	}
	return -1
}

func scoreLine(line string) int {
	// find corrupted lines
	// "Stop at the first incorrect closing character on each corrupted line."
	closing := findClosing(line)
	if closing > 0 {
		return scoreClosing(line[closing])
	}
	return 0 // return 1 on incomplete or complete lines
}

func scoreFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += scoreLine(scanner.Text())
	}
	return sum
}

func main() {
	fmt.Println(scoreFile(os.Args[1]))
}
