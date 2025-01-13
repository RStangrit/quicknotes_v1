package main

import (
	"fmt"
	"main/errHandler"
	"main/inputReader"
	"main/mainMenuNavigation"
	"main/mainMenuOpts"
)

func handleUserChoice() {
	choice, err := inputReader.GetUserInput("Choose your action: ")
	if err != nil {
		errHandler.HandleError(err)
		return
	}

	mainMenuNavigation.Selector(choice)
}

func showMenu() {
	fmt.Println("Welcome to the Note storage")
	for {
		mainMenuOpts.Show()

		handleUserChoice()
	}
}

func main() {
	showMenu()
}
