package main

import (
	"fmt"
	"math"
)

func ImproveGuess(argument float64, bestGuess float64) (bool, float64) {
	nextGuess := (bestGuess + (argument / bestGuess)) / float64(2.0)

	fmt.Printf("Improving guess from %T %v to %T %v", bestGuess, bestGuess, nextGuess, nextGuess)

	delta := math.Abs(nextGuess - bestGuess)
	return (delta < float64(math.SmallestNonzeroFloat32)), nextGuess
}

func Sqrt(argument float64) {
	isDone := false
	guess := argument * 0.5

	fmt.Printf("Calculating square root of %v.", argument)

	for !isDone {
		isDone, guess = ImproveGuess(argument, guess)
		fmt.Println("")
	}

	fmt.Printf("The square root of %v ≈ %v (%v)", argument, guess, math.Sqrt(argument))
}

func main() {

	Sqrt(10)
}
