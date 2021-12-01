package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Measurementer interface {
	Push(int)
}

type Measurements struct {
	a int
	b int
	c int
}

// returns true if sum of measurement increases
func (w *Measurements) Push(i int) bool {
	temp := w.c
	w.c = w.b
	w.b = w.a
	w.a = i
	if temp != -1 && i > temp {
		return true
	}
	return false
}

func CreateMeasurements() Measurements {
	return Measurements{-1, -1, -1}
}

func countIncrease(filename string) int {
	var count = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	measurements := CreateMeasurements()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, error := strconv.Atoi(scanner.Text())
		if error != nil {
			log.Fatal(error)
		}
		if measurements.Push(depth) {
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
