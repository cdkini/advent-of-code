package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	connections, bagCounts := inputToMap() // 594
	fmt.Println("Part One - Valid Bags:", CountValidBags(connections))
	fmt.Println("Part Two - Nested Bags:", CountNestedBags(bagCounts))
}

func CountNestedBags(bagCounts map[string]map[string]int) int {
	return countNestedBags("shiny gold", bagCounts)
}

func countNestedBags(curr string, bagCounts map[string]map[string]int) int {
	if len(bagCounts[curr]) == 0 {
		return 0
	}
	contents := bagCounts[curr]
	numBags := 0
	for bag, count := range contents {
		numBags += (count * countNestedBags(bag, bagCounts)) + count
	}
	return numBags
}

func CountValidBags(connections map[string][]string) int {
	return countValidBags("shiny gold", make(map[string]bool, 0), connections)
}

func countValidBags(curr string, valid map[string]bool, connections map[string][]string) int {
	if len(connections[curr]) == 0 {
		return 0
	}
	for _, child := range connections[curr] {
		if _, ok := valid[child]; !ok {
			valid[child] = true
			countValidBags(child, valid, connections)
		}
	}
	return len(valid)
}

func inputToMap() (map[string][]string, map[string]map[string]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	connections := make(map[string][]string, 0)
	bagCounts := make(map[string]map[string]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		parent := line[0] + " " + line[1]
		bagCounts[parent] = make(map[string]int, 0)
		for i := 4; i < len(line); i += 4 {
			if line[i] == "no" {
				break
			}
			count, _ := strconv.Atoi(line[i])
			child := line[i+1] + " " + line[i+2]
			bagCounts[parent][child] = count

			if _, ok := connections[child]; !ok {
				connections[child] = make([]string, 0)
			}
			connections[child] = append(connections[child], parent)
		}
	}

	return connections, bagCounts
}
