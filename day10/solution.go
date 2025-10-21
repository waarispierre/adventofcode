package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting day 10")

	if len(os.Args) < 2 {
		fmt.Println(fmt.Errorf("provide at least one parameter, options are 1 or 2"))
		return
	}

	switch os.Args[1] {
	case "1":
		fmt.Println("Challenge 1")
		challengeOne()
	case "2":
		fmt.Println("Challenge 2")
		challengeTwo()
	}
}

func challengeTwo() {
	panic("unimplemented")
}

func challengeOne() {
	data, err := loadData("day10/dayten.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range data {
		fmt.Println(d)
	}
}

func loadData(s string) (data [][]int, err error) {
	file, err := os.Open(s)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currLine := ""
	convertedInt := 0
	for scanner.Scan() {
		currLine = scanner.Text()
		var row []int
		for _, str := range currLine {
			convertedInt, err = strconv.Atoi(string(str))
			if err != nil {
				return
			}
			row = append(row, convertedInt)
		}
		data = append(data, row)
	}

	return
}
