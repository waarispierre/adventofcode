package dayTwo

import (
	"fmt"
	"math"
	"github.com/waarispierre/adventOfCode/loadData"
)

func Challenge() {
 	fmt.Println("Day Two")

	matrix, err := loadData.ReadData("inputTwo.txt")

	if err != nil {
		fmt.Sprintf("An error occurred: %v \n", err)
	}

	partOne(matrix)
}

func partOne(matrix [][]int) {
	safeReports := 0
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]

		ruleOne := IsIncreasingWithinBounds(row)
		ruleTwo := IsDecreasingWithinBounds(row)
		
		if ruleOne != ruleTwo {
			safeReports++
		}
	}

	fmt.Printf("Safe reports: %v\n", safeReports)

}

func IsIncreasingWithinBounds(slice []int) bool {

	for i := 0; i < len(slice) - 1; i++ {
	    difference := int(slice[i]) - int(slice[i + 1]) 
		absDifference := int(math.Abs(float64(difference)))
		if difference > 0 || (absDifference < 1 || absDifference > 3) {
			return false
		}
	}
	return true
}

func IsDecreasingWithinBounds(slice []int) bool {

	for i := 0; i < len(slice) - 1; i++ {
	    difference := int(slice[i]) - int(slice[i + 1]) 
		absDifference := int(math.Abs(float64(difference)))
		if difference < 0 || (absDifference < 1 || absDifference > 3) {
			return false
		}
	}
	return true
}
