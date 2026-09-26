package main

import "fmt"

func main() {
	var num float64

	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	truncated := int(num)
	fmt.Println(truncated)
}
