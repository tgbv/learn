package main

import "fmt"

func main() {

	array := [5]int{34, 123, 45, 765, 87}
	fmt.Println(array)

	slice := array[0:2]
	slice[1] = 22
	fmt.Println(slice)

	//	make()

	newSlice := make([]uint, 1)
	newSlice = append(newSlice, 222, 33, 444, 22)

	strings := make([]string, 10)
	strings = append(strings, "hello!!!")
	for _, v := range strings {
		fmt.Printf("%T ", v)
	}
	fmt.Println()

	twoDStrings := make([][20]int, 10)

	fmt.Println(len(twoDStrings[0]))

	//fmt.Println(cap(strings))

}
