package main

import "fmt"

func movies() {
	var age int
	fmt.Scanf("%d", &age)
	// Code your switch or if...else-if statement here.
	switch {
	case age <= 14:
		fmt.Println("Toy Story 4")
	case 15 <= age && age <= 18:
		fmt.Println("The Matrix")
	case 19 <= age && age <= 25:
		fmt.Println("John Wick")
	case 26 <= age && age <= 35:
		fmt.Println("Constantine")
	case 36 <= age:
		fmt.Println("Speed")
	default:
		fmt.Println("Invalid")
	}
}
