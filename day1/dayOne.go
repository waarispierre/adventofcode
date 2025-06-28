package dayOne

import (
	"bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
	"sort"
	"math"
)

func main() {
	col1, col2 := loadData()

	sort.Ints(col1)
	sort.Ints(col2)
	
	var totalDistance int
	for i := 0; i < len(col1); i++ {
		totalDistance = totalDistance + int(math.Abs(float64(col1[i] - col2[i])))
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

func loadData() ([]int, []int) {
	file, err := os.Open("inputOne.txt")

	if err != nil {
		fmt.Printf("Error reading the input data: %v\n", err)
	}
	
	defer file.Close()

	var column1, column2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Invalid line format: %v\n", line)
			continue
		}

		valueOne, error := strconv.Atoi(parts[0])
		if error != nil {
			fmt.Printf("Error parsing first value: %v\n", parts[0])
			continue
		}

		valueTwo, error := strconv.Atoi(parts[1])
		if error != nil {
			fmt.Printf("Error parsing the second value: %v\n", parts[1])
			continue
		}

		column1 = append(column1, valueOne)
		column2 = append(column2, valueTwo)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file: ", err)
	}

	return column1, column2
}
