package main

import (
	"fmt"
)

func main() {
	// Input: Weight (kg) and Height (m)
	var weight, height float64

	fmt.Print("Enter your weight (kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Enter your height (m): ")
	fmt.Scanln(&height)

	// Calculate BMI
	bmi := weight / (height * height)

	// Categorize BMI
	var category string
	switch {
	case bmi < 18.5:
		category = "Underweight"
	case bmi >= 18.5 && bmi < 24.9:
		category = "Normal weight"
	case bmi >= 25 && bmi < 29.9:
		category = "Overweight"
	default:
		category = "Obese"
	}

	// Output the result
	fmt.Printf("\nYour BMI: %.2f\n", bmi)
	fmt.Printf("Category: %s\n", category)
}
