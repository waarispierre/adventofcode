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
	partTwo(matrix)
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

func partTwo(matrix [][]int) {
	safeReports := 0

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		if !IsSafe(row) {
			for x := 0; x < len(row); x++ {
				newRow := removeIndex(row, x)
				if IsSafe(newRow) {
					safeReports++
					break
				}
			}
		} else {
			safeReports++	
		}
	}
	fmt.Printf("Safe reports: %v\n", safeReports)
}

func removeIndex(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice // Invalid index
    }
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
    return result
}

func IsSafe(row []int) bool {
	ruleOne := IsIncreasingWithinBounds(row)
	ruleTwo := IsDecreasingWithinBounds(row)

	return ruleOne != ruleTwo
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
