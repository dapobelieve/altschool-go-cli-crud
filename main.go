package main

import (
	"fmt"
	"os"
)

var TODOLIST = []string{"Default todo"}
var menuNumber int

func main() {
	welcome()
	menu()
}

func newLine(numberOfLines int) {
	i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

func welcome() {
	fmt.Println("*******{ Welcome to the TODO CLI app 🍾 }*******")
}

func listTodos() {
	fmt.Println("List of items in your TODO list are:")
	newLine(1)
	for index, item := range TODOLIST {
		fmt.Printf("TODO #%v: %v", index+1, item)
	}
	newLine(1)
	menu()
}

func createTodo() {
	fmt.Println("Creating Todo Item")
	menu()
}

func editTodo() {
	fmt.Println("Editing todo item")
	menu()
}

func deleteTodo() {
	fmt.Println("Deleting todo item")
	menu()
}

func exitProgram() {
	fmt.Println("Goodbye 👋")
	os.Exit(0)
}

func menu() {
	newLine(2)
	fmt.Println("Select Operation:")
	fmt.Println("1. Create Todo \t\t 2. List Todo Items")
	fmt.Println("3. Edit Todo Item \t 4. Delete Todo Item")
	fmt.Println("0. Exit program \t")
	newLine(1)

	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
	}

	switch menuNumber {
	case 1:
		createTodo()
	case 2:
		listTodos()
	case 3:
		editTodo()
	case 4:
		deleteTodo()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}
