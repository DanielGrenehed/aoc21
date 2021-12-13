package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Paper struct {
	points []Point
}

func (p *Paper) FoldX(fold int) {
	for i, pnt := range p.points {
		if pnt.x > fold {
			p.points[i].x = -(pnt.x - fold*2)
		}
	}
}

func (p *Paper) FoldY(fold int) {
	for i, pnt := range p.points {
		if pnt.y > fold {
			p.points[i].y = -(pnt.y - fold*2)
		}
	}
}

func (p *Paper) FilterDuplicates() {
	for i, pnt := range p.points {
		for j := i + 1; j < len(p.points); j++ {
			if pnt.x == p.points[j].x && pnt.y == p.points[j].y {
				// remove point
				p.points = append(p.points[:j], p.points[j+1:]...)
			}
		}
	}
}

func (p *Paper) Width() int {
	w := 0
	for _, pnt := range p.points {
		if pnt.x > w {
			w = pnt.x
		}
	}
	return w
}

func (p *Paper) Height() int {
	h := 0
	for _, pnt := range p.points {
		if pnt.y > h {
			h = pnt.y
		}
	}
	return h
}

func (p *Paper) hasPoint(x int, y int) bool {
	for _, pnt := range p.points {
		if pnt.x == x && pnt.y == y {
			return true
		}
	}
	return false
}

func (p *Paper) PrintPoints() {
	h := p.Height()
	w := p.Width()
	for y := 0; y <= h; y++ {
		out := ""
		for x := 0; x <= w; x++ {
			if p.hasPoint(x, y) {
				out += "#"
			} else {
				out += "."
			}
		}
		fmt.Println(out)
	}
}

func getFoldedPaper(filename string) *Paper {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	paper := &Paper{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		if len(split) > 1 {
			v1, err := strconv.Atoi(split[0])
			v2, err := strconv.Atoi(split[1])
			if err == nil {
				point := Point{v1, v2}
				paper.points = append(paper.points, point)
			}
		} else {
			split = strings.Split(line, "=")
			if len(split) == 2 {
				fold, err := strconv.Atoi(split[1])
				if err == nil {
					if strings.HasSuffix(split[0], "x") {
						//paper.PrintPoints()
						paper.FoldX(fold)
						paper.FilterDuplicates()
						return paper
					} else if strings.HasSuffix(split[0], "y") {
						//paper.PrintPoints()
						paper.FoldY(fold)
						paper.FilterDuplicates()
						return paper
					}
				}

			}
		}
	}
	return paper
}

func main() {
	paper := getFoldedPaper(os.Args[1])
	fmt.Println(paper.points)
	//paper.PrintPoints()

	fmt.Println(len(paper.points))
}
