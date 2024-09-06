package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	result := adder(3, 5)

	fmt.Println("result is:", result)

	proRes, _ := proAdder(2, 5, 8, 7)
	fmt.Println("pro result is:", proRes)

	// greeterTwo()
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "HI PRO RESULT FUNCTION"
}

// func greeterTwo(){
// 	fmt.Println("another method")
// }
// func (){

// }()
func greeter() {
	fmt.Println("namaste from golang")
}
