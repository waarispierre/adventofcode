package daythree

import (
	"adventofcode/shared/loaddata"
	"fmt"
	"regexp"
	"strconv"
)

func ChallengeOne() {
	fmt.Println("Day three challenge one started")
	data, err := loaddata.ReadDataInString("day03/daythree.txt")	
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	calculateTotal(&total, data)	
	fmt.Println(total)
}

func ChallengeTwo() {
	fmt.Println("Day three challenge two started")
	data, err := loaddata.ReadDataInString("day03/daythree.txt")	
	if err != nil {
		fmt.Println(err)
		return
	}

	patternDo := `do\(\)`
	regexDo := regexp.MustCompile(patternDo)
	splitsDo := regexDo.FindAllStringIndex(data, -1)

	patternDont := `don't\(\)`
	regexDont := regexp.MustCompile(patternDont)
	splitsDont := regexDont.FindAllStringIndex(data, -1)

	type instruction struct {
		pos int
		enabled bool
	}

	var instructions []instruction
	for _, split := range splitsDo {
		instructions = append(instructions, instruction{pos: split[0], enabled: true})
	}
	for _, split := range splitsDont {
		instructions = append(instructions, instruction{pos: split[0], enabled: false})
	}

	for i := 0; i < len(instructions)-1; i++ {
		for j := i + 1; j < len(instructions); j++ {
			if instructions[i].pos > instructions[j].pos {
				instructions[i], instructions[j] = instructions[j], instructions[i]
			}
		}
	}
	//Same as
	// sort.Slice(instructions, func(i, j int) bool {
	// 	return instructions[i].pos < instructions[j].pos
	// })

	var enabledSections []string
	enabled := true
	position := 0
	var instr instruction
	var inBounds bool

	for i := 0; i <= len(instructions); i++ {
		inBounds = i < len(instructions)	
		if inBounds {
			instr = instructions[i]
		}
		if enabled && inBounds {
			enabledSections = append(enabledSections, data[position:instructions[i].pos])
		} else if enabled {
			enabledSections = append(enabledSections, data[position:])
		}

		enabled = instr.enabled
		position = instr.pos
	}

	total := 0
	for _, dataToUse := range enabledSections {
		calculateTotal(&total, dataToUse)
	}
	
	fmt.Println(total)
}

func calculateTotal(total *int, dataToUse string) {
	pattern := `mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`
	regex := regexp.MustCompile(pattern)
	
	matches := regex.FindAllStringSubmatch(dataToUse, -1)
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
		
		*total += first * second
	}
}
