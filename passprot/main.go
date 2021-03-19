package main

import (
	"fmt"
)

func checkCredentials(usr string, pss string) {
	if usr == "john" && pss == "doe" {
		fmt.Println("Congrats!")
	} else {
		fmt.Println("Username and password are incorrect! Please retry")
		scanInput()
	}
}

func scanInput() (string, string) {
	var usr, pss string

	fmt.Println("Please input a username and a password.")

	i, err := fmt.Scanf("%s %s\n", &usr, &pss)
	if err != nil || i < 2 {
		fmt.Println("Wrong data.")
		return scanInput()
	}

	return usr, pss
}

func main() {
	// scan input

	// check usr && pss
	checkCredentials(scanInput())
}
