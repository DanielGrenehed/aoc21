package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Cave struct {
	name   string
	is_big bool
}

type Connection struct {
	a *Cave
	b *Cave
}

type CaveSystem struct {
	caves       []Cave
	connections []Connection
}

func isBigCave(name string) bool {
	for _, r := range name {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (cs *CaveSystem) AddCave(name string) *Cave {
	for i := 0; i < len(cs.caves); i++ {
		if cs.caves[i].name == name {
			return &cs.caves[i]
		}
	}
	cs.caves = append(cs.caves, Cave{name, isBigCave(name)})
	return &cs.caves[len(cs.caves)-1]
}

func (cs *CaveSystem) GetCave(name string) *Cave {
	for _, cave := range cs.caves {
		if cave.name == name {
			return &cave
		}
	}
	return nil
}

func (cs *CaveSystem) AddConnection(a string, b string) {
	ac := cs.AddCave(a)
	bc := cs.AddCave(b)
	cs.connections = append(cs.connections, Connection{ac, bc})
}

func (cs *CaveSystem) PrintConnections() {
	fmt.Println("Connections:")
	for i, c := range cs.connections {
		fmt.Print(i+1, ": ")
		if c.a.is_big {
			fmt.Print(c.a.name + "(Big)")
		} else {
			fmt.Print(c.a.name)
		}
		fmt.Print("-")
		if c.b.is_big {
			fmt.Println(c.b.name + "(Big)")
		} else {
			fmt.Println(c.b.name)
		}

	}
}

func (cs *CaveSystem) GetConnected(cave string) []Cave {
	var connected []Cave
	for _, c := range cs.connections {
		if c.a.name == cave {
			connected = append(connected, *c.b)
		}
		if c.b.name == cave {
			connected = append(connected, *c.a)
		}
	}
	return connected
}

func stringify(caves []Cave) string {
	out := ""
	for i, cave := range caves {
		if i > 0 {
			out += "," + cave.name
		} else {
			out += cave.name
		}
	}
	return out
}

func visited(history []Cave, cave Cave) bool {
	for _, historic := range history {
		if cave == historic {
			return true
		}
	}
	return false
}

func (cs *CaveSystem) findEnd(history []Cave) []string {
	current := history[len(history)-1]

	if current.name == "end" {
		return []string{stringify(history)}
	}

	caves := cs.GetConnected(current.name)
	var paths []string
	for _, connection := range caves {
		if connection.is_big || !visited(history, connection) {
			paths = append(paths, cs.findEnd(append(history, connection))...)
		}
	}
	return paths
}

func (cs *CaveSystem) GetPaths() []string {
	return cs.findEnd([]Cave{*cs.GetCave("start")})
}

func testConnections() {
	cs := &CaveSystem{}
	cs.AddConnection("A", "b")
	cs.AddConnection("start", "A")
	cs.AddConnection("c", "end")
	cs.AddConnection("b", "c")

	cs.PrintConnections()
	fmt.Println(cs.GetPaths())
}

func loadCaveSystem(filename string) *CaveSystem {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cs := &CaveSystem{}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "-")
		if len(input) > 1 {
			cs.AddConnection(input[0], input[1])
		}
	}
	return cs
}

func main() {
	if len(os.Args) > 1 {
		cs := loadCaveSystem(os.Args[1])
		res := cs.GetPaths()
		fmt.Println(res)
		fmt.Println(len(res))
	} else {
		testConnections()
	}

}
