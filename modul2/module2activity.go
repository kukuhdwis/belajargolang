package main

import (
	"fmt"
)

// GenDisplaceFn takes acceleration (a), initial velocity (vo), and initial displacement (so).
// It returns a function that computes displacement s as a function of time (t)
// using the formula: s = ½ a t^2 + vo*t + so.
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fmt.Println("Displacement:", fn(t))
}
