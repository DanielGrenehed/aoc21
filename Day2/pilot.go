package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Positioner interface {
	Forward(int)
	Down(int)
	Up(int)
	Horizontal()
	Vertical()
}

type Position struct {
	horizontal int
	vertical   int
}

func (p *Position) Forward(x int) {
	p.horizontal += x
}

func (p *Position) Down(x int) {
	p.vertical += x
}

func (p *Position) Up(x int) {
	p.vertical -= x
}

func (p Position) Horizontal() int {
	return p.horizontal
}

func (p Position) Vertical() int {
	return p.vertical
}

func CreatePosition() Position {
	return Position{0, 0}
}

func pilot(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pos := CreatePosition()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		value, error := strconv.Atoi(split[1])
		if error != nil {
			log.Fatal(error)
		}
		switch split[0] {
		case "forward":
			pos.Forward(value)
			break
		case "up":
			pos.Up(value)
			break
		case "down":
			pos.Down(value)
			break
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return pos.Horizontal() * pos.Vertical()
}

func main() {
	fmt.Println(pilot(os.Args[1]))
}
