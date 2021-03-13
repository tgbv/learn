package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	fmt.Println("Hello", name, "!")

	name = "gandalf"

	fmt.Println("I am", name, "!")
}
