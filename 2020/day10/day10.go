package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := inputToSlice()
	fmt.Println("Part One - Product Of Jolt Differences:", ProductOfJoltDiffs(adapters))
}

func ProductOfJoltDiffs(adapters []int) int {
	diffs := getJoltDifferences(adapters)
	return diffs[0] * diffs[2]
}

func getJoltDifferences(adapters []int) []int {
	sort.Ints(adapters)
	diffs := []int{0, 0, 0}
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	curr := 0
	for _, num := range adapters {
		diff := num - curr
		diffs[diff-1]++
		curr = num
	}

	return diffs
}

func inputToSlice() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	return nums
}
