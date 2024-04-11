package main

import "fmt"

// Capital lettets desides public or private
const jwtToken int = 300000                        // PRIVATE VAR
const JwtTokenPub string = "##kjhkjahdj#kjhkjadkj" // PUBLIC VAR

func main() {
	var username string = "Abhinav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.123456789123456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// defaults values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "google.com"
	fmt.Println(website)

	// no var style
	numverOfUsers := 300000
	fmt.Println(numverOfUsers)

	//const
	fmt.Println(jwtToken)
	fmt.Println(JwtTokenPub)

}
