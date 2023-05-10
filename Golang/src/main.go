package main

import (
	"fmt" //format package
)




func main() {

    // var firstName string
	// var lastName string
	// var userAge int
	// var email string
	// fmt.Println("Enter your first name: ")
	// fmt.Scan(&firstName)
	// fmt.Println("Enter your second name: ")
	// fmt.Scan(&lastName)
	// fmt.Println("Enter your email: ")
	// fmt.Scan(&email)
	// fmt.Printf("Enter your age: ")
	// fmt.Scan(&userAge)
	// fmt.Printf("Hello %v %v !! Your email is %v and age is %v", firstName,lastName, email, userAge)

	var userNumber int
	fmt.Println("Enter number of users you want: ")
	fmt.Scan(&userNumber)

	 allUsers:= []string{}

	for i := 0; i < userNumber; i++ {
		fmt.Printf("Enter %v user full name	: ", i+1)
		fmt.Scan(&allUsers[i])
		
	}

	fmt.Println(allUsers)
	
	

	

	
	
}

