package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Seaer interface {
	AddFish(int)
	Cycle()
	CountFish() int
}

type Sea struct {
	fish [9]int
}

func (s *Sea) AddFish(day int) {
	s.fish[day]++
}

func (s *Sea) Cycle() {
	carry := 0
	for i := 8; i > -1; i-- {
		temp := s.fish[i]
		s.fish[i] = carry
		carry = temp
	}
	s.fish[6] += carry
	s.fish[8] += carry
}

func (s *Sea) CountFish() int {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += s.fish[i]
	}
	return sum
}

func CreateSea() *Sea {
	s := &Sea{}
	for i := 0; i < 9; i++ {
		s.fish[i] = 0
	}
	return s
}

func readInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sea := CreateSea()
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(split); i++ {
			day, error := strconv.Atoi(split[i])
			if error != nil {
				fmt.Println("error:", split[i])
			} else {
				sea.AddFish(day)
			}
		}
	}

	for i := 0; i < 256; i++ {
		sea.Cycle()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sea.CountFish()
}

func main() {
	fmt.Println(readInput(os.Args[1]))
}
