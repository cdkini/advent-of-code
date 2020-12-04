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
	passports := inputToSlice()
	CountValidPassports(passports)
	CountDataValidatedPassports(passports)
}

func CountDataValidatedPassports(passports []Passport) int {
	count := 0
	for _, passport := range passports {
		if isValidYear(passport["byr"], 1920, 2002) &&
			isValidYear(passport["iyr"], 2010, 2020) &&
			isValidYear(passport["eyr"], 2020, 2030) &&
			isValidHgt(passport["hgt"]) &&
			isValidHcl(passport["hcl"]) &&
			isValidEcl(passport["ecl"]) &&
			isValidPid(passport["pid"]) {
			count++
		}
	}
	fmt.Println("Valid passports:", count)
	return count
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func isValidYear(year string, minYear int, maxYear int) bool {
	if len(year) != 4 {
		return false
	}
	yearNum, _ := strconv.Atoi(year)
	return yearNum >= minYear && yearNum <= maxYear
}

// hgt (Height) - a number followed by either cm or in:
//   If cm, the number must be at least 150 and at most 193.
//   If in, the number must be at least 59 and at most 76.
func isValidHgt(hgt string) bool {
	// TODO: Open to implement!
	return true
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func isValidHcl(hcl string) bool {
	if len(hcl) != 7 || string(hcl[0]) != "#" {
		return false
	}
	// TODO: Open to implement!
	return true
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func isValidEcl(ecl string) bool {
	validColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return validColors[ecl]
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func isValidPid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	// TODO: Open to implement!
	return true
}

func CountValidPassports(passports []Passport) int {
	count := 0
	attrs := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, passport := range passports {
		isValid := true
		for _, attr := range attrs {
			if len(passport[attr]) == 0 {
				isValid = false
				break
			}
		}
		if isValid {
			count++
		}
	}
	fmt.Println("Valid passports:", count)
	return count
}

type Passport map[string]string

func emptyPassport() Passport {
	return Passport{"byr": "", "iyr": "", "eyr": "", "hgt": "", "hcl": "", "ecl": "", "pid": "", "cid": ""}
}

func inputToSlice() []Passport {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out []Passport
	curr := emptyPassport()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			out = append(out, curr)
			curr = emptyPassport()
		} else {
			for _, attrs := range strings.Split(line, " ") {
				attr := strings.Split(attrs, ":")
				for i := 0; i < len(attr); i += 2 {
					if _, ok := curr[string(attr[i])]; ok {
						curr[string(attr[i])] = string(attr[i+1])
					}
				}
			}
		}
	}
	return out
}
