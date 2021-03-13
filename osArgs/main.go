package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please input name!")
	}

	name := os.Args[1]

	fmt.Println("Hello", name, "!")

	name = "gandalf"

	fmt.Println("I am", name, "!")
	fmt.Printf("Hello %s!\n", name)
}
