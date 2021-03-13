package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// check if arg is provided in command line
	if len(os.Args) < 2 {
		panic("Please input number")
	}

	celsius, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic("Please input a real number !")
	}

	fahrenheit := celsius*1.8 + 32.0

	fmt.Printf("Celsius: %g\nFahrenheit: %g", celsius, fahrenheit)

}
