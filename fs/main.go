package main

import (
	"fmt"
	"os"
)

const F_NAME = "hello.txt"

func main() {

	f, err := os.Create(F_NAME)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// data
	data := []byte("I love you so so much!\r\n\t")

	n, _ := f.Write(data)
	fmt.Println("Written", n, "bytes!")

	f, err = os.OpenFile(F_NAME, os.O_APPEND, 0666)
	f.WriteString(string(data))

}
