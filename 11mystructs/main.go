package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

	//no inheritence in golang no super or parent

	yagya := User{"Yagya", "yagya@go.dev", true, 16}

	fmt.Println(yagya)
	fmt.Printf("yagya details are : %+v\n", yagya)
	fmt.Printf("yagya details are : %v\n", yagya)
	fmt.Printf("Name is %v and email is  %v", yagya.Name, yagya.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
