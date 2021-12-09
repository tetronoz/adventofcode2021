package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var points [][]int

	for sc.Scan() {
		var row []int
		for _,c := range strings.Split(sc.Text(), "") {
			d, _ := strconv.Atoi(c)
			row = append(row, d)
		}

		points = append(points, row)
	}

	var heights []int
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {

			// look up
			if i - 1 >= 0 {
				if points[i][j] >= points[i - 1][j] {
					continue
				}
			}
			// look down
			if i + 1 < len(points) {
				if points[i][j] >= points[i + 1][j] {
					continue
				}	
			}
			// look left
			if j - 1 >= 0 {
				if points[i][j] >= points[i][j - 1] {
					continue
				}
			}
			// look right
			if j + 1 < len(points[0]) {
				if points[i][j] >= points[i][j + 1] {
					continue
				}
			}
			
			heights = append(heights, points[i][j])
		}
	}
	var riskLevel int
	for _, v := range heights {
		riskLevel += v + 1
	}

	fmt.Println(riskLevel)

	var sizes []int

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			if points[i][j] == 9 {
				continue
			}
			var elem [2]int
			elem[0] = i
			elem[1] = j

			// look up
			if i - 1 >= 0 {
				if points[i][j] >= points[i - 1][j] {
					continue
				}
			}
			// look down
			if i + 1 < len(points) {
				if points[i][j] >= points[i + 1][j] {
					continue
				}	
			}
			// look left
			if j - 1 >= 0 {
				if points[i][j] >= points[i][j - 1] {
					continue
				}
			}
			// look right
			if j + 1 < len(points[0]) {
				if points[i][j] >= points[i][j + 1] {
					continue
				}
			}

			var size int
			if points[i][j] != -1 {
				size += countBasinSize(points, i, j)
			}

			sizes = append(sizes, size)
		}
	}

	sort.Ints(sizes)
	fmt.Println(sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3])
}

func countBasinSize (points [][]int, i,j int) int {
	var queue [][2]int
	var size int
	
	queue = append(queue, [2]int{i, j})
	
	for len(queue) != 0 {
		elem := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if points[elem[0]][elem[1]] != -1 {
			points[elem[0]][elem[1]] = -1
			size++

			i = elem[0]
			j = elem[1] 
		
			// look up
			if i - 1 >= 0 {
				if  points[i][j] < points[i - 1][j] && points[i - 1][j] != 9 {
					queue = append(queue, [2]int{i-1,j})
				} 
			}
			// look down
			if i + 1 < len(points) {
				if  points[i][j] < points[i + 1][j] && points[i + 1][j] != 9 {
					queue = append(queue, [2]int{i+1, j})
				}	
			}
			// look left
			if j - 1 >= 0 {
				if  points[i][j] < points[i][j - 1] && points[i][j - 1] != 9 {
					queue = append(queue, [2]int{i, j-1})
					
				}
			}
			// look right
			if j + 1 < len(points[0]) {
				if  points[i][j] < points[i][j + 1] && points[i][j + 1] != 9 {
					queue = append(queue, [2]int{i, j+1})
				}
			}
		}
	}
	return size
}
