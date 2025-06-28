package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/waarispierre/adventOfCode/loadData"
	"github.com/waarispierre/adventOfCode/dayOne"
	"github.com/waarispierre/adventOfCode/dayTwo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'deploy' subcommand")
		os.Exit(1)
	}

	loadData.readData();
}
