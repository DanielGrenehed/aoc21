package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dumboser interface {
	Charge()
	Flash()
	FlashDumbo(int)
	ClearFlashes() int
	Print()
	GetAdjecent(int) []int
}

type Dumbos struct {
	values []int
	width  int
}

func (d *Dumbos) Charge() {
	for i := 0; i < len(d.values); i++ {
		d.values[i]++
	}
}

func (d *Dumbos) Flash() {
	for i := 0; i < len(d.values); i++ {
		if d.values[i] == 10 {
			d.FlashDumbo(i)
		}
	}
}

func (d *Dumbos) FlashDumbo(pos int) {
	if d.values[pos] == 10 {
		d.values[pos]++

		adj := d.GetAdjecent(pos)
		//fmt.Println(pos, adj)
		for i := 0; i < len(adj); i++ {
			if d.values[adj[i]] != 10 {
				d.values[adj[i]]++
			}
			d.FlashDumbo(adj[i])
		}
	}
}

func (d *Dumbos) ClearFlashes() int {
	fc := 0
	for i := 0; i < len(d.values); i++ {
		if d.values[i] > 9 {
			fc++
			d.values[i] = 0
		}
	}
	return fc
}

func (d *Dumbos) GetAdjecent(center int) []int {
	var adj []int

	if int(center/d.width) == int((center-1)/d.width) {
		if center-1 > -1 {
			adj = append(adj, center-1)
		}
		if center-d.width-1 > -1 {
			adj = append(adj, center-d.width-1)
		}
		if center+d.width-1 < len(d.values) && center != 0 {
			adj = append(adj, center+d.width-1)
		}
	}

	if int(center/d.width) == int((center+1)/d.width) { // plus 1 is on same line
		if center+1 < len(d.values) {
			adj = append(adj, center+1)
		}
		if center+d.width+1 < len(d.values) {
			adj = append(adj, center+d.width+1)
		}
		if center+1-d.width > -1 {
			adj = append(adj, center+1-d.width)
		}
	}

	if center-d.width >= 0 {
		adj = append(adj, center-d.width)
	}

	if center+d.width < len(d.values) {
		adj = append(adj, center+d.width)
	}

	return adj
}

func (d *Dumbos) Print() {
	fmt.Println()
	for y := 0; y < d.width; y++ {
		out := ""
		for x := 0; x < d.width; x++ {
			out += strconv.Itoa(d.values[x+(y*d.width)])
		}
		fmt.Println(out)
	}
}

func DummyDumbo() *Dumbos {
	d := &Dumbos{}
	d.width = 3
	for i := 0; i < d.width*d.width; i++ {
		d.values = append(d.values, 1)
	}
	d.values[0] = 7
	return d
}

func cycle(d *Dumbos) int {
	d.Charge()
	d.Flash()
	return d.ClearFlashes()
}

func testOverflow() {
	d := DummyDumbo()
	for i := 0; i < 8; i++ {
		d.Print()
		cycle(d)
	}
	d.Print()
}

func sumFlashes(d *Dumbos, cycles int) int {
	sum := 0
	for i := 0; i < cycles; i++ {
		sum += cycle(d)
		d.Print()
	}
	return sum
}

func findFlashSync(d *Dumbos) int {
	count := 1
	for {
		if cycle(d) == len(d.values) {
			return count
		}
		count++
	}
	return 0
}

func ImportDumbos(filename string) *Dumbos {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	octopuses := &Dumbos{}
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			v, e := strconv.Atoi(string(line[i]))
			if e == nil {
				octopuses.values = append(octopuses.values, v)
			}
		}
		octopuses.width = len(line)
	}
	return octopuses
}

func main() {

	if len(os.Args) > 1 {
		d := ImportDumbos(os.Args[1])
		fmt.Println(findFlashSync(d))
	} else {
		testOverflow()
	}

}
