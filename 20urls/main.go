package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://go.dev/"

func main() {
	fmt.Println("welcome to handlling URLs in golang")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T \n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("params is:", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "go.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
