package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func main() {
	part1()
	part2()
}


func part1() {
	fp, _ := os.Open("input.txt")
	sc := bufio.NewScanner(fp)

	sc.Scan()
	
	rawData := strings.Split(sc.Text(), ",")

	var positions []int
	minVal := math.MaxInt
	maxVal := math.MinInt

	for _, data := range rawData {
		pos, _ := strconv.Atoi(data)
		if pos > maxVal {
			maxVal = pos
		}

		if pos < minVal {
			minVal = pos
		}
		positions = append(positions, pos)
	}

	var bestPosition int
	
	bestPosition = -1
	minFuel := math.MaxFloat64

	for i := minVal; i <= maxVal; i++ {
		var fuel float64
		for _, val := range positions {
			fuel += math.Abs(float64(i - val))
			//cost := sumOfN(fuel)
			//fuel += cost
		}
		if fuel < minFuel {
			minFuel = fuel
			bestPosition = i
		}
	}

	fmt.Println(bestPosition, minFuel)
}

func part2() {
	fp, _ := os.Open("input.txt")
	sc := bufio.NewScanner(fp)

	sc.Scan()
	
	rawData := strings.Split(sc.Text(), ",")

	var positions []int
	minVal := math.MaxInt
	maxVal := math.MinInt

	for _, data := range rawData {
		pos, _ := strconv.Atoi(data)
		if pos > maxVal {
			maxVal = pos
		}

		if pos < minVal {
			minVal = pos
		}
		positions = append(positions, pos)
	}

	var bestPosition int
	
	bestPosition = -1
	minFuel := math.MaxFloat64

	for i := minVal; i <= maxVal; i++ {
		var fuel float64
		for _, val := range positions {
			distance := math.Abs(float64(i - val))
			cost := sumOfN(distance)
			fuel += cost
		}
		if fuel < minFuel {
			minFuel = fuel
			bestPosition = i
		}
	}

	fmt.Println(bestPosition, int(minFuel))
}

func sumOfN(n float64) float64{
	return (n * (n + 1)) / 2
}