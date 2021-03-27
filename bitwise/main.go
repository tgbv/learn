package main

import (
	"fmt"
	"strconv"
)

const (
	FIRST = 1 << iota
	SECOND
	THIRD
	FOURTH
	MORTII_TAI
)

func b(n uint8) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {

	//fmt.Println(b(FIRST), b(FOURTH))
	fmt.Println(MORTII_TAI)
}
