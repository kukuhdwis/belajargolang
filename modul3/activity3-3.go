package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// sortSubarray is executed by each goroutine to sort its assigned partition
func sortSubarray(id int, arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Print the subarray that this goroutine will sort
	fmt.Printf("Goroutine %d sorting subarray: %v\n", id, arr)

	// Sort the partition in-place
	sort.Ints(arr)
}

// merge merges two sorted integer slices into a single sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int

	// Prompt the user to enter integers
	for {
		fmt.Println("Enter a series of integers separated by spaces (minimum 4 integers):")
		fmt.Print("> ")

		if !scanner.Scan() {
			fmt.Println("No input provided. Exiting.")
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		numbers = make([]int, 0, len(tokens))
		valid := true

		for _, token := range tokens {
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Invalid input '%s'. Please enter only integers.\n\n", token)
				valid = false
				break
			}
			numbers = append(numbers, num)
		}

		if !valid {
			continue
		}

		if len(numbers) < 4 {
			fmt.Println("You must enter at least 4 integers to partition across 4 goroutines.\n")
			continue
		}

		break
	}

	n := len(numbers)
	fmt.Printf("\nTotal elements to sort: %d\n", n)
	fmt.Printf("Initial array: %v\n\n", numbers)

	// Partition the array into 4 parts of approximately equal size
	const numPartitions = 4
	baseSize := n / numPartitions
	remainder := n % numPartitions

	partitions := make([][]int, numPartitions)
	start := 0

	var wg sync.WaitGroup

	// Launch 4 goroutines to sort each partition
	for i := 0; i < numPartitions; i++ {
		end := start + baseSize
		if i < remainder {
			end++ // Distribute remainder elements evenly across first partitions
		}

		// Create a separate slice for each partition
		partitions[i] = make([]int, end-start)
		copy(partitions[i], numbers[start:end])
		start = end

		wg.Add(1)
		go sortSubarray(i+1, partitions[i], &wg)
	}

	// Wait for all 4 goroutines to complete their sorting
	wg.Wait()

	// The main goroutine merges the 4 sorted subarrays into one large sorted array
	fmt.Println("\nAll goroutines finished sorting. Merging subarrays in main goroutine...")
	mergedFirstHalf := merge(partitions[0], partitions[1])
	mergedSecondHalf := merge(partitions[2], partitions[3])
	finalSorted := merge(mergedFirstHalf, mergedSecondHalf)

	// Print the entire sorted list
	fmt.Printf("\nEntire sorted list: %v\n", finalSorted)
}
