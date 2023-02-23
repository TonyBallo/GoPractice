package main

import (
	"fmt"
	"strings"
)

// Profile structure for people 
type Profile struct {
	firstName string
	lastName string
	username string
	password string
}

func main () {
	// create the database map
	db := make(map[int]Profile)
	db[1] = Profile{
		firstName: "ant",
		lastName: "ball",
		username: "user",
		password: "pass",
	}
	switch input := strings.ToUpper(Login_o_Signup()); input {
	case "S":
		id := Signup(db)
		fmt.Printf("\nProfile Created, Your ID is %d", id)
	case "L":
		Login()
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

func Signup(database map[int]Profile) int {

	var username, password, fname, lname string

	// ask for data to make profile
	fmt.Println("##### Lets Get You Signed Up! #####\n\n")
	fmt.Println("Username : >")
	fmt.Scan(&username)
	fmt.Println("Password : >")
	fmt.Scan(&password)
	fmt.Println("First Name : >")
	fmt.Scan(&fname)
	fmt.Println("Last Name : >")
	fmt.Scan(&lname)

	// initialize new Profile with given data
	newProfile := Profile{
		username: username, 
		password: password, 
		firstName: fname, 
		lastName: lname,
	}

	// make id for profile
	var newID int 
	if len(database) != 0 {
		newID = len(database) + 1
	} else {
		newID = 0
	}

	// add profile to database
	database[newID] = newProfile

	// return the new id for made profile
	return newID
	
}

func Login() {

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

