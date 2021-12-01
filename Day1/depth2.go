package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Slider interface {
	Push(int)
}

type Slide struct {
	a int
	b int
	c int
}

// returns true if sum of measurement increases
func (s *Slide) Push(i int) bool {
	temp := s.c
	s.c = s.b
	s.b = s.a
	s.a = i
	if temp != -1 && i > temp {
		return true
	}
	return false
}

func CreateSlide() Slide {
	return Slide{-1, -1, -1}
}

func countIncrease(filename string) int {
	var count = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	slider := CreateSlide()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, error := strconv.Atoi(scanner.Text())
		if error != nil {
			log.Fatal(error)
		}
		if slider.Push(depth) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}

func main() {
	fmt.Println(countIncrease(os.Args[1]))
}
