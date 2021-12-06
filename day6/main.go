package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()
	
	sc := bufio.NewScanner(fp)
	sc.Scan()

	data := strings.Split(sc.Text(), ",")

	var ages [9]int
	for _, age := range data {
		idx_age, _ := strconv.Atoi(age)
		ages[idx_age]++
	}

	for i := 0; i < 256; i++ {
		newborns := ages[0]
		for idx := 0; idx < len(ages) - 1; idx++ {
			ages[idx] = ages[idx+1]
		}
		ages[8] = newborns
		ages[6] += newborns
		if i == 79 {
			getSum(ages[:])
		}
		
	}

	getSum(ages[:])
}	

func getSum(a []int) {
	var sum int
	for _, v := range a {
		sum += v
	}

	fmt.Println(sum)
}