package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("data/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile("[0-9]")
		ints := re.FindAllString(scanner.Text(), -1)
		first, _ := strconv.Atoi(ints[0])
		last, _ := strconv.Atoi(ints[len(ints)-1])
		sum += first*10 + last
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
