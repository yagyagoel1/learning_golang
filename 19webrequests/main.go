package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev"

func main() {
	fmt.Println("Lco web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type %T \n", response)

	defer response.Body.Close() //calllers responsiblilty to close the connection

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
