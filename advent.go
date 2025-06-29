package main

import (
	"fmt"
	"github.com/waarispierre/adventOfCode/dayOne"
	"github.com/waarispierre/adventOfCode/dayTwo"
	"github.com/waarispierre/adventOfCode/dayThree"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Expected 'challenge' subcommand")
		os.Exit(1)
	}

	switch os.Args[2] {
	case "dayOne":
		dayOne.Challenge()
	case "dayTwo":
		dayTwo.Challenge()
	case "dayThree":
		dayThree.Challenge()
	default:
		fmt.Println("Please provide a parameter to say which challenge to run.\nOptions examples: dayOne, dayTwo etc.")
	}

	fmt.Printf("Duration of run: %v", time.Since(start))
}
