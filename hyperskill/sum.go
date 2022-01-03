package main

import "fmt"

func sum() {
	a, b, s := 0, 0, 0
	fmt.Scan(&a, &b)
	s = a + b
	fmt.Println(s)
}
