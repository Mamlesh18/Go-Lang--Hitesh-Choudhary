package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://localhost:3000/items"

func main() {
	fmt.Println("URL IN GOLANG")
	fmt.Println(GetRequest())
	fmt.Println(PostRequest())
	fmt.Println(GetRequest())
}

func GetRequest() string {
	response, err := http.Get(url)
	errorHandling(err)

	defer response.Body.Close()
	fmt.Println("Status Code", response.StatusCode)

	content, err := ioutil.ReadAll(response.Body)
	errorHandling(err)
	return string(content)
}

func PostRequest() string {
	RequestBody := strings.NewReader(`{
		"name": "Mamleshva"
	}`)

	response, err := http.Post(url, "application/json", RequestBody)
	errorHandling(err)
	fmt.Println(GetRequest())
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	errorHandling(err)
	return string(content)
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
