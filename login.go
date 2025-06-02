package main

import (
	"errors"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Users []User

func Login(users *Users) (string, error) {

	println("Welcome to the Todo CLI Application!")
	println("Please enter your username:")

	var username string
	_, err := fmt.Scanf("%s", &username)
	if err != nil {

		return "", err
	}

	println("Please enter your password:")
	var password string
	_, err = fmt.Scanf("%s", &password)
	if err != nil {

		return "", err
	}
	for _, user := range *users {
		if user.Username == username && user.Password == password {
			println("Login successful!")
			return username, nil
		}
	}

	return "", errors.New("invalid username or password, please try again or create a new user")

}

func CreateUser(users *Users) (string, error) {
	println("Please enter a username:")
	var username string
	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		return "", err
	}
	for _, user := range *users {
		if user.Username == username {
			return "", errors.New("username already exists")
		}
	}
	println("Please enter a password:")
	var password string
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		return "", err
	}

	*users = append(*users, User{Username: username, Password: password})
	println("User created successfully!")
	return username, nil
}
