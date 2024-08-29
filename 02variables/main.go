package main

import "fmt"

const LoginToken string = "sibrafd" // public
func main() {
	var username string = "Yagya"
	fmt.Println(username)
	fmt.Printf("variable is of type :%T \n", username)
	// boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("varibale is of type:%T \n", isLoggedIn)

	// smallval
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("varibale is of type:%T \n", smallVal)

	// smallfloat
	var smallFloat float32 = 255.343424344324243
	fmt.Println(smallFloat)
	fmt.Printf("varibale is of type:%T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type:%T \n", anotherVariable)

	//implicit type
	var website = "yagyagoel.dev"
	fmt.Println(website)

	//no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)

}
