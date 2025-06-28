package dayOne

import (
	"fmt"
	"github.com/waarispierre/adventOfCode/loadData"
	"math"
	"sort"
)

func Challenge() {
	matrix, err := loadData.ReadData("inputOne.txt")

	if err != nil {
		fmt.Sprintf("An error occured loading the data %v\n", err)
	}

	fmt.Sprintf("loook: %v\n", matrix)

	col1 := loadData.GetColumn(matrix, 0)
	col2 := loadData.GetColumn(matrix, 1)

	sort.Ints(col1)
	sort.Ints(col2)

	var totalDistance int
	for i := 0; i < len(col1); i++ {
		totalDistance = totalDistance + int(math.Abs(float64(col1[i]-col2[i])))
	}

	col2Counts := make(map[int]int)
	for _, value := range col2 {
		col2Counts[value]++
	}

	var total int
	for _, value := range col1 {
		total = total + (col2Counts[value] * value)
	}

	fmt.Println(totalDistance)
	fmt.Println(total)
}
