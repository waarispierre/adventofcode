package main

import (
	"adventofcode/shared/loaddata"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		Challenge()
	case "2":
		ChallengeTwo()
	default:
		fmt.Println("Options 1 or 2")
	}

	fmt.Println("Duration:", time.Since(start))
}

func Challenge() {
	fmt.Println("Starting day one part one")
	data, err := loaddata.ReadData("day01/dayone.txt")
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return
	}

    bufferSize := 105
	lch := make(chan int, bufferSize) 
	rch := make(chan int, bufferSize)
	resCh := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	// Start the calculation goroutine
	go func() {
		defer wg.Done()
		var totalDistance int
		for left := range lch {
			right := <-rch
			distance := left - right
			if distance < 0 {
				distance = -1 * distance
			}
			totalDistance = totalDistance + distance
		}
		resCh <- totalDistance
	}()

	// Start populating channels concurrently with calculation
	go createList(data, lch, rch)

	// Wait for calculation to complete
	wg.Wait()
	fmt.Println(<-resCh)
}

func ChallengeTwo() {
	fmt.Println("Starting day one part two")
	data, err := loaddata.ReadData("day01/dayone.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	frequencyMap := make(map[int]int)
	var row []int
	
	for i := 0; i < len(data); i++ {
		row = data[i]
		frequencyMap[row[1]]++
	}

	var leftValue int
	var total int
	
	for i := 0; i < len(data); i++ {
		row = data[i]
		leftValue = row[0]
		total += (frequencyMap[leftValue] * leftValue)
	}

	fmt.Println(total)
}

// ChallengeSimple does the same calculation as Challenge but without goroutines
func ChallengeSimple() {
	fmt.Println("Starting day one (simple version)")
	data, err := loaddata.ReadData("dayone.txt")
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return
	}

	// Create and populate left and right lists
	var leftList []int
	var rightList []int

	for _, row := range data {
		leftList = append(leftList, row[0])
		rightList = append(rightList, row[1])
	}

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate total distance
	var totalDistance int
	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	fmt.Printf("Total Distance: %d\n", totalDistance)
}

func createList(data [][]int, lch chan<- int, rch chan<- int) {
	var rTempList []int
	var lTempList []int

	for i := 0; i < len(data); i++ {
		row := data[i]
		lTempList = append(lTempList, row[0])
		rTempList = append(rTempList, row[1])
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Sort both lists in parallel
	go func() {
		defer wg.Done()
		sortList(&lTempList)
	}()

	go func() {
		defer wg.Done()
		sortList(&rTempList)
	}()

	wg.Wait() // Wait for sorting to complete

	// Now populate channels in parallel
	go func() {
		populateChannel(lch, lTempList)
		close(lch)
	}()

	go func() {
		populateChannel(rch, rTempList)
		close(rch)
	}()

	return
}

func populateChannel(ch chan<- int, list []int) {
	for i := 0; i < len(list); i++ {
        ch <- list[i]
	}
}

func sortList(list *[]int) {
	sort.Slice(*list, func(i, j int) bool {
		return (*list)[i] < (*list)[j]
	})
}

