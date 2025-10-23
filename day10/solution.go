package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	routes := make(map[string]map[string]int)
	for row, d := range data {
		for col, pos := range d {
			if pos == 0 {
				routes[fmt.Sprintf("%v:%v", row, col)] = make(map[string]int)
			}
		}
	}

	var recursive func(origanalRow, origanalCol, row, col int)
	recursive = func(origanalRow, origanalCol, row, col int) {
		elev := data[row][col]
		if elev == 9 {
			routes[fmt.Sprintf("%v:%v", origanalRow, origanalCol)][fmt.Sprintf("%v:%v", row, col)] = 0
			return
		}

		stopRecursion := true
		//left
		if col > 0 && data[row][col-1]-elev == 1 {
			recursive(origanalRow, origanalCol, row, col-1)
			stopRecursion = false
		}
		//right
		if col < len(data[row])-1 && data[row][col+1]-elev == 1 {
			recursive(origanalRow, origanalCol, row, col+1)
			stopRecursion = false
		}
		//up
		if row > 0 && data[row-1][col]-elev == 1 {
			recursive(origanalRow, origanalCol, row-1, col)
			stopRecursion = false
		}
		//down
		if row < len(data)-1 && data[row+1][col]-elev == 1 {
			recursive(origanalRow, origanalCol, row+1, col)
			stopRecursion = false
		}
		if stopRecursion {
			return
		}
	}

	for k := range routes {
		parts := strings.Split(k, ":")
		row, _ := strconv.Atoi(parts[0])
		col, _ := strconv.Atoi(parts[1])
		recursive(row, col, row, col)
	}

	total := 0
	for _, v := range routes {
		total += len(v)
	}

	fmt.Println("Result:", total)
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
