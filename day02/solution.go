package main

import (
	"adventofcode/shared/loaddata"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
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

func ChallengeOne() {
	data, err := loaddata.ReadData("day02/daytwo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowChannel := make(chan []int, int(len(data)/3))
	var validReports int32

	var wg sync.WaitGroup
	numberOfWorkers := 10

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go validateReport(rowChannel, &validReports, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(rowChannel)
		for i := 0; i < len(data); i++ {
			rowChannel <- data[i]
		}
	}()

	wg.Wait()

	// Read the final result
	fmt.Println("Valid reports:", validReports)
}

func ChallengeTwo() {
	data, err := loaddata.ReadData("day02/daytwo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowChannel := make(chan []int, int(len(data)/3))
	var validReports int32

	var wg sync.WaitGroup
	numberOfWorkers := 1

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go validateReportForChallengeTwo(rowChannel, &validReports, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(rowChannel)
		for i := 0; i < len(data); i++ {
			rowChannel <- data[i]
		}
	}()

	wg.Wait()

	// Read the final result
	fmt.Println("Valid reports:", validReports)
}

func validateReport(rowCh <-chan []int, validReports *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for row := range rowCh {
		if validateIncrements(row) {
			atomic.AddInt32(validReports, 1)
		}
	}
}

func validateReportForChallengeTwo(rowCh <-chan []int, validReports *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for row := range rowCh {
		if validateIncrements(row) {
			atomic.AddInt32(validReports, 1)
		} else {
			for i := 0; i < len(row); i++ {
				newRow := make([]int, 0, len(row)-1)
				newRow = append(newRow, row[:i]...)
				newRow = append(newRow, row[i+1:]...)
				if validateIncrements(newRow) {
					atomic.AddInt32(validReports, 1)
					break 
				}
			}
		}
	}
}

func validateIncrements(row []int) bool {
	absIncrement := 0
	var increasing bool
	for i := 0; i < len(row)-1; i++ {
		if i == 0 {
			increasing = row[0] < row[1]
		}
		if increasing != (row[i] < row[i+1]) {
			return false
		}
		absIncrement = abs(row[i] - row[i+1])
		oneOrMoreRule := absIncrement >= 1
		threeOrLessRule := absIncrement <= 3
		if !oneOrMoreRule || !threeOrLessRule {
			return false
		}
	}
	return true
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
