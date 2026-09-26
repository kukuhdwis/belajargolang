package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt and read name
	fmt.Print("Enter name: ")
	nameInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name := strings.TrimSpace(nameInput)

	// Prompt and read address
	fmt.Print("Enter address: ")
	addressInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading address:", err)
		return
	}
	address := strings.TrimSpace(addressInput)

	// Create map and store values
	personMap := map[string]string{
		"name":    name,
		"address": address,
	}

	// Marshal map into JSON
	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print the resulting JSON object as a string
	fmt.Println(string(jsonData))
}
