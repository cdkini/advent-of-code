package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := inputToSlice()
	fmt.Println("Part One Tree Count:", GetTreeCount(1, 3, input))
	a := GetTreeCount(1, 1, input)
	b := GetTreeCount(1, 3, input)
	c := GetTreeCount(1, 5, input)
	d := GetTreeCount(1, 7, input)
	e := GetTreeCount(2, 1, input)
	fmt.Println("Part Two Tree Count:", a*b*c*d*e)
}

func GetTreeCount(rowMvmt int, colMvmt int, input []string) int {
	count := 0
	row := 0
	col := 0
	for row < len(input) {
		if input[row][col] != 46 {
			count++
		}
		row += rowMvmt
		col += colMvmt
		col %= len(input[0])
	}
	return count
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
		out = append(out, scanner.Text())
	}

	return out
}
