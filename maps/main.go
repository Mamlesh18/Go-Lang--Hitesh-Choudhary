package main

import "fmt"

func main() {
	var mapping = make(map[string]int)
	mapping["JS"] = 1
	mapping["JAS"] = 2
	mapping["SAS"] = 3

	fmt.Println(mapping)
	fmt.Println(mapping["JS"])

	for key, value := range mapping {
		fmt.Printf("key - %v value - %v", key, value)

	}

}
