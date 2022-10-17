package main

import "fmt"
import "os"
import "log"
import "strings"

func forbidden() {
	name, word := "", ""
	fmt.Scan(&name)
	file, err := os.ReadFile(name) 
    if err != nil {
        log.Fatal(err)
    }
    file_content := strings.ToLower(string(file))

    for true {

        fmt.Scan(&word)
        if word == "exit" {
            fmt.Println("Bye!")
            return
        } else if strings.Contains(file_content, strings.ToLower(word)) {
            for i:=0; i<len(word); i++ {
                fmt.Print("*")
            }
            fmt.Println()
        } else {
            fmt.Println(word)
        }
    }
}
