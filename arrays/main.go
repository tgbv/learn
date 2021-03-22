package main

import "fmt"

type meme struct {
	length uint
	pepe   func()
}

func (m *meme) Throw() uint {
	return m.length*10000 + 1
}

func main() {

	var fruits [3]string

	fruits[0] = "watermellon"
	fruits[1] = "apples"
	fruits[2] = "bananas"

	fmt.Println(fruits)

	numbers := [3]uint32{545, 23, 66}
	fmt.Println(numbers)

	multiDimensional := [3][2]uint32{}
	fmt.Println(multiDimensional)

	someMemeStruct := [2]meme{
		{length: 0, pepe: func() {}},
		{length: 111},
	}
	someMemeStruct[0].pepe()
	fmt.Println()

	for i, v := range multiDimensional {
		for j, b := range v {
			multiDimensional[i][j] = 22
			fmt.Println(b)
			fmt.Println(multiDimensional[i][j])
		}
	}
}
