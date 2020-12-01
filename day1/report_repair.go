package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func inputToIntSlice() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var out []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, num)
	}

	return out
}

func PairAddsTo2020() {
	seen := make(map[int]struct{})
	nums := inputToIntSlice()
	for _, num := range nums {
		if _, ok := seen[2020-num]; ok {
			fmt.Println("Found pair:", num, 2020-num)
			fmt.Println("Answer:", num*(2020-num))
			return
		}
		seen[num] = struct{}{}
	}

	fmt.Println("Could not find pair")
}

func TripletAddsTo2020() {
	nums := inputToIntSlice()
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 2020 {
				right--
			} else if sum < 2020 {
				left++
			} else {
				fmt.Println("Found triplet:", nums[i], nums[left], nums[right])
				fmt.Println("Answer:", nums[i]*nums[left]*nums[right])
				return
			}
		}
	}

	fmt.Println("Could not find triplet")
}

func main() {
	PairAddsTo2020()
	TripletAddsTo2020()
}
