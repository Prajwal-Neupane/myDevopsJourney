package main

import (
	"fmt" //format package
	// "strings"
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

	// var userNumber int
	// fmt.Println("Enter number of users you want: ")
	// fmt.Scan(&userNumber)

	//  allUsers:= []string{}

	// for i := 0; i < userNumber; i++ {
	// 	fmt.Printf("Enter %v user full name	: ", i+1)
	// 	fmt.Scan(&allUsers[i])
		
	// }

	// To check for the valid user FirstName, LastName and Email
	// validUser := len(firstName) >= 2 && len(lastName) >= 2
	// validEmail := strings.Contains(email, "@")


	// if validUser && validEmail {
	// 	fmt.Println("User Input Data is True")
		
	// } else {
	// 	fmt.Println("The statement is not true")
	// }


	
	
	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <=i ; j++ {
	// 		fmt.Printf("*")
			
	// 	}
	// 	fmt.Printf("\n")
		
	// }
	

	// fmt.Println(allUsers)

	var a int
	var b int
	var c string

	fmt.Println("Enter the value of a: ")
	fmt.Scan(&a)
	fmt.Println("Enter the value of b: ")
	fmt.Scan(&b)
	fmt.Println("Enter the operator: +, -, *, /  : ")
	fmt.Scan(&c)

	switch c {
	case "+":
		fmt.Printf("The sum of %v and %v is %v", a, b, a+b)
		
	case "-":
		fmt.Printf("The difference of %v and %v is %v", a, b, a-b)
		
	case "*":
		fmt.Printf("The product of %v and %v is %v", a, b, a*b)
		
	case "/":
		fmt.Printf("The quotient of %v and %v is %v", a, b, a/b)
		
	 default:
		fmt.Printf("Error in selecting operator")

		
	}
	
	

	

	
	
}

