package main

import "fmt"

func shout(userName string) {
	fmt.Printf("%s is learning how to call functions!\n", userName)
}

func function() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)

	// call the function directly, or within a fmt.Println statement here.
	shout(userName)
}
