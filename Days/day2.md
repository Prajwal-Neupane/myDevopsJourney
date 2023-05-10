Today I started my devOPs journey with the same Golang tutorial of Techworld with Nana video.

Some of the basic keypoints of Golang are..

Printf() function is actually like C Programming and %v is used as a hold of the value..

%T for knowing the type of the variable.

uint variable is specifically for unsigned integer. So the number assigned to that variable must not be less than 0.
For eg. var myNumber uint = -6 // It is not applicable.. The value should be positive


Well the beginner way of declaring variable is like : var myName = "Prajwal Neupane"

Well the mentos way of declaring variable is like: myName:= "Prajwal Neupane"

The way to take user input is as;;

var userName string
var userAge int

fmt.Print("Enter your name: ")
fmt.Scan(&userName)
fmt.Print("Enter your age: ")
fmt.Scan(&userAge)

Another key take away is that;

fmt.Print(userName) // Prints the value of the variable
fmt.Print(&userName) // Prints the address of the variable where the value is stored in the memory


Learned about pointer in it and didn't do much in the night time. 
So, This much only for day 2...