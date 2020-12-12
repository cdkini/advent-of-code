package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	seats := inputToSlice()
	fmt.Println("Part One - Num Occupants:", CountOccupancies(seats))
}

// Part One - Main func
func CountOccupancies(seats [][]string) int {
	runSimulation(seats)
	occupancies := 0
	for row := range seats {
		for col := range seats[0] {
			if seats[row][col] == "#" {
				occupancies++
			}
		}
	}
	return occupancies
}

// Part One - Wrapper around updateSeats helper
func runSimulation(seats [][]string) {
	for updateSeats(seats) != 0 {
		updateSeats(seats)
	}
}

// Part One - Iterates through seats, identifies necessary changes, and updates seats with changes
func updateSeats(seats [][]string) int {
	newSeats := copyMatrix(seats)
	updates := 0
	for row := range seats {
		for col := range seats[0] {
			if seats[row][col] == "L" && checkNeighbors(row, col, seats) == 0 {
				newSeats[row][col] = "#"
				updates++
			}
			if seats[row][col] == "#" && checkNeighbors(row, col, seats) >= 4 {
				newSeats[row][col] = "L"
				updates++
			}
		}
	}
	copy(seats, newSeats)
	return updates
}

// Part One - Check adjacent neighbors to determine occupancy
func checkNeighbors(row int, col int, seats [][]string) int {
	neighbors := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			if r+row < 0 || r+row >= len(seats) || c+col < 0 || c+col >= len(seats[0]) {
				continue
			}
			if seats[r+row][c+col] == "#" {
				neighbors++
			}
		}
	}
	return neighbors
}

// Both Parts - Used to make a deep copy of our seats matrix
func copyMatrix(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

// Both Parts - Stores input as a slice of string slices
func inputToSlice() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seats [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		seats = append(seats, row)
	}

	return seats
}
