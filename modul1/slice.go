package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create an empty integer slice with initial capacity of 3
	sli := make([]int, 0, 3)

	for {
		var input string
		fmt.Print("Enter an integer (or 'X' to exit): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		// Exit condition
		if strings.EqualFold(input, "x") {
			fmt.Println("Exiting program.")
			break
		}

		// Convert string input to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X' to exit.")
			continue
		}

		// Append the new integer to the slice (grows dynamically)
		sli = append(sli, num)

		// Sort the slice in ascending order
		sort.Ints(sli)

		// Print the sorted slice
		fmt.Println(sli)
	}
}
