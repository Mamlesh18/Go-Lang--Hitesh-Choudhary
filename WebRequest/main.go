package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://mamlesh18.github.io/mamleshva_portfolio/"

func main() {
	fmt.Println("Web Rqeuset")
	fmt.Println(HttpHand(url))
}

func HttpHand(url string) string {

	response, err := http.Get(url)
	errorHandling(err)
	databytes, err := ioutil.ReadAll(response.Body)
	errorHandling(err)

	return string(databytes)
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
