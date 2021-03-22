package main

import "fmt"

var users []map[string]string

func addUser(usr string, pss string) {
	users = append(users, map[string]string{
		"username": usr,
		"password": pss,
	})
}

func delUser(usr string) {
	for i, v := range users {
		if v["username"] == usr {
			firstCut := users[:i]
			secondCut := users[i+1:]
			users = append(firstCut, secondCut...)
			break
		}
	}
}

func main() {
	addUser("kane", "some password")
	addUser("johnson", "some password")
	addUser("david", "some password")
	delUser("johnson")
	fmt.Println(users)
}
