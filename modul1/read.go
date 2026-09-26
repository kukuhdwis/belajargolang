package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define struct named name with fields fname and lname
type name struct {
	fname string
	lname string
}

func main() {
	// Prompt user for file name
	var fileName string
	fmt.Print("Enter the name of the text file: ")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Slice of name structs
	names := make([]name, 0)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split line by whitespace into first name and last name
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			n := name{
				fname: parts[0],
				lname: parts[1],
			}
			// Add struct to slice
			names = append(names, n)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	// Iterate through the slice and print each first and last name pair
	fmt.Println("\nNames found in file:")
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}
