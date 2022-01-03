package main

import "fmt"

func reverse() int {
	input := 0
	fmt.Scan(&input)
	output := 0

	for i := 0; i < 100; i++ {
		output = output*10 + input%10
		input = input / 10
		if input == 0 {
			break
		}
	}

	fmt.Println(output)
	return output
}
