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
	part2()
}

func part1() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()
	sc := bufio.NewScanner(fp)

	digits := make(map[int][]string)

	var counter int
	for sc.Scan() {
		signals := strings.Split(sc.Text(), "|")[1]

		for _, signal := range strings.Fields(signals) {
			r := []rune(signal)
			sort.Slice(r, func(i int, j int) bool { return r[i] < r[j]})
			switch len(signal) {
			case 2:
				// 1
				digits[1] = append(digits[1], string(r))
				counter++
			case 4:
				// 4
				counter++
			case 3:
				// 7
				digits[7] = append(digits[7], string(r))
				counter++
			case 7:
				//8
				counter++
			}

		}
	}

	fmt.Println(counter)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j]})

	return string(r)
}

func part2() {
	fp, _ := os.Open("input.txt")
	defer fp.Close()
	sc := bufio.NewScanner(fp)

	bigSum := 0

	for sc.Scan() {
		data := strings.Split(sc.Text(), "|")
		signals := data[0]
		//digits := data[1]

		lettersToDigit := make(map[string]int)
		digitsToLetters := make(map[int]string)

		for _, signal := range strings.Fields(signals) {
			switch len(signal) {
			case 2:
				lettersToDigit[signal] = 1
				digitsToLetters[1] = signal
			case 3:
				lettersToDigit[signal] = 7
				digitsToLetters[7] = signal
			case 4:
				lettersToDigit[signal] = 4
				digitsToLetters[4] = signal
			case 7:
				lettersToDigit[signal] = 8
				digitsToLetters[8] = signal
			}
		}

		// searchin for 3
		for _, signal := range strings.Fields(signals) {
			if len(signal) == 5 {
				matches := 0
				for _, ch := range strings.Split(digitsToLetters[1], "") {
					if strings.Contains(signal, ch) {
						matches++
					}

					if matches == 2 {
						lettersToDigit[signal] = 3
						digitsToLetters[3] = signal
					}
				}
			}
		}

		// searchin for 9
		for _, signal := range strings.Fields(signals) {
			if len(signal) == 6 {
				m := make(map[string]bool)
				for _, ch := range strings.Split(digitsToLetters[3], "") {
					m[ch] = true
				}

				for _, ch := range strings.Split(digitsToLetters[4], "") {
					m[ch] = true
				}

				matches := 0
				for _, ch := range strings.Split(signal, "") {
					_, ok := m[ch]
					if ok {
						matches += 1
					} else {
						break
					}
				}

				if matches == 6 {
					lettersToDigit[signal] = 9
					digitsToLetters[9] = signal
				}
			}
		}

		// searchin for 5
		for _, signal := range strings.Fields(signals) {
			_, ok := lettersToDigit[signal]
			if len(signal) == 5 &&  !ok {
				m := make(map[string]bool)
				for _, ch := range strings.Split(digitsToLetters[1], "") {
					m[ch] = true
				}

				for _, ch := range strings.Split(signal, "") {
					m[ch] = true
				}

				matches := 0
				for _, ch := range strings.Split(digitsToLetters[9], "") {
					_, ok := m[ch]
					if ok {
						matches += 1
					} else {
						break
					}
				}

				if matches == 6 {
					lettersToDigit[signal] = 5
					digitsToLetters[5] = signal
				}
			}
		}

		// searchin for 2
		for _, signal := range strings.Fields(signals) {
			_, ok := lettersToDigit[signal]
			if len(signal) == 5 &&  !ok {
				lettersToDigit[signal] = 2
				digitsToLetters[2] = signal
			}
		}

		// search 6
		for _, signal := range strings.Fields(signals) {
			_, ok := lettersToDigit[signal]
			if len(signal) == 6 &&  !ok {
				m := make(map[string]bool)
				for _, ch := range strings.Split(digitsToLetters[1], "") {
					m[ch] = true
				}

				for _, ch := range strings.Split(signal, "") {
					m[ch] = true
				}

				matches := 0
				for _, ch := range strings.Split(digitsToLetters[8], "") {
					_, ok := m[ch]
					if ok {
						matches += 1
					} else {
						break
					}
				}

				if matches == 7 {
					lettersToDigit[signal] = 6
					digitsToLetters[6] = signal
				}
			}
		}

		//search 0
		for _, signal := range strings.Fields(signals) {
			_, ok := lettersToDigit[signal]
			if len(signal) == 6 && !ok {
				lettersToDigit[signal] = 0
				digitsToLetters[0] = signal
			}
		}

		var res []string 

		for _, digit := range strings.Fields(data[1]) {
			for k, v := range lettersToDigit {
				if sortString(digit) == sortString(k) {
					res = append(res, strconv.Itoa(v))	
				}
			}
		}
		d, _ := strconv.Atoi(strings.Join(res, ""))
		bigSum += d 
	}

	fmt.Println(bigSum)

}