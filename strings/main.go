package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0.33

	x := fmt.Sprintf("%.1f", n)

	fmt.Println(x)

	fox := "The quick brown fox jumps!"

	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, fox))

	fmt.Println(fox < "Time") // true

	fmt.Println(strings.Compare(fox, "The")) // fox > The
}
