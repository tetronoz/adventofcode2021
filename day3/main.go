package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fileName := "input.txt"
	fp, _ := os.Open(fileName)
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	var digits[12]int

	for sc.Scan() {
		s := sc.Text()
		for i, c := range s {
			if rune(c) == '0' {
				digits[i] -= 1
			} else {
				digits[i] += 1
			}
		}
	}

	var gammaRate []string
	var epsilonRate []string

	for _, digit := range digits {
		if digit > 0 {
			gammaRate = append(gammaRate, "1")
			epsilonRate = append(epsilonRate, "0")
		} else {
			gammaRate = append(gammaRate, "0")
			epsilonRate = append(epsilonRate, "1")
		}
	}

	gamma, _ := strconv.ParseInt(strings.Join(gammaRate, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilonRate, ""), 2, 64)

	fmt.Println(gamma * epsilon)
}

func part2() {
	fileName := "input.txt"
	fp, _ := os.Open(fileName)
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	
	var report [][]string

	for sc.Scan() {
		s := sc.Text()
		report = append(report, strings.Split(s, ""))
	}

	oxygenRating := findOxygenRating(report)
	co2Rating := findCO2Rating(report)
	fmt.Printf("%v - %v\n", oxygenRating, co2Rating)
	fmt.Println(oxygenRating * co2Rating)
}

func findOxygenRating(data [][]string) int64 {
	positions := make([]bool, len(data))

	for i := range positions {
		positions[i] = true
	}

	options := len(data)

	for options > 1 {
		for j := range data[0] {
			var ones, zeroes int
			m := make(map[int][]int)
			
			for i := range(data) {
				if positions[i] == true {
					if data[i][j] == "1" {
						ones += 1
						m[1] = append(m[1], i) 
					} else {
						zeroes += 1
						m[0] = append(m[0], i)
					}
				}
			}

			if ones >= zeroes {
				for _, idx := range m[0] {
					positions[idx] = false
					options -= 1
				}
			} else {
				for _, idx := range m[1] {
					positions[idx] = false
					options -= 1
				}
			}
		}
	}
	
	if options == 1 {
		for i := range positions {
			if positions[i] == true {
				res, _ := strconv.ParseInt(strings.Join(data[i], ""), 2, 64)
				return res
			}
		}
	}
	
	return 0
}


func findCO2Rating(data [][]string) int64 {
	positions := make([]bool, len(data))

	for i := range positions {
		positions[i] = true
	}

	options := len(data)

	for options > 1 {
		for j := range data[0] {
			var ones, zeroes int
			m := make(map[int][]int)
			
			for i := range(data) {
				if positions[i] == true {
					if data[i][j] == "1" {
						ones += 1
						m[1] = append(m[1], i) 
					} else {
						zeroes += 1
						m[0] = append(m[0], i)
					}
				}
			}

			if zeroes <= ones {
				for _, idx := range m[1] {
					positions[idx] = false
					options -= 1
				}
			} else if ones < zeroes {
				for _, idx := range m[0] {
					positions[idx] = false
					options -= 1
				}
			}

			if options == 1 {
				break
			}
		}
	}
	
	if options == 1 {
		for i := range positions {
			if positions[i] == true {
				res, _ := strconv.ParseInt(strings.Join(data[i], ""), 2, 64)
				return res
			}
		}
	}
	
	return 0
}


func proper() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var oxygenRatingValues []string
	for sc.Scan(){
		oxygenRatingValues = append(oxygenRatingValues, sc.Text())
	}
	co2RatingValues := append([]string{}, oxygenRatingValues...) //A copy of oxygenRatingValues
}