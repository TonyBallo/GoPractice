package main

import (
	"fmt"
	"strings"
)

func main () {
	
	switch input := strings.ToUpper(Login_o_Signup()); input {
	case "S":
		fmt.Println("Signing up...")
	case "L":
		fmt.Println("Logging in...")
	default:
		fmt.Println("Something is wrong...")
	}

}

func Login_o_Signup() string {
	var input string

	fmt.Println("Would You like to Login (L) or Signup (S)? \n>")
	fmt.Scan(&input)
	for strings.ToUpper(input) != "L" && strings.ToUpper(input) != "S" {
		fmt.Println("Please enter a valid option\n>")
		fmt.Scanf("%s", &input)
	}

	return input
}

/* Outline 

Ask for login or Signup
Login
	-get id 
	-check if exists
		-if not
			-ask for signup
		-if yes
			-"Succesfull Login"
Signup
	-get username
	-get password
	-get Fistname
	-get Lastname
	-give ID

*/

