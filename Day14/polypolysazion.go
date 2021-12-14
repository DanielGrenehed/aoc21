package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Reaction struct {
	elements map[string]int64
}

func (r *Reaction) IncreaseElement(element string, inc int64) {
	res, found := r.elements[element]
	if found {
		r.elements[element] = res + inc
	} else {
		r.elements[element] = inc
	}
}

type Polymer struct {
	state      Reaction
	rules      map[string]string
	charcounts map[byte]int64
}

func (p *Polymer) IncrementStat(char byte, inc int64) {
	_, f := p.charcounts[char]
	if f {
		p.charcounts[char] += inc
	} else {
		p.charcounts[char] = inc
	}
}

func CreatePolymer(template string) Polymer {
	p := Polymer{Reaction{make(map[string]int64)}, make(map[string]string), make(map[byte]int64)}
	p.IncrementStat(template[0], 1)
	for i := 1; i < len(template); i++ {
		p.state.IncreaseElement(template[i-1:i+1], 1)
		p.IncrementStat(template[i], 1)
	}
	return p
}

func (p *Polymer) AddRule(match string, insert string) {
	p.rules[match] = insert
}

func (p *Polymer) React() {
	next_state := Reaction{make(map[string]int64)}
	for key, value := range p.state.elements {
		insert, found := p.rules[key]
		if found {
			a := string(key[0]) + insert
			b := insert + string(key[1])
			next_state.IncreaseElement(a, value)
			next_state.IncreaseElement(b, value)

			p.IncrementStat(insert[0], value)
		} else {
			next_state.IncreaseElement(key, value)
		}
	}
	p.state = next_state
}

func (p *Polymer) MostCommonCharCount() int64 {
	var res int64 = 0
	for _, value := range p.charcounts {
		if value > res {
			res = value
		}
	}
	return res
}

func (p *Polymer) LeastCommonCharCount() int64 {
	var res int64
	i := 0
	for _, value := range p.charcounts {
		if i == 0 {
			res = value
		}
		if value < res {
			res = value
		}
		i++
	}
	return res
}

func GetPolymer(filename string) Polymer {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := CreatePolymer(scanner.Text())

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " -> ")

		if len(split) > 1 {
			polymer.AddRule(split[0], split[1])
		}
	}

	return polymer
}

func main() {
	polymer := GetPolymer(os.Args[1])

	for i := 0; i < 40; i++ {
		polymer.React()
	}

	fmt.Println(polymer.MostCommonCharCount() - polymer.LeastCommonCharCount())
}
