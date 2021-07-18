package main

import (
	"fmt"
	"os"
)

// set at compile time
var version = "unknown"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Print(version)
		os.Exit(0)
	}
	fmt.Printf("Hello! This is version '%s'\n", version)

	fmt.Scanln()
}
