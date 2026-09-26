package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a string: ")

	// Use bufio.NewReader so we can read the entire line including spaces
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		fmt.Println("Not Found!")
		return
	}

	// Remove newline and leading/trailing whitespace
	cleaned := strings.TrimSpace(input)

	// Convert to lowercase to make it case-insensitive
	lower := strings.ToLower(cleaned)

	// Check if:
	// 1. Starts with 'i'
	// 2. Contains 'a'
	// 3. Ends with 'n'
	if strings.HasPrefix(lower, "i") && strings.Contains(lower, "a") && strings.HasSuffix(lower, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
