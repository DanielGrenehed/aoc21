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

/*
	Returns true if all letters in string are uppercase letters
*/
func isUppercase(name string) bool {
	for _, r := range name {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/*
	Returns a Cave with the given name, creates one if not found
*/
func (cs *CaveSystem) AddCave(name string) *Cave {
	for i := 0; i < len(cs.caves); i++ {
		if cs.caves[i].name == name {
			return &cs.caves[i]
		}
	}
	cs.caves = append(cs.caves, Cave{name, isUppercase(name)})
	return &cs.caves[len(cs.caves)-1]
}

/*
	Returns a Cave with the given name if found, else nil
*/
func (cs *CaveSystem) GetCave(name string) *Cave {
	for _, cave := range cs.caves {
		if cave.name == name {
			return &cave
		}
	}
	return nil
}

/*
	Creates a connection between caves with the given names
*/
func (cs *CaveSystem) AddConnection(a string, b string) {
	ac := cs.AddCave(a)
	bc := cs.AddCave(b)
	cs.connections = append(cs.connections, Connection{ac, bc})
}

/*
	Prints all connections between caves
*/
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

/*
	Returns an array of all caves connected to the cave with the given argument as name
*/
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

/*
	Turns an array of caves into a string of comma-separated cave names
*/
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

/*
	Returns true if a small cave exists twice in array
*/
func smallVisitedTwice(history []Cave) bool {
	for i, cave := range history {
		if !cave.is_big {
			for j := i + 1; j < len(history); j++ {
				if cave == history[j] {
					return true
				}
			}
		}
	}
	return false
}

/*
	Returns true if array contains cave
*/
func visited(history []Cave, cave Cave) bool {
	for _, historic := range history {
		if cave == historic {
			return true
		}
	}
	return false
}

/*
	Returns an array of strings contaning all paths to "end" cave
*/
func (cs *CaveSystem) findEnd(history []Cave) []string {
	current := history[len(history)-1]

	if current.name == "end" {
		return []string{stringify(history)}
	}

	caves := cs.GetConnected(current.name)

	var paths []string
	for _, connected := range caves {
		if connected.name == "start" {
			continue
		}
		if connected.is_big || !visited(history, connected) || !smallVisitedTwice(history) {
			if connected.name == "start" {
				log.Fatal("Should have skipped start")
			}
			paths = append(paths, cs.findEnd(append(history, connected))...)
		}
	}
	return paths
}

/*
	Returns an array of strings containing all paths from CaveSystems start to its end
*/
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

/*
	Creates a cave system from file
*/
func loadCaveSystem(filename string) *CaveSystem {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cs := &CaveSystem{}
	for scanner.Scan() {
		input := strings.Split(strings.TrimSpace(scanner.Text()), "-")
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
		/* for _, str := range res {
			fmt.Println(str)
		} */
		fmt.Println(len(res))
	} else {
		testConnections()
	}

}
