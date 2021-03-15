package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "alex șmecher"

	fmt.Println(len(name))

	fmt.Println(utf8.RuneCountInString(name))

}
