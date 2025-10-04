package main

import (
	"adventofcode/shared/loaddata"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		ChallengeOne()
	case "2":
		ChallengeTwo()
	default:
		fmt.Println("Options 1 or 2")
	}

	fmt.Println("Duration:", time.Since(start))
}

func ChallengeTwo() {
	data, err := loaddata.ReadDataInRuneMatrix("day04/dayfour.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for i, row := range data {
		for j, r := range row {
			if r == 'A' && i >= 1 && j >= 1 && i < len(data)-1 && j < len(row)-1 {
				if !(data[i-1][j-1] == 'M' && data[i+1][j+1] == 'S') && !(data[i-1][j-1] == 'S' && data[i+1][j+1] == 'M') {
					continue
				}
				if (data[i+1][j-1] == 'M' && data[i-1][j+1] == 'S') || (data[i+1][j-1] == 'S' && data[i-1][j+1] == 'M') {
					total += 1
				}
			}
		}
	}
	fmt.Println(total)
}

func ChallengeOne() {
	data, err := loaddata.ReadDataInRuneMatrix("day04/dayfour.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var total int
	for i, row := range data {
		for j, r := range row {
			if r == 'X' {
				if j < len(row)-3 && data[i][j+1] == 'M' && data[i][j+2] == 'A' && data[i][j+3] == 'S' {
					total += 1
				}
				if i < len(data)-3 && data[i+1][j] == 'M' && data[i+2][j] == 'A' && data[i+3][j] == 'S' {
					total += 1
				}
				if i > 2 && data[i-1][j] == 'M' && data[i-2][j] == 'A' && data[i-3][j] == 'S' {
					total += 1
				}
				if j > 2 && data[i][j-1] == 'M' && data[i][j-2] == 'A' && data[i][j-3] == 'S' {
					total += 1
				}
				if j > 2 && i > 2 && data[i-1][j-1] == 'M' && data[i-2][j-2] == 'A' && data[i-3][j-3] == 'S' {
					total += 1
				}
				if j < len(row)-3 && i < len(data)-3 && data[i+1][j+1] == 'M' && data[i+2][j+2] == 'A' && data[i+3][j+3] == 'S' {
					total += 1
				}
				if j > 2 && i < len(data)-3 && data[i+1][j-1] == 'M' && data[i+2][j-2] == 'A' && data[i+3][j-3] == 'S' {
					total += 1
				}
				if j < len(row)-3 && i > 2 && data[i-1][j+1] == 'M' && data[i-2][j+2] == 'A' && data[i-3][j+3] == 'S' {
					total += 1
				}
			}
		}
	}
	fmt.Println(total)
}
