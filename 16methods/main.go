package main

import "fmt"

func main() {
	fmt.Println("welcome to structs")

	//no inheritence in golang no super or parent

	yagya := User{"Yagya", "yagya@go.dev", true, 16}

	fmt.Println(yagya)
	fmt.Printf("yagya details are : %+v\n", yagya)
	fmt.Printf("yagya details are : %v\n", yagya)
	fmt.Printf("Name is %v and email is  %v \n", yagya.Name, yagya.Email)
	yagya.GetStatus()
	yagya.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}
