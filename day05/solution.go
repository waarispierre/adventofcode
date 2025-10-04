package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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
	rules := make(map[int][]int)
	updates, err := getdata(&rules)
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int

	var updatedPageOrder []int
	for _, update := range updates {
		updatedPageOrder = update.pages
		invalidUpdate := true
		for invalidUpdate {
			invalidUpdate = !checkIfOrderIsValid(updatedPageOrder, rules)
			if !invalidUpdate {
				continue
			}
			for i, value := range updatedPageOrder {
				if i == 0 {
					continue
				}
				rule := rules[value]
				valuesBefore := update.pages[:i]
				for _, v := range valuesBefore {
					if violateRule(rule, v) {
						indexOfV := slices.Index(updatedPageOrder, v)
						updatedPageOrder = swapValuesInSliceByIndex(updatedPageOrder, indexOfV, i)
					} 
				}
			}

			invalidUpdate = !checkIfOrderIsValid(updatedPageOrder, rules)
			if !invalidUpdate {
				total += updatedPageOrder[(len(update.pages)-1)/2]
			}
		}
	}
	fmt.Println("Answer:", total)
}

func swapValuesInSliceByIndex(updatedPageOrder []int, indexOfV, i int) []int {
	firstValue := updatedPageOrder[indexOfV]
	secondValue := updatedPageOrder[i]
	updatedPageOrder[indexOfV] = secondValue
	updatedPageOrder[i] = firstValue
	return updatedPageOrder
}

func ChallengeOne() {
	rules := make(map[int][]int)
	updates, err := getdata(&rules)
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for _, update := range updates {
		if checkIfOrderIsValid(update.pages, rules) {
			total += update.pages[(len(update.pages)-1)/2]
		}
	}
	fmt.Println("Answer:", total)
}

func checkIfOrderIsValid(pageOrders []int, rules map[int][]int) bool {
	isValid := true
	for i, value := range pageOrders {
		if i == 0 && !isValid{
			continue
		}
		rule := rules[value]
		valuesBefore := pageOrders[:i]
		for _, v := range valuesBefore {
			if violateRule(rule, v) {
				isValid = false
				continue
			}
		}
	}
	return isValid
}

func violateRule(slice []int, search int) bool {
	for _, v := range slice {
		if search == v {
			return true
		}
	}
	return false
}

type update struct {
	pages []int
}

func getdata(rules *map[int][]int) (updates []update, err error){
	file, err := os.Open("day05/dayfive.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := ""
	readingRules := true
	var ruleSplit []string
	for scanner.Scan() {
		currentLine = scanner.Text()

		if currentLine == "" {
			readingRules = false
			continue
		}

		if readingRules {
			ruleSplit = strings.Split(currentLine, "|")
			firstValue, err1 := strconv.Atoi(ruleSplit[0])
			if err1 != nil {
				err = err1
				return
			}
			secondValue, err2 := strconv.Atoi(ruleSplit[1])
			if err2 != nil {
				err = err2
				return
			}
			(*rules)[firstValue] = append((*rules)[firstValue], secondValue)
		} else {
			pages := strings.Split(currentLine, ",")
			var pageSlice []int
			pageNumber := 0
			for _, page := range pages {
				pageNumber, _ = strconv.Atoi(page)
				pageSlice = append(pageSlice, pageNumber)
			}
			updates = append(
				updates, 
				update {
					pages: pageSlice,
				},
			)
		}
	}
	return
}
