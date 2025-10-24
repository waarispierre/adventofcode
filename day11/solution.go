package main

import (
	"fmt"
	"os"
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

	fmt.Println(data)
}

func loadData(s string) (data []string, err error) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	return
}

func challengeTwo() {
	panic("unimplemented")
}
