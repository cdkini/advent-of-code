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
	connections := inputToMap() // 594
	fmt.Println(CountValidBags(connections))
}

func CountValidBags(connections map[string]map[string]int) int {
	count := 0
	for color := range connections {
		if color != "shiny gold" && isValid(color, connections) {
			count++
		}
	}
	return count
}

func isValid(curr string, connections map[string]map[string]int) bool {
	if curr == "shiny gold" {
		return true
	}
	children := connections[curr]
	if len(children) == 0 {
		return false
	}
	for child := range children {
		return isValid(child, connections)
	}
	return false
}

func inputToMap() map[string]map[string]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	connections := make(map[string]map[string]int, 0)
	scanner := bufio.NewScanner(file)
	var a []string

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		color := line[0] + " " + line[1]
		connections[color] = make(map[string]int, 0)
		for i := 4; i < len(line); i += 4 {
			if line[i] == "no" {
				break
			}
			count, _ := strconv.Atoi(line[i])
			nestedColor := line[i+1] + " " + line[i+2]
			connections[color][nestedColor] = count
		}

		a = append(a, color)
	}

	return connections
}
