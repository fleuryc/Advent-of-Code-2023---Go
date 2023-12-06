package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	strToInt := map[string]int{
		"0":     0,
		"zero":  0,
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}

	file, err := os.Open("data/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile("([0-9]|zero|one|two|three|four|five|six|seven|eight|nine)")
		matches_start := re.FindAllString(scanner.Text(), -1)
		first, _ := strToInt[matches_start[0]]
		last, _ := strToInt[matches_start[len(matches_start)-1]]
		indexes := re.FindAllStringIndex(scanner.Text(), -1)
		matches_last := re.FindAllString(scanner.Text()[indexes[len(indexes)-1][0]+1:], -1)
		if len(matches_last) > 0 {
			last, _ = strToInt[matches_last[0]]
		}
		add := first*10 + last
		fmt.Println(add)

		sum += first*10 + last
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
