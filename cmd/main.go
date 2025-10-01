package main

import (
	"adventofcode/dayone"
	"adventofcode/daythree"
	"adventofcode/daytwo"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dayOne":
		dayone.Challenge()
	case "dayOneSimple":
		dayone.ChallengeSimple()
	case "dayOne-2":
		dayone.ChallengeTwo()
	case "dayTwo":
		daytwo.ChallengeOne()
	case "dayTwo-2":
		daytwo.ChallengeTwo()
	case "dayThree":
		daythree.ChallengeOne()
	case "dayThree-2":
		daythree.ChallengeTwo()


	default:
		fmt.Println("Please provide a parameter to say which challenge to run.\nOptions examples: dayOne, dayOneSimple, dayTwo etc.")
	}

	fmt.Printf("Duration of run: %v", time.Since(start))
}
