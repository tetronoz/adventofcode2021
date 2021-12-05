package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth int
	Aim int
}

func main() {
	part1()
	part2()
}

func part1() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	defer fp.Close()
	if err != nil {
		log.Fatalln(err)
	}
	
	pos := Position{Horizontal: 0, Depth: 0}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		unit, _ := strconv.Atoi(splitted[1])
		switch splitted[0] {
		case "forward":
			pos.Horizontal += unit
		case "up":
			pos.Depth -= unit
		case "down":
			pos.Depth += unit
		}
	}

	fmt.Println(pos.Horizontal * pos.Depth)
}

func part2() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	defer fp.Close()
	if err != nil {
		log.Fatalln(err)
	}
	
	pos := Position{Horizontal: 0, Depth: 0, Aim: 0}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		unit, _ := strconv.Atoi(splitted[1])
		switch splitted[0] {
		case "forward":
			pos.Horizontal += unit
			pos.Depth += pos.Aim * unit
		case "up":
			pos.Aim -= unit
		case "down":
			pos.Aim += unit
		}
	}

	fmt.Println(pos.Horizontal * pos.Depth)
}