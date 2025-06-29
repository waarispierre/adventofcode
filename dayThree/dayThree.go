package dayThree

import (
	"fmt"
	"unicode"
	"strconv"
	"github.com/waarispierre/adventOfCode/loadData"
)

func Challenge() {
	x, err := loadData.ReadDataInString("inputThree.txt")

	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	result := partOne(x)
	fmt.Printf("The result is: %v\n", result)
}

func partOne(x string) int {
	var mulResults [][]int  
	for i := 0; i < len(x) - 3; i++ {
		if x[i:i+4] == "mul(" {
			start := i + 4
			step := start

			for step < len(x) && unicode.IsDigit(rune(x[step])) {
				step++
			}
			if step == start || step >= len(x) || x[step] != ',' {
				continue
			}

			firstNumber, err := strconv.Atoi(x[start:step])
			if err != nil {
				fmt.Errorf("error parsing '%s' as integer: %v", x[start:step], err)
            }

			step++

			secondStart := step
			for step < len(x) && unicode.IsDigit(rune(x[step])) {
				step++
			}
			if step == secondStart || step >= len(x) || x[step] != ')' {
				continue
			}
			secondNumber, err := strconv.Atoi(x[secondStart:step])
			if err != nil {
				fmt.Errorf("error parsing '%s' as integer: %v", x[secondStart:step], err)
            }
			mulResults = append(mulResults, []int{int(firstNumber),int(secondNumber)})
		}
	}

	result := 0
	for i := 0; i < len(mulResults); i++ {
		result = result + (mulResults[i][0] * mulResults[i][1]) 
	}
	return result
}
