package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countIncrease(filename string) int {
	var count = 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pr_depth := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, error := strconv.Atoi(scanner.Text())
		if error != nil {
			log.Fatal(error)
		}
		if pr_depth > -1 {
			if depth > pr_depth {
				count++
			}
		}
		pr_depth = depth

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count
}

func main() {
	fmt.Println(countIncrease(os.Args[1]))
}
