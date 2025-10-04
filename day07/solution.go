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
	if len(os.Args) < 2 {
		fmt.Println("Please provide day challenge number 1 or 2 as a parameter")
		return
	}

	switch os.Args[1]{
	case "1":
		fmt.Println("Run challenge one")
		challengeOne()
	case "2":
	 	fmt.Println("Run challenge two")
		challengeTwo()
	}
	fmt.Println("Duration:", time.Since(start))
}

func challengeTwo() {
	total, err := getTotalOfCalibration([]string{"+", "*", "||"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result:", total)
}

func challengeOne() {
	total, err := getTotalOfCalibration([]string{"+", "*"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result:", total)
}

func getTotalOfCalibration(operators []string) (total int, err error) {
	data, err := getData()
	if err != nil {
		return
	}
	total = 0
	sequenceMap := make(map[int][][]string)
	for _, values := range data {
		_, exists := sequenceMap[len(values)-1]
		if !exists {
			sequenceMap[len(values)-1] = getOptions(len(values)-1, operators)
		}
		if canBeTrue(values, sequenceMap[len(values)-1]) {
			total += values[0]
		}
	}
	return
}

func canBeTrue(values []int, sequences [][]string) bool {
	testValue := values[0]
	var parseErr error
	for _, seq := range sequences {
		var total int = 0
		for i, value := range values[1:] {
			if i == 0 {
				total = value
				continue
			}
			switch seq[i-1] {
			case "+":
				total += value
			case "*":
				total = total * value
			case "||":
				total, parseErr = strconv.Atoi(fmt.Sprintf("%d%d", total, value))
				if parseErr != nil {
					panic("Could not parse new value")
				}
			} 		
		}
		if total == testValue {
			return true
		}
	}
	return false
}

func getOptions(n int, operators []string) [][]string {
	var differentOptions [][]string
	var slots int = n - 1
	var backtrack func(current []string)
	backtrack = func(current []string) {
		if len(current) == slots {
			seq := make([]string, slots)
			copy(seq, current)
			differentOptions = append(differentOptions, seq)
			return
		}
		for _, operator := range operators {
			backtrack(append(current, operator))
		}
	}
	backtrack([]string{})
	return differentOptions
}

func getData() (data [][]int, err error) {
	
	file, err := os.Open("day07/dayseven.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			continue
		}
		
		splitOne := strings.Split(currentLine, ":")
		possibleAnswer, err1 := strconv.Atoi(splitOne[0])	
		if err1 != nil {
			err = err1
			return
		}
		
		splitTwo := strings.Split(strings.TrimSpace(splitOne[1]), " ")
		var intSlice []int = []int{possibleAnswer}
		for _, str := range splitTwo {
			integer, parseErr := strconv.Atoi(str)
			if parseErr != nil {
				err = parseErr
				return
			}
			intSlice = append(intSlice, integer)	
		}
	
		data = append(data, intSlice)
	}
	return
}
