package main

import "fmt"

func main() {
	numbers := make([]int, 0, 10)
	fmt.Println("Enter up to 10 integers:")
	for len(numbers) < 10 {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		numbers = append(numbers, n)
	}

	BubbleSort(numbers)

	for i, n := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(n)
	}
	fmt.Println()
}

func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func Swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}
