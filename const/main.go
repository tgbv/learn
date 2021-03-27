package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	magic  = iota
	M_DIFF = 0.3048
	Y_DIFF = 0.3333
)

func main() {

	if len(os.Args) < 2 {
		os.Args = append(os.Args, "0")
	}

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * M_DIFF
	yards := feet * Y_DIFF

	fmt.Printf("%gfeet is %g meters\n", feet, meters)
	fmt.Printf("%gfeet is %g yards\n", feet, yards)
}
