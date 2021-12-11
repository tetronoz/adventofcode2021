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

	var energy [][]int

	for sc.Scan() {
		var row []int
		for _, c := range strings.Split(sc.Text(), "") {
			d, _ := strconv.Atoi(c)
			row = append(row, d)
		}
		energy = append(energy, row)
	}


	var totalFlashes int
	var allFlashed bool
	var step int

	//for step := 0; step < 100; step++ {
	for !allFlashed {
		step++
		// make a step
		for i := 0; i < len(energy); i++ {
			for j := 0; j < len(energy[0]); j++ {
				energy[i][j]++
			}
		}

		// check for > 9
		var counter int
		for i := 0; i < len(energy); i++ {
			for j := 0; j < len(energy[0]); j++ {
				if energy[i][j] > 9 {
					// check neighbours
					counter += flashNeighbours(energy, i, j)
				}
			}
		}

		totalFlashes += counter

		if step == 100 {
			fmt.Println(totalFlashes)
		}

		if counter == 100 {
			allFlashed = true
		}

	}
	fmt.Println(step)
}

func flashNeighbours(e [][]int, i, j int) int {
	var queue [][]int
	queue = append(queue, []int{i, j})

	var flashCount int

	dirs := [][]int{ {-1,0}, {1,0}, {0,1}, {0, -1}, {-1, 1}, {-1, -1}, {1, -1}, {1,1} }

	for len(queue) > 0 {
		el := queue[len(queue)-1]
		i := el[0]
		j := el[1]

		queue = queue[:len(queue)-1]

		if e[i][j] > 9 {
			flashCount++
			e[i][j] = 0
			for _, dir := range dirs {
				row := i + dir[0]
				col := j + dir[1]

				if 0 <= col && col < len(e[0]) && 0 <= row && row < len(e) {
					if e[row][col] != 0 {
						e[row][col]++

						if e[row][col] > 9 {
							queue = append(queue, []int{row, col})
						}
					}
				}
			}
		}
	}

	return flashCount
}