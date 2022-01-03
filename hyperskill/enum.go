package main

import "fmt"

func enum() {
	const (
		Read = 1 << iota // << A bit operation
		Write
		Execute
	)

	fmt.Println(Read)
	fmt.Println(Write)
	fmt.Println(Execute)
}
