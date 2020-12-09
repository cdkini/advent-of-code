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
	instructions := inputToSlice() // 594
	fmt.Println("Part One - Value of Accumulator:", ExecuteInstructions(instructions))
}

func ExecuteInstructions(instructions []Instruction) int {
	accumulator := 0
	seen := make(map[int]bool, 0)
	i := 0
	for i < len(instructions) {
		if _, ok := seen[i]; ok {
			break
		}
		seen[i] = true
		curr := instructions[i]
		if curr.operation == "acc" {
			accumulator += curr.argument
			i++
		} else if curr.operation == "jmp" {
			i += curr.argument
		} else if curr.operation == "nop" {
			i++
		}
	}
	return accumulator
}

type Instruction struct {
	operation string
	argument  int
}

func newInstruction(operation string, argument int) Instruction {
	return Instruction{operation, argument}
}

func inputToSlice() []Instruction {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		op := line[0]
		sign := line[1][0]
		arg, _ := strconv.Atoi(line[1][1:])
		if sign == '-' {
			arg = -arg
		}
		instruction := newInstruction(op, arg)
		instructions = append(instructions, instruction)
	}

	return instructions
}
