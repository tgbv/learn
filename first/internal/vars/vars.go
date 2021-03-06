package vars

import "fmt"

/*
SomeFunc function
*/
func SomeFunc() {

	// declare differnet var types
	var (
		apples     int
		pineapples int
	)

	// set different var values
	apples = 1
	pineapples = 2

	// declare && initialize
	size, width, height := 0, 100, 100

	// update values
	size, width, height = width*height, 0, 0

	// print them
	fmt.Println(apples, pineapples)
	fmt.Println(size, width, height)

}
