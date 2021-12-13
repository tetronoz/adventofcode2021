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
}

func part1() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	dots := make(map[[2]int]bool)

	var max_x, max_y int
	var foldings []string

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		
		if strings.HasPrefix(line, "fold along") {
			foldings = append(foldings, strings.Split(line, " ")[2])
		} else {

			var x, y int
			fmt.Sscanf(line, "%d,%d", &x, &y)
			if x > max_x {
				max_x = x
			}

			if y > max_y {
				max_y = y
			}
			k := [2]int{y, x}
			dots[k] = true
		}
	}

	var paper [][]string

	for i := 0; i <= max_y; i++ {
		var row []string
		for j := 0; j <= max_x; j++ {
			dot := [2]int{i,j}
			if _, ok := dots[dot]; ok {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		paper = append(paper, row)
	}

	var foldNum int
	for _, folding := range foldings {
		paper = runFolding(paper, folding)
		foldNum++
		if foldNum == 1 {
			fmt.Println(countVisible(paper))
		}
	}
	for _, row := range paper {
		fmt.Println(row)
	} 
}

func runFolding(dots [][]string, inst string) [][]string {

	s := strings.Split(inst, "=")

	switch s[0] {
	case "x":
			line,_ := strconv.Atoi(s[1])
			dots = foldVertically(dots, line)
	case "y":
		line,_ := strconv.Atoi(s[1])
			dots = foldHorizontal(dots, line)
	}

	return dots
}

func foldHorizontal(dots [][]string, line int) [][]string {

	rows := len(dots) - 1

	for row := rows; rows - row < line; row-- {
		for col :=0; col < len(dots[0]); col++ {
			if dots[row][col] == "#" {
				dots[rows - row][col] = dots[row][col]
			}
		}
	}

	return dots[:line]
}

func foldVertically(dots [][]string, line int) [][]string {

	cols := len(dots[0]) - 1
	for row := 0; row < len(dots); row++ {
		for col := cols; col > line; col-- {
			if dots[row][col] == "#" {
				dots[row][cols - col] = dots[row][col]
			}
		}
	}

	var res [][]string

	for _, r := range dots {
		res = append(res, r[:line])
	}

	return res
}

func countVisible(dots [][]string) int {
	var counter int
	for _, row := range dots {
		for _, ch := range row {
			if ch == "#" {
				counter++
			}
		}
	}

	return counter
}