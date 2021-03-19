package main

import (
	"fmt"
	"time"
)

func main() {
	h := time.Now().Hour()

	switch {

	case h >= 6 && h < 14:
		fmt.Println("Good morning!")
	case h >= 14 && h < 22:
		fmt.Println("Good afternoon!")
	case h >= 22,
		h > 0 && h < 6:
		fmt.Println("Good evening!")

	}
}
