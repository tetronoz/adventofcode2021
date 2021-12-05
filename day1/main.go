package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gammazero/deque"
)

func main() {
	part1()
	part2()
}

func part1() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()

	var counter int

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		
		if num > prev {
			counter += 1
		}
		prev = num
	}
	fmt.Println(counter)
}

func part2() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer fp.Close()

	d := deque.New()
	
	var currentSum int
	var prevSum int
	var counter int

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to convert string to int %v", err)
		}
		
		if d.Len() < 3 {
			d.PushBack(num)
			currentSum += num
		} else {
			prevSum = currentSum
			leftVal := d.PopFront().(int)
			
			d.PushBack(num)

			currentSum += num - leftVal
			if currentSum > prevSum {
				counter += 1
			}
		}
	}

	fmt.Println(counter)
}