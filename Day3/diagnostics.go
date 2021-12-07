package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BitCountNoder interface {
	Increment(int)
	Decrement(int)
	At(int) *BitCountNode
	Epsilon() int
	Gamma() int
}

type BitCountNode struct {
	next  *BitCountNode
	count int
}

func (b *BitCountNode) Increment(bit int) {
	if bit < 0 {
		return
	}
	b.At(bit).count++
}

func (b *BitCountNode) Decrement(bit int) {
	if bit < 0 {
		return
	}
	b.At(bit).count--
}

func (b *BitCountNode) At(i int) *BitCountNode {
	c := b
	if i < 0 {
		return nil
	}
	for j := 0; j < i; j++ {
		if c.next == nil {
			c.next = &BitCountNode{nil, 0}
		}
		c = c.next
	}
	return c
}

func (b *BitCountNode) Epsilon() int {
	epsilon := 0
	for c := b; c != nil; c = c.next {
		epsilon <<= 1
		if c.count < 0 {
			epsilon++
		}
	}
	return epsilon
}

func (b *BitCountNode) Gamma() int {
	gamma := 0
	for c := b; c != nil; c = c.next {
		gamma <<= 1
		if c.count >= 0 {
			gamma++
		}
	}
	return gamma
}

func submarinePowerConsumption(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	bytes := &BitCountNode{nil, 0}
	for scanner.Scan() {
		line := scanner.Text()
		bitpos := 0
		for i := 0; i < len(line); i++ {
			if line[i] == '0' {
				bytes.Decrement(bitpos)
				bitpos++
			} else if line[i] == '1' {
				bytes.Increment(bitpos)
				bitpos++
			}
		}
	}

	return bytes.Epsilon() * bytes.Gamma()
}

func main() {
	fmt.Println(submarinePowerConsumption(os.Args[1]))
}
