package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	nums := inputToSlice()
	fmt.Println("Part One - Invalid Num:", TraverseXMAS(nums))
	fmt.Println("Part Two - Encryption Weakness:", FindWeakness(nums, 15353384)) // Answer for part one
}

func FindWeakness(nums []int, target int) int {
	sum := 0
	left := 0
	right := 0
	for right < len(nums) {
		if sum > target {
			sum -= nums[left]
			left++
		} else if sum < target {
			sum += nums[right]
			right++
		} else {
			min, max := minMax(nums[left:right])
			return min + max
		}
	}
	return -1
}

func minMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, value := range nums {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func TraverseXMAS(nums []int) int {
	preamble := initPreamble(nums)
	for i := 25; i < len(nums); i++ {
		if isValid(nums[i], preamble) {
			updatePreamble(nums[i], nums[i-25], preamble)
			continue
		}
		return nums[i]
	}
	return -1
}

func initPreamble(nums []int) map[int]int {
	preamble := make(map[int]int, 0)
	i := 0
	for i < 25 {
		preamble[nums[i]]++
		i++
	}
	return preamble
}

func updatePreamble(newNum int, oldNum int, preamble map[int]int) {
	preamble[oldNum]--
	if preamble[oldNum] == 0 {
		delete(preamble, oldNum)
	}
	preamble[newNum]++
}

func isValid(curr int, preamble map[int]int) bool {
	for val := range preamble {
		diff := curr - val
		if _, ok := preamble[diff]; ok && val != diff {
			return true
		}
	}
	return false
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
