package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument. Options are 1 or ")
		return
	}

	switch os.Args[1] {
	case "1":
		challengeOne()
	case "2":
		challengeTwo()
	}

	fmt.Println("Duration:", time.Since(start))
}

type FreeSpace struct {
	startPos int
	endPos   int
}

func (fs FreeSpace) NumberOfFreeSpaces() int {
	return fs.endPos - fs.startPos + 1
}

func challengeTwo() {
	data, err := getData("day09/daynine.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	individualBlocks, freeSapces := rearrangeFiles(data)
	individualBlocks = moveFilesToBlocks(individualBlocks, freeSapces)

	result := checkSumTwo(individualBlocks)
	fmt.Println("Result:", result)
}

func checkSumTwo(individualBlocks []string) int {
	result := 0
	for pos, str := range individualBlocks {
		value, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		result += value * pos
	}

	return result
}

func moveFilesToBlocks(individualBlocks []string, freeSapces []FreeSpace) []string {
	id := 0
	var err error
	done := make(map[int]bool)
	
	//start is the starting position of the last character that is used to determine the file sizes
	for start := len(individualBlocks) - 1; start >= 0; start-- {
		id, err = strconv.Atoi(individualBlocks[start])
		exist := done[id]
		if err != nil || id == 0 || exist {
			continue
		}

		idCount := 0
		for i := start; i >= 0; i-- {
			currId, _ := strconv.Atoi(individualBlocks[i])
			if id == currId {
				idCount++
			} else {
				start = start - (start - i) + 1
				break
			}
		}

		done[id] = true

		for ind, fs := range freeSapces {
			if fs.startPos >= start {
				break
			}
			if idCount <= fs.NumberOfFreeSpaces() {
				for n := range idCount {
					individualBlocks[fs.startPos+n] = fmt.Sprint(id)
				}

				freeSapces[ind] = FreeSpace{startPos: fs.startPos + idCount, endPos: fs.endPos}
				if freeSapces[ind].endPos < freeSapces[ind].startPos {
					freeSapces = append(freeSapces[0:ind], freeSapces[ind+1:]...)
				}
				end := (start + idCount)
				for u := start; u < end && u < len(individualBlocks); u++ {
					individualBlocks[u] = "."
				}
				break
			}
		}
	}
	return individualBlocks
}

func challengeOne() {
	data, err := getData("day09/daynine.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	individualBlocks, _ := rearrangeFiles(data)
	individualBlocks = moveFiles(individualBlocks)
	result := checkSum(individualBlocks)

	fmt.Println("Result:", result)
}

func checkSum(individualBlocks []string) any {
	checksum := 0
	val := 0
	res := 0
	for i, j := range individualBlocks {
		val, _ = strconv.Atoi(j)
		res = i * val
		checksum += res
	}
	return checksum
}

func moveFiles(individualBlocks []string) []string {
	lastVal := len(individualBlocks) - 1
	for pos, r := range individualBlocks {
		if pos >= len(individualBlocks) {
			break
		}
		if r == "." {
			individualBlocks[pos] = individualBlocks[lastVal]
			individualBlocks = individualBlocks[:lastVal]
			lastVal--
			for individualBlocks[lastVal] == "." {
				individualBlocks = individualBlocks[:lastVal]
				lastVal--
			}
		}
	}
	return individualBlocks
}

func rearrangeFiles(data []int) (individualBlocks []string, fs []FreeSpace) {
	id := 0
	numOfBlocks := 0
	freeSpace := 0
	fsp := FreeSpace{}

	for i := 0; i < len(data); i += 2 {
		numOfBlocks = data[i]
		for range numOfBlocks {
			individualBlocks = append(individualBlocks, fmt.Sprint(id))
		}

		if i+1 >= len(data) {
			continue
		}
		freeSpace = data[i+1]
		if freeSpace != 0 {
			fsp.startPos = len(individualBlocks)
			fsp.endPos = fsp.startPos + freeSpace - 1
			fs = append(fs, fsp)
		}
		for range freeSpace {
			individualBlocks = append(individualBlocks, ".")
		}
		id++
	}
	return
}

func getData(path string) (data []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	currentLine := ""
	for scanner.Scan() {
		currentLine = scanner.Text()
		if currentLine == "" {
			continue
		}

		num := 0
		for _, n := range currentLine {
			num, _ = strconv.Atoi(
				string(n),
			)
			data = append(
				data,
				num,
			)
		}
		continue
	}

	return
}
