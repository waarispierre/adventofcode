package daythree

import (
	"adventofcode/loaddata"
	"fmt"
	"regexp"
	"strconv"
)

func ChallengeOne() {
	fmt.Println("Day three challenge one started")
	data, err := loaddata.ReadDataInString("daythree.txt")	
	if err != nil {
		fmt.Println(err)
		return
	}

	pattern := `mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`
	regex := regexp.MustCompile(pattern)
	
	matches := regex.FindAllStringSubmatch(data, -1)
	total := 0
	for _, match := range matches {
		first, err1 := strconv.Atoi(match[1])
		second, err2 := strconv.Atoi(match[2])
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		
		total += first * second
	}

	fmt.Println(total)
}
