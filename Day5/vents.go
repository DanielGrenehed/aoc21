package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Liner interface {
	IsValid() bool
	IsHorizontal() bool
	IsVertical() bool
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (l *Line) IsValid() bool {
	if l.IsHorizontal() || l.IsVertical() {
		return true
	}
	return false
}

func (l *Line) IsHorizontal() bool {
	return l.y1 == l.y2
}

func (l *Line) IsVertical() bool {
	return l.x1 == l.x2
}

type Maper interface {
	SetLine(*Line)
	OverlappingCount() int
	Print()
}

type Map struct {
	arr         [1000 * 1000]int
	overlapping int
}

func (m *Map) SetLine(l *Line) {
	if l.IsHorizontal() {
		for x := l.x1; x <= l.x2; x++ {
			i := x + (l.y1 * 1000)
			m.arr[i]++
			if m.arr[i] == 2 {
				m.overlapping++
			}
		}
	} else if l.IsVertical() {
		for y := l.y1; y <= l.y2; y++ {
			i := l.x1 + (y * 1000)
			m.arr[i]++
			if m.arr[i] == 2 {
				m.overlapping++
			}
		}
	}
}

func (m *Map) OverlappingCount() int {
	count := 0
	for i := 0; i < 1000*1000; i++ {
		if m.arr[i] > 1 {
			count++
		}
	}
	return count
}

func (m *Map) Print() {
	for y := 0; y < 1000; y++ {
		out := ""
		for x := 0; x < 1000; x++ {
			out += strconv.Itoa(m.arr[x+y*1000])
		}
		fmt.Println(out)
	}
}

func ConstructMap() Map {
	aMap := Map{}
	for i := 0; i < 1000*1000; i++ {
		aMap.arr[i] = 0
	}
	aMap.overlapping = 0
	return aMap
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func readInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := ConstructMap()
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), " -> ", ",", 1)
		split := strings.Split(line, ",")
		if len(split) > 3 {
			x1, error := strconv.Atoi(split[0])
			y1, error := strconv.Atoi(split[1])
			x2, error := strconv.Atoi(split[2])
			y2, error := strconv.Atoi(split[3])
			if error != nil {
				fmt.Println("error:", line)
			} else {
				l := Line{Min(x1, x2), Min(y1, y2), Max(x1, x2), Max(y1, y2)}
				if l.IsValid() {
					fmt.Println(l)

				} else {
					fmt.Println("Invalid:", l)
				}
				m.SetLine(&l)
			}

		} else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m.OverlappingCount()
}

func main() {
	fmt.Println(readInput(os.Args[1]))
}
