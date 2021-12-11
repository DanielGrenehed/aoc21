package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	unique := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		sfound := false
		for i := 0; i < len(split); i++ {
			if split[i] == "|" {
				sfound = true
				continue
			}

			if sfound {
				switch len(split[i]) {
				case 2:
					unique++
					break
				case 3:
					unique++
				case 4:
					unique++
				case 7:
					unique++
				}
			}
		}
	}
	return unique
}

func main() {
	fmt.Println(readFile(os.Args[1]))
}
