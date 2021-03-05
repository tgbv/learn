package main

// why Go, why won't you support relative paths ? ;-;
import (
	"fmt"
)

var short = 222
var hello = "congrats broo"

func main() {

	short = 33

	fmt.Println(short, hello)

}
