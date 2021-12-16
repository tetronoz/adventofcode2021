package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"container/heap"
)

type Coord struct {
	row int
	col int
}

type Node struct {
	pos   Coord
	dist  int
	val   int
	prev  *Node
	isInf bool
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(n interface{}) { *pq = append(*pq, n.(*Node)) }

func (pq *PriorityQueue) Pop() interface{} {
	t := *pq
	var n interface{}

	n, *pq = t[len(t)-1], t[:len(t)-1]
	return n
}

func main() {
	part1()
	part2()
}

func part1() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	nodeMap := make(map[Coord]*Node)
	
	row := 0
	col := 0
	for sc.Scan() {
		col = 0
		for _, risk := range strings.Split(sc.Text(), "") {
			node := Node{}
			node.pos.row = row
			node.isInf = true
			node.pos.col = col
			riskValue, _ := strconv.Atoi(risk)
			node.val = riskValue
			nodeMap[Coord{row, col}] = &node
			col++
		}
		//col--
		row++
	}

	pq := make(PriorityQueue, 1)

	row--
	col--
	var start, end *Node
	start = nodeMap[Coord{0,0}]
	end = nodeMap[Coord{row, col}]

	start.isInf = false
	start.dist = 0
	
	pq[0] = start

	heap.Init(&pq)

	visited := make(map[*Node]bool)
	visited[start] = false

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if node == end {
			fmt.Println(node.dist)
			return
		}

		visited[node] = true

		up := Coord{node.pos.row-1, node.pos.col}
		
		down := Coord{node.pos.row+1, node.pos.col}
		
		right := Coord{node.pos.row, node.pos.col+1}
		
		left := Coord{node.pos.row, node.pos.col-1}

		for _, pos := range []Coord{up, down, right, left} {
			if _, ok := nodeMap[pos]; ok {
				neighbourNode := nodeMap[pos]
				if visited[neighbourNode] {
					continue
				}

				newDist := neighbourNode.val + node.dist
				if neighbourNode.isInf || newDist < neighbourNode.dist {
					neighbourNode.dist = newDist
					neighbourNode.prev = node
					neighbourNode.isInf = false
					heap.Push(&pq, neighbourNode)
				}
			}
		}
	}
}

func part2() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	nodeMap := make(map[Coord]*Node)
	
	row := 0
	col := 0
	for sc.Scan() {
		col = 0
		for _, risk := range strings.Split(sc.Text(), "") {
			node := Node{}
			node.pos.row = row
			node.isInf = true
			node.pos.col = col
			riskValue, _ := strconv.Atoi(risk)
			node.val = riskValue
			nodeMap[Coord{row, col}] = &node
			col++
		}
		row++
	}

	// expand to right

	var end_col_idx int
	start_col := col
	for i := 0; i < row; i++ {
		end_col := start_col + col

		for s := 0; s < 4; s++ {
			end_col = start_col + col
			for j := start_col; j < end_col; j++ {
				v := nodeMap[Coord{i, j - col}].val
				v++
				if v > 9 {
					v = 1
				}

				node := Node{}
				node.pos.row = i
				node.isInf = true
				node.pos.col = j
				node.val = v
				nodeMap[Coord{i, j}] = &node
			}
			end_col_idx = end_col - 1
			start_col += col
		}
		start_col = col
	}

	// expand down

	start_row := row
	var end_row_idx int
	for s := 0; s < 4; s++ {
		end_row := start_row + row
		
		for i := start_row; i < end_row ; i++ {
			for j := 0; j <= end_col_idx; j++ {
				v := nodeMap[Coord{i - row, j}].val
				v++
				if v > 9 {
					v = 1
				}
				node := Node{}
				node.pos.row = i
				node.isInf = true
				node.pos.col = j
				node.val = v
				nodeMap[Coord{i, j}] = &node
			}
		}
		end_row_idx = end_row - 1
		start_row += row
	}

	pq := make(PriorityQueue, 1)

	var start, end *Node
	start = nodeMap[Coord{0,0}]
	end = nodeMap[Coord{end_row_idx, end_col_idx}]

	start.isInf = false
	start.dist = 0

	pq[0] = start

	heap.Init(&pq)

	visited := make(map[*Node]bool)
	visited[start] = false

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if node == end {
			fmt.Println(node.dist)
			return
		}

		visited[node] = true

		up := Coord{node.pos.row-1, node.pos.col}
		
		down := Coord{node.pos.row+1, node.pos.col}
		
		right := Coord{node.pos.row, node.pos.col+1}
		
		left := Coord{node.pos.row, node.pos.col-1}

		for _, pos := range []Coord{up, down, right, left} {
			if _, ok := nodeMap[pos]; ok {
				neighbourNode := nodeMap[pos]
				if visited[neighbourNode] {
					continue
				}

				newDist := neighbourNode.val + node.dist
				if neighbourNode.isInf || newDist < neighbourNode.dist {
					neighbourNode.dist = newDist
					neighbourNode.prev = node
					neighbourNode.isInf = false
					heap.Push(&pq, neighbourNode)
				}
			}
		}
	}
}
