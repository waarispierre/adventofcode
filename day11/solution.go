package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Start day 11")
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument, options are 1 or 2")
		return
	}

	switch os.Args[1] {
	case "1":
		challengeOne()
	case "2":
		challengeTwo()
	}
	fmt.Println("Duration:", time.Since(start))
}

func challengeOne() {
	data, err := loadData("day11/dayeleven.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	blinks := 25
	result := getNumberOfStones(blinks, data)
	fmt.Println("Result:", len(result))
}

func getNumberOfStones(blinks int, data []string) []string {
	var updatedSlice []string
	for range blinks {
		updatedSlice = make([]string, len(data))
		copy(updatedSlice, data)
		data = []string{}
		for _, stone := range updatedSlice {
			if stone == "0" {
				stone = "1"
				data = append(data, stone)
			} else if len(stone)%2 == 0 {
				n := len(stone) / 2
				stoneRange := strings.Split(stone, "")
				data = append(data, strings.Join(stoneRange[0:n], ""))
				num, _ := strconv.Atoi(strings.Join(stoneRange[n:], ""))
				data = append(data, fmt.Sprintf("%v", num))
			} else {
				number, _ := strconv.Atoi(stone)
				number = number * 2024
				stone = fmt.Sprintf("%v", number)
				data = append(data, stone)
			}
		}
	}
	updatedSlice = make([]string, len(data))
	copy(updatedSlice, data)
	return updatedSlice
}

func loadData(s string) (data []string, err error) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			continue
		}

		parts := strings.SplitSeq(currentLine, " ")
		for part := range parts {
			data = append(data, part)
		}
	}
	return
}

func challengeTwo() {
	data, err := loadData("day11/dayeleven.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	blinks := 75
	splitData := make([][]string, len(data))
	for i, d := range data {
		splitData[i] = append(splitData[i], d)
	}

	total := 0
	for i, d := range splitData {
		fmt.Println(i)
		result := getNumberOfStones(blinks, d)
		total += len(result)
	}
	fmt.Println("Result:", total)
}
