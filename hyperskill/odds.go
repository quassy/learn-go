package main

import "fmt"

func odds() {
	for i := 2; i < 1024; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

	f := 0
	p := 1
	fmt.Scanf("%d", &f)

	for i := 1; i <= f; i++ {
		p *= i
	}
	fmt.Println(p)
}
