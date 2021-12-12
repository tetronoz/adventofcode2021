package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	part1()
}


func part1() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	graph := make(map[string][]string)
	var paths1 []string
	var paths2 []string

	
	for sc.Scan() {
		node := strings.Split(sc.Text(), "-")
		graph[node[0]] = append(graph[node[0]], node[1])
		graph[node[1]] = append(graph[node[1]], node[0])
	}

	for _, node := range graph["start"] {
		seen := make(map[string]bool)
		seen["start"] = true
		findEndOnce([]string{"start"}, &paths1, node, graph, seen)
	}

	fmt.Println(len(paths1))


	for _, node := range graph["start"] {
		seen := make(map[string]int)
		seen["start"] = 1
		for _, s := range findEndTwice([]string{"start"}, node, graph, seen) {
			paths2 = append(paths2, s)
		}
	}

	fmt.Println(len(paths2))
}

func findEndOnce(current_p []string, paths *[]string, node string, graph map[string][]string, seen map[string]bool) {
	if node == "end" {
		current_p = append(current_p, "end")
		*paths = append(*paths, strings.Join(current_p, ","))
		return
	}

	current_p = append(current_p, node)
	_, found := seen[node]

	if strings.ToLower(node) == node && found {
		return
	} 
	
	seen[node] = true

	for _, n := range graph[node] {
		seen_copy := make(map[string]bool)
		for k,v := range seen {
			seen_copy[k] = v
		}
		findEndOnce(current_p, paths, n, graph, seen_copy)
	}

	return
}


func findEndTwice(current_p []string, node string, graph map[string][]string, seen map[string]int) []string {
	if node == "end" {
		current_p = append(current_p, "end")
		return []string{strings.Join(current_p, ",")}
	}

	current_p = append(current_p, node)

	if node == "start" {
		return []string{}
	}

	_, found := seen[node]
	
	var seen_twice bool
	
	for k := range seen {
		if seen[k] > 1 {
			seen_twice = true
		}
	}

	if strings.ToLower(node) == node && found && seen[node] > 0 && seen_twice {
		return []string{}
	} 
	
	if strings.ToLower(node) == node {
		seen[node]++
	}

	res := []string{}
	for _, n := range graph[node] {
		seen_copy := make(map[string]int)
		for k,v := range seen {
			seen_copy[k] = v
		}
		
		for _, s := range findEndTwice(current_p, n, graph, seen_copy) {
			res = append(res, s)
		}
	}
	return res
}