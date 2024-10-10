package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json video ")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	LcoCourses := []course{
		{"reactJs Bootcamp", 299, "yagyagoel1.dev", "abjdfjkdsk", []string{"web-dev", "js"}},
		{"AngularJs Bootcamp", 199, "yagyagoel1.dev", "abjdfjkdsk", nil},
		{"mern bootcamp", 299, "yagyagoel1.dev", "abjdfjkdsk", []string{"fullstack", "js"}},
	}
	//package this data as json data
	finalJson, err := json.MarshalIndent(LcoCourses, "", "\t")
	// finalJson, err := json.Marshal(LcoCourses)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "reactJs Bootcamp",
                "Price": 299,
                "website": "yagyagoel1.dev",
                "tags": ["web-dev","js"]
        }
				`)

	var LcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &LcoCourses)
		fmt.Printf("%#v\n", LcoCourses)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
