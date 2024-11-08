package main

import (
	"fmt"
)

func main() {

	logins := 24
	var res string
	if logins < 10 {
		res = "Less"

	} else if logins == 23 {
		res = "same"

	} else {
		res = "Hight"
	}

	fmt.Println(res)
}
