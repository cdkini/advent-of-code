package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	bags := inputToMap()
	for color, bag := range bags {
		fmt.Println(color, bag.children)
	}
	fmt.Println("Part One - Num of Bags: ", CountNumValidBags(bags))
}

func CountNumValidBags(adjList map[string]*Bag) int {
	count := 0
	for _, curr := range adjList {
		if curr.color == "shinygold" {
			continue
		}
		if helper(curr, adjList) {
			count++
		}
	}
	return count
}

func helper(curr *Bag, adjList map[string]*Bag) bool {
	if adjList[curr.color].color == "shinygold" {
		return true
	}
	for _, child := range curr.children {
		return helper(child, adjList)
	}
	return false
}

type Bag struct {
	color    string
	children []*Bag
}

func newBag(color string) *Bag {
	return &Bag{color, make([]*Bag, 0)}
}

func (b *Bag) String() string {
	return fmt.Sprintf("%s", b.color)
}

func (b *Bag) addChild(bag *Bag) {
	b.children = append(b.children, bag)
}

func inputToMap() map[string]*Bag {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bags := make(map[string]*Bag, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		color := line[0] + line[1]

		var curr *Bag
		if bag, ok := bags[color]; ok {
			curr = bag
		} else {
			curr = newBag(color)
		}

		for i := 4; i < len(line); i += 4 {
			if line[i] == "no" {
				break
			}
			child := addChildToBags(i, line, bags)
			curr.addChild(child)
		}
		bags[color] = curr
	}

	return bags
}

func addChildToBags(i int, line []string, bags map[string]*Bag) *Bag {
	color := line[i+1] + line[i+2]
	if _, ok := bags[color]; ok {
		return bags[color]
	}
	bags[color] = newBag(color)
	return bags[color]
}
