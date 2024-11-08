package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// STRING INPUT
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(input)
	}

	// INT INPUT
	fmt.Println("enter number")
	inputs, err := reader.ReadString('\n')
	fmt.Println("your number", inputs)
	numAdd, err := strconv.ParseFloat(strings.TrimSpace(inputs), 64)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("your new add", 1+numAdd)

	}

}
