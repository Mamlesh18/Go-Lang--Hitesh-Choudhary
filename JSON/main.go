package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Password string
	Tage     []string
}

func main() {
	fmt.Println("Starting JSON")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	total := []course{
		{"React", 100, "mamlesh", []string{"react", "web dev"}},
		{"GO", 200, "mamleasdsh", []string{"GO", "backend dev"}},
	}

	finalJson, err := json.MarshalIndent(total, "", "\t")
	ErrorHandling(err)
	fmt.Printf("%s\n", string(finalJson))

}

func DecodeJson() {

	TotalJson := []byte(` {
                "Name": "GO",
                "Price": 200,
                "Password": "mamleasdsh",
                "Tage": ["GO","backend dev"]
        }`)
	var lcocourse course
	checkValid := json.Valid(TotalJson)
	if checkValid {
		fmt.Println("JSON was Valid")
		err := json.Unmarshal(TotalJson, &lcocourse)
		ErrorHandling(err)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("ERROR")
	}

}

func ErrorHandling(err error) {
	if err != nil {
		panic(err)
	}

}
