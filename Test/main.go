package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type animal struct {
	Name  string `json:"name"` // Exported and added JSON tags for marshaling
	Count int    `json:"count"`
}

var ani_details = []animal{
	{Name: "Tiger", Count: 2},
	{Name: "Lion", Count: 3},
	{Name: "Cheetah", Count: 5},
	{Name: "Elephant", Count: 10},
}

const url = "https://mamlesh18.github.io/mamleshva_portfolio/"
const url2 = "http://localhost:3000/items"

func main() {
	fmt.Println("STARTING SERVER")
	// fmt.Println(HttpHandle(url)) // You can uncomment this to use the HttpHandle
	Posting()
}

func Posting() {
	fmt.Println("Posting data to localhost")

	// Marshal ani_details to JSON
	jsonConvert, err := json.Marshal(ani_details)
	if err != nil {
		panic(err)
	}

	// Create a POST request with the JSON data
	requestBody := bytes.NewReader(jsonConvert)
	response, err := http.Post(url2, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Read and log the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Status: %s\n", response.Status)
	fmt.Printf("Response Body: %s\n", string(responseBody))
}

func HttpHandle(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(dataBytes)
}
