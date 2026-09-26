package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct contains the traits of an animal.
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints the food that the animal eats.
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints the locomotion method of the animal.
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints the spoken sound of the animal.
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Predefined animals with hard-coded data
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

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
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter 2 words: <cow|bird|snake> <eat|move|speak>")
			continue
		}

		animalName := strings.ToLower(parts[0])
		action := strings.ToLower(parts[1])

		var selectedAnimal Animal
		switch animalName {
		case "cow":
			selectedAnimal = cow
		case "bird":
			selectedAnimal = bird
		case "snake":
			selectedAnimal = snake
		default:
			fmt.Println("Invalid animal name. Must be: cow, bird, or snake")
			continue
		}

		switch action {
		case "eat":
			selectedAnimal.Eat()
		case "move":
			selectedAnimal.Move()
		case "speak":
			selectedAnimal.Speak()
		default:
			fmt.Println("Invalid action. Must be: eat, move, or speak")
		}
	}
}
