package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)


func main() {
	part1()
	part2()
}


func part1() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	sc.Scan()
	tpl := sc.Text()

	rules := make(map[string]string)

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		rules[fields[0]] = fields[2]
	}


	for step := 1; step < 11; step++ {
		tpl = applyRules(tpl, rules)
	}

	counts := make(map[string]int)
	for _, el := range tpl {
		counts[string(el)]++
	}

	max := math.MinInt
	min := math.MaxInt

	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	fmt.Println(max - min)
}

func applyRules(t string, r map[string]string) string {

	var res []string

	for i := 0; i < len(t) - 1; i++ {
		p := string(t[i:i+2])
		if _, ok := r[p]; ok {
			res = append(res, string(p[0]))
			res = append(res, r[p])
		} else if !ok {
			res = append(res, string(p[0]))
			res = append(res, string(p[1]))
		}
		if i == len(t) - 2 { 
			res = append(res, string(p[1]))
		}
	}
	return strings.Join(res, "")
}

func part2() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	sc.Scan()
	tpl := sc.Text()

	rules := make(map[string]string)
	pairs := make(map[string]int)
	counts := make(map[string]int)

	for i := 0; i < len(tpl) -1; i++ {
		pairs[tpl[i:i+2]]++
	}
	pairs[tpl[len(tpl) - 1:]]++

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		rules[fields[0]] = fields[2]
	}

	for i := 0; i < 40; i++ {
		new_pairs := make(map[string]int)

		for k, v := range pairs {
			if _, ok := rules[k]; ok {
				new_pairs[k[0:1] + rules[k]] += v
				new_pairs[rules[k] + k[1:2]] += v
			} else {
				new_pairs[k] += v
			}
		}

		pairs = new_pairs
	}

	max := math.MinInt
	min := math.MaxInt

	for k, v := range pairs {
		counts[k[0:1]] += v
	}

	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	fmt.Println(max - min)
}

