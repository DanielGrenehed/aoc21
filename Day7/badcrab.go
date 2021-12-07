package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AATNoder interface {
	isLeaf() bool
	hasLeft() bool
	hasRight() bool
	fuelConsumption(int) int
	find(int) *AATNode
}

type AATNode struct {
	value int
	left  *AATNode
	right *AATNode
	level int
	mult  int
}

func (n *AATNode) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *AATNode) hasLeft() bool {
	return n.left != nil
}

func (n *AATNode) hasRight() bool {
	return n.right != nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func fuelCounter(dist uint64) uint64 {
	return dist * (dist + 1) / 2
}

func (n *AATNode) fuelConsumption(dist int) uint64 {
	sum := fuelCounter(uint64(abs(dist-n.value))) * uint64(n.mult)
	if n.hasLeft() {
		sum += n.left.fuelConsumption(dist)
	}
	if n.hasRight() {
		sum += n.right.fuelConsumption(dist)
	}
	return sum
}

func (n *AATNode) find(val int) *AATNode {
	if val == n.value {
		return n
	} else if val < n.value {
		if n.hasLeft() {
			return n.left.find(val)
		}
	} else {
		if n.hasRight() {
			return n.right.find(val)
		}
	}
	return nil
}

func NewNode(val int) *AATNode {
	return &AATNode{val, nil, nil, 1, 1}
}

func skew(T *AATNode) *AATNode {
	if T == nil {
		return nil
	} else if !T.hasLeft() {
		return T
	} else if T.left.level == T.level {
		L := T.left
		T.left = L.right
		L.right = T
		return L
	}
	return T
}

func split(T *AATNode) *AATNode {
	if T == nil {
		return nil
	} else if !T.hasRight() || !T.right.hasRight() {
		return T
	} else if T.level == T.right.right.level {
		R := T.right
		T.right = R.left
		R.left = T
		R.level++
		return R
	}
	return T
}

func insert(val int, T *AATNode) *AATNode {
	if T == nil {
		return NewNode(val)
	} else if val < T.value {
		T.left = insert(val, T.left)
	} else if val > T.value {
		T.right = insert(val, T.right)
	} else {
		T.mult++
	}
	T = skew(T)
	T = split(T)
	return T
}

type AATreer interface {
	insert(int)
	rasTarget() int
}

type AATree struct {
	root *AATNode
}

func (t *AATree) insert(val int) {
	if t.root == nil {
		t.root = NewNode(val)
	} else {
		t.root = insert(val, t.root)
	}
}

func (t *AATree) rasTarget() uint64 {
	tpos := t.root.value
	base := t.root.fuelConsumption(tpos)
	dir := -1
	if t.root.fuelConsumption(tpos+1) < base {
		dir = 1
	}

	res := base
	for {
		tpos += dir
		val := t.root.fuelConsumption(tpos)
		if val > res {
			break
		} else {
			fmt.Println("better val: ", val, t.root.find(tpos))
			res = val
		}
	}

	return res
}

func readInput(filename string) uint64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	t := &AATree{nil}
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(split); i++ {
			pos, error := strconv.Atoi(split[i])
			if error != nil {
				fmt.Println("error:", split[i])
			} else {
				t.insert(pos)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return t.rasTarget()
}

func main() {
	fmt.Println(readInput(os.Args[1]))
}
