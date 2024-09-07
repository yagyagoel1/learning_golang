package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	myDefer()

}

//output hello Two one world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//outpput hello 4 3 2 1 0 two one world
