package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all languages:", languages)
	fmt.Println("JS shorts for :", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages:", languages)

	//loops are intersitng in golang

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
