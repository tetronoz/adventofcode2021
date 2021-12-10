package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)


func main() {
	part1()
	part2()
}

func part1() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)


	var score int
	for sc.Scan() {
		score += getIllegalChar(sc.Text())
	}

	fmt.Println(score)
}

func part2() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	var missingScore []int
	for sc.Scan() {
		l := sc.Text()
		score := getIllegalChar(l)
		if score == 0 {
			missingScore = append(missingScore, getMissingChars(l))
		}
	}

	sort.Ints(missingScore)

	middle := len(missingScore) / 2

	fmt.Println(missingScore[middle])
}

func getIllegalChar(s string) int {
	scoreMap := map[string]int {
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	openings := map[string]bool {
		"[": true,
		"(": true,
		"{": true,
		"<": true,
	}

	closings := map[string]bool {
		")": true,
		"]": true,
		"}": true,
		">": true,
	}

	m := map[string]string {
		"[": "]",
		"(": ")",
		"{": "}",
		"<": ">",
	}

	var stack []string

	for _, c := range strings.Split(s, "") {

		_, in_openings := openings[c]
		_, in_closings := closings[c]

		if len(stack) == 0 || in_openings {
			stack = append(stack, c)
		} else {
			top := stack[len(stack)-1]
			if in_closings && m[top] == c {
				stack = stack[:len(stack)-1]
			} else {
				return scoreMap[c]
			}
		}
	}
	return 0
}


func getMissingChars(s string) int {
	scoreMap := map[string]int {
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	openings := map[string]bool {
		"[": true,
		"(": true,
		"{": true,
		"<": true,
	}

	closings := map[string]bool {
		")": true,
		"]": true,
		"}": true,
		">": true,
	}

	m := map[string]string {
		"[": "]",
		"(": ")",
		"{": "}",
		"<": ">",
	}

	var stack []string

	for _, c := range strings.Split(s, "") {

		_, in_openings := openings[c]
		_, in_closings := closings[c]

		if len(stack) == 0 || in_openings {
			stack = append(stack, c)
		} else {
			top := stack[len(stack)-1]
			if in_closings && m[top] == c {
				stack = stack[:len(stack)-1]
			}
		}
	}

	var score int
	for i := len(stack) - 1; i >=0; i-- {
		score *= 5
		c := stack[i]
		expected := m[c]
		score += scoreMap[expected]
	}

	return score
}