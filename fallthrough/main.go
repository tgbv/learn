package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Input integer!")
	}

	// fallthrough :()
	switch num = 212; {

	case num > 100:
		fmt.Print("big ")
		fallthrough

	case num < 0:
		fmt.Print("positive ")

	default:
		fmt.Print("number.")

	}
}
