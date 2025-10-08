package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Please parse an argument. Options are 1 or 2")
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
	fmt.Println("Starting challenge 1")
	data, err := getData("day07/daysevent.txt")
}
	
func challengeTwo() {
	fmt.Println("Starting challenge 2")
}

func getData(path string) (data [][]rune, err error) {
	
}
