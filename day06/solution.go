package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var mut sync.RWMutex

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		ChallengeOne()
	case "2":
		ChallengeTwo()
	default:
		fmt.Println("Options 1 or 2")
	}

	fmt.Println("Duration:", time.Since(start))
}

func ChallengeTwo() {
	data, pos, dir, err := getData()
	if err != nil {
		fmt.Println(err)
		return
	}
	var originalPos []int = pos
	var originalDir rune = dir
	maxY := len(data)
	maxX := len(data[0])
	counter := 0

	mapCh := make(chan [][]rune, 30)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
	 	go findLoop(mapCh, &counter, originalDir, originalPos, maxX, maxY, &wg)	
	}

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		defer close(mapCh)
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				if data[y][x] == '#' || (originalPos[0] == x && originalPos[1] == y) {
					continue
				}
				pos = make([]int, len(originalPos))
				dir = originalDir
				copy(pos, originalPos)
				updateData := make([][]rune, len(data))
				for i, row := range data {
					updateData[i] = make([]rune, len(data[0]))
					copy(updateData[i], row)
				}
				addObstacle(&updateData, x, y)
				mapCh <- updateData
			}
		}
	}()

	wg.Wait()
	
	fmt.Println("Number of options:", counter)
}

func findLoop(mapCh chan [][]rune, counter *int, origDir rune, origPos []int, maxX, maxY int, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := range mapCh {
		dir := origDir
		pos := make([]int, len(origPos))
		copy(pos, origPos)
		positions := make(map[string]int)
		addPostionToLoopMap(&positions, pos, dir)	
		var onMappedArea bool = true
		var foundLoop bool = false
		for onMappedArea && !foundLoop {
			moveGuard(&ch, &dir, &pos, &onMappedArea, maxX, maxY)
			addPostionToLoopMap(&positions, pos, dir)
			if !onMappedArea {
				continue
			}
			if positions[fmt.Sprintf("%c:%v:%v", dir, pos[0], pos[1])] > 1 {
				incrementCounter(counter)
				foundLoop = true
				continue
			}
		}
	}
}

func incrementCounter(counter *int) {
	mut.Lock()
	*counter++
	mut.Unlock()
}


func addPostionToLoopMap(positions *map[string]int, pos []int, dir rune) {
	(*positions)[fmt.Sprintf("%c:%v:%v", dir, pos[0], pos[1])] += 1 
}

func addObstacle(data *[][]rune, x, y int) {
	(*data)[y][x] = 'O'	
}

func ChallengeOne() {
	data, pos, dir, err := getData()
	if err != nil {
		fmt.Println(err)
		return
	}
	positions := make(map[string]int)
	addPostionToMap(&positions, pos)
	var onMappedArea bool = true
	maxY := len(data)
	maxX := len(data[0])
	steps := 1
	for onMappedArea {
		moveGuard(&data, &dir, &pos, &onMappedArea, maxX, maxY)
		if !onMappedArea {
			continue
		}
		addPostionToMap(&positions, pos)
		steps += 1
	}
	fmt.Println("Distinct Positions", len(positions))
}

func moveGuard(data *[][]rune, dir *rune, pos *[]int, onMappedArea *bool, maxX, maxY int) {
	curX := (*pos)[0]
	curY := (*pos)[1]
	incrX := 0
	incrY := 0

	switch *dir {
	case 'v':
		incrY = 1
	case '<':
		incrX = -1
	case '^':
		incrY = -1
	case '>':
		incrX = 1
	}

	newY := curY + incrY
	newX := curX + incrX
	if !inbounds(newX, newY, maxX, maxY) {
		*onMappedArea = false
		return
	}
	if (*data)[newY][newX] == '.' {
		(*data)[newY][newX] = *dir
		(*data)[curY][curX] = '.'
	} else {
		switch *dir {
		case 'v':
			incrY = 0
			incrX = -1
			*dir = '<'
			if (*data)[curY][curX-1] != '.' {
				incrY = -1
				incrX = 0	
				*dir = '^'
			}
		case '<':
			incrY = -1
			incrX = 0
			*dir = '^'
			if (*data)[curY-1][curX] != '.' {
				incrY = 0
				incrX = 1	
				*dir = '>'
			}
		case '^':
			incrY = 0
			incrX = 1
			*dir = '>'
			if (*data)[curY][curX+1] != '.' {
				incrY = 1
				incrX = 0	
				*dir = 'v'
			}
		case '>':
			incrY = 1
			incrX = 0
			*dir = 'v'
			if (*data)[curY+1][curX] != '.' {
				incrY = 0
				incrX = -1	
				*dir = '<'
			}
		}
		newX = curX + incrX
		newY = curY + incrY
		(*data)[curY][curX] = '.'
		(*data)[newY][newX] = *dir
	}
	*pos = updatePosition(newX, newY)
}

func addPostionToMap(positions *map[string]int, pos []int) {
	(*positions)[fmt.Sprintf("%v:%v", pos[0], pos[1])] += 1 
}

func printMap(data [][]rune) {
	for _, row := range data {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	} 
}

func inbounds(x, y, maxX, maxY int) bool {
	if x < 0 {
		return false
	}
	if x >= maxX {
		return false
	}
	if y < 0 {
		return false
	}
	if y >= maxY {
		return false
	}
	return true
}

func updatePosition(x , y int) (pos []int) {
	pos = append(pos, x)
	pos = append(pos, y)
	return
}

func getData() (data [][]rune, pos []int, dir rune, err error) {
	file, err := os.Open("day06/daysix.txt")
	if err != nil {
		return 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentLine string
	var x int
	var y int = 0
	for scanner.Scan() {
		x = 0
		currentLine = scanner.Text()
		if currentLine == "" {
			continue
		}
		var row []rune	
		for _, r := range currentLine {
			row = append(row, r)
			if r == '>' || r == '<' || r == '^' || r == 'v' {
				dir = r
				pos = append(pos, x)
				pos = append(pos, y)
			} 
			
			x += 1
		}
		data = append(data, row)
		y += 1
	}
	
	return
}