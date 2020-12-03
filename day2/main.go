package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	min      int
	max      int
	target   string
	password string
}

func main() {
	input := inputToSlice()
	GetValidPasswordCountByFrequency(input)
	GetValidPasswordCountByPosition(input)
}

func GetValidPasswordCountByFrequency(input []Input) {
	count := 0
	for _, in := range input {
		freq := findFrequency(in.target, in.password)
		if in.min <= freq && in.max >= freq {
			count++
		}
	}
	fmt.Println("Valid password count:", count)
}

func GetValidPasswordCountByPosition(input []Input) {
	count := 0
	for _, in := range input {
		validMin := string(in.password[in.min-1]) == in.target
		validMax := string(in.password[in.max-1]) == in.target
		if (validMin && !validMax) || (!validMin && validMax) {
			count++
		}
	}
	fmt.Println("Valid password count:", count)
}

func inputToSlice() []Input {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []Input
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		min, max := defineRange(temp[0])
		target := string(temp[1][0])
		password := temp[2]
		instance := Input{min, max, target, password}
		out = append(out, instance)
	}

	return out
}

func defineRange(str string) (int, int) {
	i := strings.Index(str, "-")
	min, _ := strconv.Atoi(str[:i])
	max, _ := strconv.Atoi(str[i+1:])
	return min, max
}

func findFrequency(target string, password string) int {
	freq := 0
	for _, char := range password {
		if string(char) == target {
			freq++
		}
	}
	return freq
}
