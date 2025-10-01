package loaddata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData(filename string) ([][]int, error) {
	path := fmt.Sprintf("inputdata/%v", filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		parts := strings.Fields(line) // Split by any whitespace
		row := make([]int, len(parts))

		for i, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing '%s' as integer: %v", part, err)
			}
			row[i] = val
		}

		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

func ReadDataInString(filename string) (string, error) {
	path := fmt.Sprintf("inputdata/%v", filename)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}
		builder.WriteString(line)
	}

	return builder.String(), scanner.Err()
}

func GetColumn(matrix [][]int, colIndex int) []int {
	if len(matrix) == 0 {
		return nil
	}

	column := make([]int, len(matrix))
	for i, row := range matrix {
		if colIndex >= len(row) {
			return nil // Column index out of bounds
		}
		column[i] = row[colIndex]
	}
	return column
}

// GetRow extracts a specific row from the matrix
func GetRow(matrix [][]int, rowIndex int) []int {
	if rowIndex >= len(matrix) || rowIndex < 0 {
		return nil
	}

	// Return a copy to avoid modifying original
	row := make([]int, len(matrix[rowIndex]))
	copy(row, matrix[rowIndex])
	return row
}
