package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap swaps the contents of the slice at position i with position i+1.
func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

// BubbleSort sorts a slice of integers in ascending order (least to greatest)
// by modifying the slice in-place.
func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				swapped = true
			}
		}
		// If no elements were swapped in this pass, the slice is already sorted
		if !swapped {
			break
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter up to 10 integers (separated by spaces or one per line, press Enter when done):")

	var numbers []int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(numbers) > 0 {
				break
			}
			continue
		}

		fields := strings.Fields(line)
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Warning: '%s' is not a valid integer and will be skipped.\n", field)
				continue
			}
			numbers = append(numbers, num)
			if len(numbers) == 10 {
				break
			}
		}

		// Finish input if 10 numbers have been entered or multiple numbers were entered on a single line
		if len(numbers) >= 10 || len(fields) > 1 {
			break
		}
	}

	if len(numbers) == 0 {
		fmt.Println("No integers were entered.")
		return
	}

	BubbleSort(numbers)

	// Print the sorted integers on one line separated by spaces
	for i, v := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
