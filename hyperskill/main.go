package main

import (
	"fmt"
)

func main() {
	fmt.Print("Choose a program: ")
	program := "coffee"
	fmt.Scan(&program)
	switch program {
	case "reverse":
		reverse()
	case "sum":
		sum()
	case "enum":
		enum()
	case "coffee":
		coffee()
	case "movies":
		movies()
	case "positive":
		positive()
	case "function":
		function()
	case "coffee2":
		coffee2()
	case "odds":
		odds()
	default:
		fmt.Println("Unknown.")
	}
}
