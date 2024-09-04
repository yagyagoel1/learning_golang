package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golnag")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("Dice value is 6 and you can open")
		fallthrough //make it work the same as in other lanugage without break
	default:
		fmt.Println("wHAT WAS THAT")
	}
}
