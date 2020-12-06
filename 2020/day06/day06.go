package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	groups := inputToSlice()
	fmt.Println("Part One - Sum of Groups:", SumOfGroups(groups))
	fmt.Println("Part Two - Sum of Groups:", SumOfUnanimousGroups(groups))
}

func SumOfUnanimousGroups(groups []map[rune]int) int {
	sum := 0
	for _, group := range groups {
		for k, v := range group {
			if k != 0 && v == group[rune(0)] {
				sum++
			}
		}
	}
	return sum
}

func SumOfGroups(groups []map[rune]int) int {
	sum := 0
	for _, group := range groups {
		for k, v := range group {
			if k != 0 && v > 0 {
				sum++
			}
		}
	}
	return sum
}

type Group map[rune]int

func newGroup() Group {
	group := make(map[rune]int)
	for i := 0; i < 26; i++ {
		key := rune(i + 97)
		group[key] = 0
	}
	group[rune(0)] = 0
	return group
}

func inputToSlice() []map[rune]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var groups []map[rune]int
	curr := newGroup()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groups = append(groups, curr)
			curr = newGroup()
		} else {
			curr[rune(0)]++
			for _, question := range line {
				if _, ok := curr[question]; ok {
					curr[question]++
				}
			}
		}
	}
	groups = append(groups, curr)
	return groups
}
