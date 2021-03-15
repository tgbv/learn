package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// check num of args
	if len(os.Args) < 2 {
		panic(`Please input a value`)
	}

	val := os.Args[1]
	len := utf8.RuneCountInString(val)
	ems := strings.Repeat("!", len)

	fmt.Println(ems + val + ems)
}
