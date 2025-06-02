package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	userChoice := 0
	UserStorage := NewStorage[Users]("users.json")
	users := Users{}
	UserStorage.Load(&users)
	userchoice1 := 0
	passed := false
	loggedinUser := ""

	for userchoice1 != 3 && !passed {
		println("Welcome to the Todo CLI Application!")
		println("1. Login")
		println("2. Create a new user")
		println("3. Exit")
		print("Please enter your choice (1-3): ")
		_, err := fmt.Scanf("%d", &userchoice1)
		if err != nil {
			println("Invalid input. Please enter a number between 1 and 3.")
			return
		}
		switch userchoice1 {
		case 1:
			username, err := Login(&users)
			if err != nil {
				println(err.Error())
				continue
			}
			passed = true
			println("Welcome,", username)
			loggedinUser = username
			break
		case 2:
			username, err := CreateUser(&users)
			if err != nil {
				println(err.Error())
				continue
			}
			passed = true
			println("User created successfully! Welcome,", username)
			loggedinUser = username
			break
		case 3:
			println("Exiting the application. Goodbye!")
			return
		default:
			println("Invalid choice. Please enter a number between 1 and 3.")
		}
		UserStorage.Save(users)

	}
	todos := Datas{}
	storage := NewStorage[Datas]("todos.json")
	storage.Load(&todos)
	reader := bufio.NewReader(os.Stdin)
	for userChoice != 6 {
		println("1. Add a new todo")
		println("2. Edit an existing todo")
		println("3. Delete a todo")
		println("4. Complete a todo")
		println("5. Display all todos")
		println("6. Exit")
		print("Please enter your choice (1-6): ")
		_, err := fmt.Scanf("%d", &userChoice)
		if err != nil {
			println("Invalid input. Please enter a number between 1 and 6.")
			return
		}
		switch userChoice {
		case 1:
			println("You chose to add a new todo.")

			print("Please enter the title of the new todo: ")
			title, err := reader.ReadString('\n')
			if err != nil {
				println("Invalid input. Please try again.")
				return
			}
			println("Please Enter the priority of the todo (1(low)-4(high)): ")
			var priority int
			_, err = fmt.Scanf("%d", &priority)
			if err != nil || priority < 1 || priority > 5 {
				println("Invalid input. Please enter a valid priority.")
				return
			}
			todos.AddTodo(title, priority, loggedinUser)
		case 2:
			println("You chose to edit an existing todo.")
			print("Please enter the index of the todo to edit: ")
			var index int
			_, err := fmt.Scanf("%d", &index)
			if err != nil {
				println("Invalid input. Please enter a valid index.")
				return
			}
			println("Please enter the new title for the todo: ")
			title, err := reader.ReadString('\n')
			if err != nil {
				println("Invalid input. Please try again.")
				return
			}
			if err != nil {
				println("Invalid input. Please try again.")
				return
			}
			todos.EditTodo(index, title, loggedinUser)
		case 3:
			println("You chose to delete a todo.")
			print("Please enter the index of the todo to delete: ")
			var index int
			_, err := fmt.Scanf("%d", &index)
			if err != nil {
				println("Invalid input. Please enter a valid index.")
				return
			}
			todos.DeleteTodo(index, loggedinUser)
		case 4:
			println("You chose to complete a todo.")
			print("Please enter the index of the todo to complete: ")
			var index int
			_, err := fmt.Scanf("%d", &index)
			if err != nil {
				println("Invalid input. Please enter a valid index.")
				return
			}
			todos.CompleteTodo(index, loggedinUser)
		case 5:
			println("You chose to display all todos.")
			todos.DisplayTodos(loggedinUser)
		case 6:
			println("Exiting the application. Goodbye!")
		default:
			println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}

	// cmdFlags := NewCommandFlags()
	// cmdFlags.Execute(&todos)
	storage.Save(todos)
}
