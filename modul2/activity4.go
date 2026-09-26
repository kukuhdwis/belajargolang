package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface describes the actions an animal can perform.
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow represents the cow animal type.
type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird represents the bird animal type.
type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake represents the snake animal type.
type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Println("Invalid command. Format: 'newanimal <name> <cow|bird|snake>' or 'query <name> <eat|move|speak>'")
			continue
		}

		command := strings.ToLower(parts[0])
		name := parts[1]
		param := strings.ToLower(parts[2])

		switch command {
		case "newanimal":
			switch param {
			case "cow":
				animals[name] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[name] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[name] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal type. Allowed types: cow, bird, snake")
			}

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Printf("Animal '%s' not found.\n", name)
				continue
			}

			switch param {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query type. Allowed queries: eat, move, speak")
			}

		default:
			fmt.Println("Unknown command. Must be 'newanimal' or 'query'")
		}
	}
}
