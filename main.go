package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the city of New York.")
	err := os.Remove("save.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func execute() {
	//create a character
	//create a city
	//create a district
	//create a district
	main()
	//basegame.Player.Character()
}
