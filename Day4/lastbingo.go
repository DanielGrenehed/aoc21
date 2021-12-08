package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	value  int
	marked bool
}

type Boarder interface {
	HasBingo(int, int)
	Mark(int)
	Print()
	Sum()
}

type Board struct {
	tiles [5 * 5]Tile
}

func (b *Board) HasBingo(x int, y int) bool {
	xmarked := 0
	ymarked := 0
	for i := 0; i < 5; i++ {
		if b.tiles[i+(y*5)].marked {
			xmarked++
		}
		if b.tiles[x+(i*5)].marked {
			ymarked++
		}
	}
	return xmarked == 5 || ymarked == 5
}

func (b *Board) Mark(number int) int {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b.tiles[x+(y*5)].value == number {
				b.tiles[x+(y*5)].marked = true
				if b.HasBingo(x, y) {
					//b.Print()
					//fmt.Println(number)
					return b.Sum() * number
				}
			}
		}
	}
	return 0
}

func (b *Board) Print() {
	for y := 0; y < 5; y++ {
		fmt.Println(b.tiles[y*5], b.tiles[1+y*5], b.tiles[2+y*5], b.tiles[3+y*5], b.tiles[4+y*5])
	}
}

func (b *Board) Sum() int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			tile := b.tiles[x+(y*5)]
			if !tile.marked {
				sum += tile.value
			}
		}
	}
	return sum
}

type ListNoder interface {
	Remove(*Board)
}

type ListNode struct {
	next  *ListNode
	board *Board
}

func (l *ListNode) Remove(b *Board) *ListNode {
	if l.board == b {
		return l.next
	}
	l.next = l.next.Remove(b)
	return l
}

type Lister interface {
	Append(*Board)
	Remove(*Board)
	MarkBoards(int)
	Size()
}

type List struct {
	head *ListNode
}

func (l *List) Append(b *Board) {
	if l.head == nil {
		l.head = &ListNode{nil, b}
		return
	}
	c := l.head
	for ; c.next != nil; c = c.next {
	}
	c.next = &ListNode{nil, b}
}

func (l *List) Remove(b *Board) {
	if l.head == nil {
		return
	}
	l.head = l.head.Remove(b)
}

func (l *List) MarkBoards(number int) int {
	res := 0
	for c := l.head; c != nil; c = c.next {
		tmp := c.board.Mark(number)
		if tmp != 0 {
			l.Remove(c.board)
			res = tmp
		}
	}
	return res
}

func (l *List) Size() int {
	size := 0
	c := l.head
	if c == nil {
		return size
	}
	for ; c != nil; c = c.next {
		size++
	}
	return size
}

func readInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	randomlist := strings.Split(scanner.Text(), ",")
	boards := &List{}

	// construct boards
	cboard := &Board{}
	bline := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) < 1 {
			boards.Append(cboard)
			cboard = &Board{}
			bline = 0
		} else {
			split := strings.Split(line, " ")
			ifound := 0
			for i := 0; i < len(split); i++ {
				value, error := strconv.Atoi(split[i])
				if error != nil {
				} else {
					cboard.tiles[ifound+(bline*5)] = Tile{value, false}
					ifound++
				}
			}
			if ifound > 0 {
				bline++
			}
		}

	}

	// Run bingo game
	res := 0
	for i := 0; i < len(randomlist); i++ {
		value, error := strconv.Atoi(randomlist[i])
		if error != nil {
			log.Fatal(error)
		} else {
			tmp := boards.MarkBoards(value)
			if tmp != 0 {
				res = tmp
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func main() {
	fmt.Println(readInput(os.Args[1]))
}
