package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	matrix [][]int
	numbers map[int][]int
	marked map[int]bool
	winner bool
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func NewBoard(input [][]string) Board {
	board := Board{}
	board.numbers = make(map[int][]int)
	board.marked = make(map[int]bool)

	for row := 0; row < len(input); row++ {
		rowSum := 0
		var tempRow []int
		for col := 0; col < len(input[0]); col++ {
			digit, _ := strconv.Atoi(input[row][col])
			board.numbers[digit] = append(board.numbers[digit], row)
			board.numbers[digit] = append(board.numbers[digit], col)
			board.marked[digit] = false
			rowSum += digit
			tempRow = append(tempRow, digit)
		}
		tempRow = append(tempRow, rowSum)
		board.matrix = append(board.matrix, tempRow)
	}
	var colSumRow []int
	for col := 0; col < 5; col++ {
		colSum := 0
		for row :=0; row < 5; row++ {
			colSum += board.matrix[row][col]
		}
		colSumRow = append(colSumRow, colSum)
	}
	board.matrix = append(board.matrix, colSumRow)
	board.winner = false
	return board
}

func part1() int {
	fileName := "input.txt"
	fp, _ := os.Open(fileName)
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	sc.Scan()
	drawNumbers := strings.Split(sc.Text(), ",")
	
	var boards []Board

	var input [][]string

	for sc.Scan() {
		if sc.Text() == "" {
			if len(input) > 0 {
				board := NewBoard(input)
				boards = append(boards, board)
			}
			input = nil
			continue
		} else {
			input = append(input, strings.Fields(sc.Text()))
		}
	}
	board := NewBoard(input)
	boards = append(boards, board)

	for _, drawNumber := range drawNumbers {
		number, _ := strconv.Atoi(drawNumber)
		for i := 0; i < len(boards); i++ {
			b := &boards[i]
			if _, ok := b.numbers[number]; ok {
				r := b.numbers[number][0]
				c := b.numbers[number][1]
				b.marked[number] = true
				b.matrix[5][c] -= number
				b.matrix[r][5] -= number

				if b.matrix[5][c] == 0 || b.matrix[r][5] == 0 {
					unmarkedSum := 0
					for val := range b.marked {
						if b.marked[val] == false {
							unmarkedSum += val
						}
					}
					return (number * unmarkedSum)
				}
			} 
		} 
	}

	return 0
}


func part2() int {
	fileName := "input.txt"
	fp, _ := os.Open(fileName)
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	sc.Scan()
	drawNumbers := strings.Split(sc.Text(), ",")
	
	var boards []Board

	var input [][]string

	for sc.Scan() {
		if sc.Text() == "" {
			if len(input) > 0 {
				board := NewBoard(input)
				boards = append(boards, board)
			}
			input = nil
			continue
		} else {
			input = append(input, strings.Fields(sc.Text()))
		}
	}
	board := NewBoard(input)
	boards = append(boards, board)
	winners := 0

	for _, drawNumber := range drawNumbers {
		number, _ := strconv.Atoi(drawNumber)
		for i := 0; i < len(boards); i++ {
			b := &boards[i]
			if !b.winner {
				if _, ok := b.numbers[number]; ok {
					if !b.marked[number] {
						r := b.numbers[number][0]
						c := b.numbers[number][1]
						b.marked[number] = true
						b.matrix[5][c] -= number
						b.matrix[r][5] -= number

						if b.matrix[5][c] == 0 || b.matrix[r][5] == 0 {
							b.winner = true
							winners++
							if winners == len(boards) {
								unmarkedSum := 0
								for val := range b.marked {
									if b.marked[val] == false {
										unmarkedSum += val
									}
								}
								return (number * unmarkedSum)
							}
						}
					}
				} 
			}
		} 
	}

	return 0
}