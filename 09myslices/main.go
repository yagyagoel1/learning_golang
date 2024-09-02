package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "Banana")

	fmt.Println("fruitList", fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println("fruitlist", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321) //all the values get alloncated again

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted((highScores)))

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
