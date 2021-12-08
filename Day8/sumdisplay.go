package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Digit struct {
	chars string
	value int
}

type SDisplayer interface {
	GetNumber(string) int
	AddDigit(string)
	FindSegments()
	Solve6()
	Solve5()
	Solve3()
	PrintDigits()
	Clear()
}

type SDisplay struct {
	digits [10]Digit
	apos   int
}

func (d *SDisplay) GetNumber(num string) int {
	for i := 0; i < 10; i++ {
		if len(num) == len(d.digits[i].chars) {
			if contains(num, d.digits[i].chars) {
				return d.digits[i].value
			}
		}
	}
	return -1
}

func (d *SDisplay) AddDigit(chars string) {

	switch len(chars) {
	case 2:
		d.digits[d.apos] = Digit{chars, 1}
		break
	case 3:
		d.digits[d.apos] = Digit{chars, 7}
		break
	case 4:
		d.digits[d.apos] = Digit{chars, 4}
		break
	case 7:
		d.digits[d.apos] = Digit{chars, 8}
		break
	default:
		d.digits[d.apos] = Digit{chars, -1}
	}
	d.apos++
	if d.apos >= 10 {
		d.apos = 0
	}
}

func (d *SDisplay) GetDigit(i int) *Digit {
	for p := 0; p < 10; p++ {
		if d.digits[p].value == i {
			return &d.digits[p]
		}
	}
	return nil
}

func (d *SDisplay) Solve6() {
	for i := 0; i < 10; i++ {
		dig := d.digits[i]
		if len(dig.chars) == 6 {
			if contains(dig.chars, d.GetDigit(7).chars) {
				if contains(dig.chars, d.GetDigit(4).chars) {
					d.digits[i].value = 9
				} else {
					d.digits[i].value = 0
				}
			} else {
				d.digits[i].value = 6
			}
		}
	}
}

func remove(a string, b string) string {
	out := ""
	for i := 0; i < len(a); i++ {
		if find(a[i], b) == -1 {
			out += string(a[i])
		}
	}
	return out
}

func (d *SDisplay) Solve5() {
	comp := remove(d.GetDigit(4).chars, d.GetDigit(3).chars)
	for i := 0; i < 10; i++ {
		dig := d.digits[i]
		if len(dig.chars) == 5 && dig.value == -1 {
			if contains(dig.chars, comp) {
				d.digits[i].value = 5
			} else {
				d.digits[i].value = 2
			}
		}
	}
}

func (d *SDisplay) Solve3() {
	for i := 0; i < 10; i++ {
		dig := d.digits[i]
		if len(dig.chars) == 5 {
			if contains(dig.chars, d.GetDigit(1).chars) {
				d.digits[i].value = 3
			}
		}
	}
}

func (d *SDisplay) FindSegments() {
	d.Solve6()
	d.Solve3()
	d.Solve5()
}

func (d *SDisplay) PrintDigits() {
	for i := 0; i < 10; i++ {
		fmt.Println("", d.digits[i].value, ": ", d.digits[i].chars)
	}
}

func (d *SDisplay) Clear() {
	for i := 0; i < 10; i++ {
		d.digits[i] = Digit{"", -1}
	}
	d.apos = 0
}

func find(f byte, s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == f {
			return i
		}
	}
	return -1
}

func contains(b string, a string) bool {
	for i := 0; i < len(a); i++ {
		if find(a[i], b) == -1 {
			return false
		}
	}
	return true
}

func readFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	disp := SDisplay{}
	disp.Clear()
	out := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		sfound := false
		res := 0
		for i := 0; i < len(split); i++ {
			if split[i] == "|" {
				sfound = true
				disp.FindSegments()
				continue
			}

			if sfound {
				res *= 10
				res += disp.GetNumber(split[i])
			} else {
				disp.AddDigit(split[i])
			}
		}
		out += res
		disp.Clear()
	}
	return out
}

func main() {
	fmt.Println(readFile(os.Args[1]))
}
