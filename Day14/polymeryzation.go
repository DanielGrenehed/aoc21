package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rule struct {
	match  string
	insert string
}

type Polymer struct {
	template string
	rules    []Rule
}

type Insertion struct {
	value string
	index int
}

func (p *Polymer) ApplyRules() {
	var insertions []Insertion
	for i := 1; i < len(p.template); i++ {
		for _, rule := range p.rules {
			if p.template[i-1:i+1] == rule.match {

				insertions = append(insertions, Insertion{rule.insert, i})
				break
			}
		}
	}
	p.Insert(insertions)
}

func (p *Polymer) Insert(insertions []Insertion) {
	for i, insertion := range insertions {
		p.template = p.template[:i+insertion.index] + insertion.value + p.template[i+insertion.index:]
	}
}

type ByteCount struct {
	char  byte
	count int
}

type PolyStat struct {
	stats []ByteCount
}

func (p *PolyStat) GetIndex(char byte) int {
	for i, bc := range p.stats {
		if char == bc.char {
			return i
		}
	}
	p.stats = append(p.stats, ByteCount{char, 0})
	return len(p.stats) - 1
}

func (p *PolyStat) GetMostCommonCharCount() int {
	res := p.stats[0].count
	for _, stat := range p.stats {
		if stat.count > res {
			res = stat.count
		}
	}
	return res
}

func (p *PolyStat) GetLeastCommonCharCount() int {
	res := p.stats[0].count
	for _, stat := range p.stats {
		if stat.count < res {
			res = stat.count
		}
	}
	return res
}

func (p *Polymer) CreateStat() PolyStat {
	stat := PolyStat{}
	for _, char := range p.template {
		index := stat.GetIndex(byte(char))
		stat.stats[index].count++
	}
	return stat
}

// get count of most common char
// get count of least common char
//

func GetPolymer(filename string) Polymer {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	polymer := Polymer{}
	scanner.Scan()
	polymer.template = scanner.Text()

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " -> ")

		if len(split) > 1 {
			polymer.rules = append(polymer.rules, Rule{split[0], split[1]})
		}
	}

	return polymer
}

func main() {
	polymer := GetPolymer(os.Args[1])

	for i := 0; i < 10; i++ {
		polymer.ApplyRules()
	}
	stat := polymer.CreateStat()
	fmt.Println(stat.GetMostCommonCharCount() - stat.GetLeastCommonCharCount())
}
