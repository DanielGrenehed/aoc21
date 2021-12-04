package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BitCountNoder interface {
	Increment(int)
	MultiplyGammaEpsilon(int)
}

type BitCountNode struct {
	next  *BitCountNode
	count int
}

func (b *BitCountNode) Increment(bit int) {
	c := b
	if bit < 0 {
		return
	}
	for j := 0; j < bit; j++ {
		if c.next == nil {
			c.next = &BitCountNode{nil, 0}
		}
		c = c.next
	}
	c.count++
}

/*
	returns an int with all bit-counts larger than min set to 1
*/
func (b *BitCountNode) MultiplyGammaEpsilon(min int) int {
	gamma := 0
	epsilon := 0
	c := b
	for ; c != nil; c = c.next {
		gamma <<= 1
		epsilon <<= 1
		if c.count >= min {
			gamma++
		} else {
			epsilon++
		}
	}
	return gamma * epsilon
}

func submarinePowerConsumption(filename string) int {
	lcount := 0
	bitCounter := &BitCountNode{nil, 0}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// for each char in ['0', '1']
		line := scanner.Text()
		found := false
		bpos := 0
		for i := 0; i < len(line); i++ {
			if line[i] == '0' {
				bpos++
				found = true
			} else if line[i] == '1' {
				bitCounter.Increment(bpos)
				bpos++
				found = true
			}
		}
		if found {
			lcount++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return bitCounter.MultiplyGammaEpsilon(lcount / 2)
}

func main() {
	fmt.Println(submarinePowerConsumption(os.Args[1]))
}
