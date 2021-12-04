package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type StringListNoder interface {
	Append(string)
	Filter(int, byte)
	BitCount()
}

type StringListNode struct {
	next *StringListNode
	prev *StringListNode
	data string
}

func (s *StringListNode) Append(str string) {
	c := s
	for ; c.next != nil; c = c.next {
	}
	c.next = &StringListNode{nil, c, str}
}

// returns number of elements removed
func (s *StringListNode) Filter(bitpos int, include byte) int {
	c := s
	p := s
	fcount := 0
	for ; c != nil; c = c.next {

		if c.data[bitpos] != include || bitpos >= len(c.data) {
			/// remove from list
			p.next = c.next
			if c == s {
				s = p.next
			}
			fcount++
		}
		p = c
	}
	return fcount
}

func (s *StringListNode) BitCount() int {
	l := len(s.data)
	if l == 0 && s.next != nil {
		return s.next.BitCount()
	}
	return l
}

type StringLister interface {
	Append(string)
	Filter(int, byte)
	MostCommonBit(int)
	LeastCommonBit(int)
	BitCount()
	Size()
}

type StringList struct {
	head *StringListNode
	size int
}

func (s *StringList) Append(str string) {
	if s.head == nil {
		s.head = &StringListNode{nil, nil, str}
	} else {
		s.head.Append(str)
	}
	s.size++
}

func (s *StringList) Filter(bitpos int, include byte) {
	if s.head == nil {
		return
	} else {
		c := s.head
		for ; c != nil; c = c.next {
			if len(c.data) <= bitpos || c.data[bitpos] != include {
				if c.prev == nil {
					s.head = c.next
					if c.next != nil {
						c.next.prev = nil
					}
				} else {
					c.prev.next = c.next
					if c.next != nil {
						c.next.prev = c.prev
					}
				}
				s.size--
			}
		}

	}
}

func (s *StringList) MostCommonBit(bitpos int) byte {
	zcount := 0
	itcount := 0
	for c := s.head; c != nil; c = c.next {
		if c.data[bitpos] == '0' {
			zcount++
		}
		itcount++
	}
	if zcount > itcount/2 {
		return '0'
	}
	return '1'
}

func (s *StringList) LeastCommonBit(bitpos int) byte {
	mcommon := s.MostCommonBit(bitpos)
	if mcommon == '0' {
		return '1'
	}
	return '0'
}

func (s StringList) BitCount() int {
	if s.head == nil {
		return 0
	}
	return s.head.BitCount()
}

func (s StringList) Size() int {
	return s.size
}

func getLineArray(filename string) StringList {
	var bytearray StringList
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytearray.Append(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return bytearray
}

func binaryStringToInt(bs string) int {
	out := 0
	for i := 0; i < len(bs); i++ {
		if bs[i] == '0' {
			out <<= 1
		} else if bs[i] == '1' {
			out <<= 1
			out++
		}
	}
	return out
}

func oxygenGeneratorRating(byteArray StringList) int {
	blen := byteArray.BitCount()
	for i := 0; i < blen; i++ {
		byteArray.Filter(i, byteArray.MostCommonBit(i))
		if byteArray.Size() == 1 {
			return binaryStringToInt(byteArray.head.data)
		}
		fmt.Println("array size", byteArray.Size())
	}
	return 0
}

func CO2ScrubberRating(byteArray StringList) int {
	blen := byteArray.BitCount()
	for i := 0; i < blen; i++ {
		byteArray.Filter(i, byteArray.LeastCommonBit(i))
		if byteArray.Size() == 1 {
			return binaryStringToInt(byteArray.head.data)
		}
		fmt.Println("array size", byteArray.Size())
	}
	return 0
}

func main() {
	input := getLineArray(os.Args[1])
	fmt.Println("InputLines:", input.Size())
	ogr := oxygenGeneratorRating(input)
	fmt.Println("ogr:", ogr)
	input = getLineArray(os.Args[1])
	fmt.Println("InputLines:", input.Size())
	csr := CO2ScrubberRating(input)
	fmt.Println("csr:", csr)
	fmt.Println(ogr * csr)
}
