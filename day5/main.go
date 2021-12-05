package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func part1() int {
	fileName := "input.txt"
	fp, _ := os.Open(fileName)
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	var diagram = make(map[string]int)
	var counter int

	for sc.Scan() {
		var x1, y1, x2, y2 int
		fmt.Sscanf(sc.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		//x - col, y - row
		if x1 == x2 {
			col := x1
			startRow := getMin(y1, y2)
			endRow := getMax(y1, y2)

			for row := startRow; row <= endRow; row++ {
				diagram[fmt.Sprintf("%d,%d", col, row)]++
			}
		} else if y1 == y2 {
			row := y1
			startCol := getMin(x1, x2)
			endCol := getMax(x1, x2)

			for col := startCol; col <= endCol; col++ {
				diagram[fmt.Sprintf("%d,%d", col, row)]++
			}
		}
	}

	for _, val := range diagram {
		if val > 1 {
			counter++
		}
	}

	return counter
}


func direction(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return 1
	}
	return -1
}


func part2() int {
	fileName := "input.txt"
	fp, _ := os.Open(fileName)
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	var diagram = make(map[string]int)
	var counter int

	for sc.Scan() {
		var x1, y1, x2, y2 int
		fmt.Sscanf(sc.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		col := x1
		row := y1
		endCol := x2
		endRow := y2

		for {
			diagram[fmt.Sprintf("%d,%d", col, row)]++
			if col == endCol && row == endRow {
				break
			}
			col += direction(col, endCol)
			row += direction(row, endRow)
		}
	}

	for _, val := range diagram {
		if val > 1 {
			counter++
		}
	}

	return counter
}
