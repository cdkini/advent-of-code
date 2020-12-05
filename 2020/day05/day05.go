package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	seats := inputToSlice()
	fmt.Println("Part One - Highest Seat ID:", GetHighestSeatID(seats))
	fmt.Println("Part Two - Your Seat ID:", FindYourSeat(seats))
}

func FindYourSeat(seats []string) int {
	var ids []int
	for _, seat := range seats {
		ids = append(ids, getSeatID(seat))
	}
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i] != ids[i-1]+1 {
			return ids[i] - 1
		}
	}
	return -1
}

func GetHighestSeatID(seats []string) int {
	highest := 0
	for _, seat := range seats {
		curr := getSeatID(seat)
		if curr > highest {
			highest = curr
		}
	}
	return highest
}

func getSeatID(seat string) int {
	row := getSeatRow(seat)
	col := getSeatCol(seat)
	return row*8 + col
}

func getSeatRow(seat string) int {
	low := 0
	high := 127
	for i := 0; i < len(seat)-3; i++ {
		mid := low + (high-low)/2
		letter := string(seat[i])
		if letter == "F" {
			high = mid
		} else if letter == "B" {
			low = mid + 1
		}
	}
	return low + (high-low)/2
}

func getSeatCol(seat string) int {
	low := 0
	high := 7
	for i := len(seat) - 3; i < len(seat); i++ {
		mid := low + (high-low)/2
		letter := string(seat[i])
		if letter == "R" {
			low = mid + 1
		} else if letter == "L" {
			high = mid
		}
	}
	res := float64(low) + (float64(high)-float64(low))/2.0
	return int(math.Round(res))
}

func inputToSlice() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}
	return out
}
