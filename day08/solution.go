package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Please parse an argument. Options are 1 or 2")
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

func challengeOne() {
	fmt.Println("Starting challenge 1")
	data, maxX, maxY, err := getData("day08/dayeight.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	posMap := make(map[coordinates]int)
	for _, i := range data {
		for z, j := range i {
			for _, pos := range i[z+1:] {
				vertDiff := j.y - pos.y
				horizDiff := j.x - pos.x
				first := coordinates{
					x: j.x + horizDiff,
					y: j.y + (1 * vertDiff),
				}
				second := coordinates{
					x: pos.x + (-1 * horizDiff),
					y: pos.y - vertDiff,
				}
				if first.x >= 0 && first.x <= maxX && first.y >= 0 && first.y <= maxY {
					posMap[first] += 1
				}
				if second.x >= 0 && second.x <= maxX && second.y >= 0 && second.y <= maxY {
					posMap[second] += 1
				}
			}
		}
	}
	fmt.Println("Result:", len(posMap))
}

func challengeTwo() {
	fmt.Println("Starting challenge 2")
	data, maxX, maxY, err := getData("day08/dayeight.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	posMap := make(map[coordinates]int)
	for _, i := range data {
		for z, j := range i {
			for _, pos := range i[z+1:] {
				VertDiff := j.y - pos.y
				HorizDiff := j.x - pos.x

				firstInbouncs := true
				secondInbouncs := true

				dirOneNode := coordinates{x: j.x, y: j.y}
				posMap[dirOneNode] += 1
				for firstInbouncs {
					dirOneNode = coordinates{
						x: dirOneNode.x + HorizDiff,
						y: dirOneNode.y + (1 * VertDiff),
					}
					if dirOneNode.x >= 0 && dirOneNode.x <= maxX && dirOneNode.y >= 0 && dirOneNode.y <= maxY {
						posMap[dirOneNode] += 1
					} else {
						firstInbouncs = false
						continue
					}
				}

				dirTwoNode := coordinates{x: j.x, y: j.y}
				posMap[dirTwoNode] += 1
				for secondInbouncs {
					dirTwoNode = coordinates{
						x: dirTwoNode.x + (-1 * HorizDiff),
						y: dirTwoNode.y - VertDiff,
					}
					if dirTwoNode.x >= 0 && dirTwoNode.x <= maxX && dirTwoNode.y >= 0 && dirTwoNode.y <= maxY {
						posMap[dirTwoNode] += 1
					} else {
						secondInbouncs = false
						continue
					}
				}
			}
		}
	}
	fmt.Println("Result:", len(posMap))
}

type coordinates struct {
	x int
	y int
}

func getData(path string) (map[rune][]coordinates, int, int, error) {
	data := make(map[rune][]coordinates)
	file, err := os.Open(path)
	if err != nil {
		return data, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	maxX := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			continue
		}
		maxX = len(currentLine)
		for x, r := range currentLine {
			if r != '.' && r != '#' {
				data[r] = append(
					data[r],
					coordinates{
						x: x,
						y: y,
					},
				)
			}
		}
		y += 1
	}
	return data, maxX - 1, y - 1, nil
}
