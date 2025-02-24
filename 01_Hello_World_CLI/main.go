package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Define command-line flags
	uppercase := flag.Bool("uppercase", false, "Print output in uppercase")
	lowercase := flag.Bool("lowercase", false, "Print output in lowercase")
	colored := flag.Bool("color", false, "Print colored output")

	// Parse the command-line flags
	flag.Parse()

	// Default message
	normal := "Hello, World!"

	// Check if user provided a custom message
	if flag.NArg() > 0 {
		normal = strings.Join(flag.Args(), " ")
	}

	// Handle different formats
	if *uppercase {
		fmt.Println(strings.ToUpper(normal)) // Uppercase
	} else if *lowercase {
		fmt.Println(strings.ToLower(normal)) // Lowercase
	} else if *colored {
		color.Red("%s | Red Color", normal) // Red text
		color.Green("%s | Green Color", normal) // Green text
		color.Blue("%s | Blue Color", normal) // Blue text
	} else {
		fmt.Println(normal)
	}
}