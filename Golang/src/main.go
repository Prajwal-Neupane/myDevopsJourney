package main

import "fmt" //format package




func main() {
	var i int


	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your number is: ", i)
	userInput()
}
func userInput() {

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(name)
	fmt.Println("Hello", name, "! Welcome to our website")

}
