package main

import (
	"fmt"
)

func main() {
	var arr = []string{"os", "db", "cn", "hi", "hey"}
	fmt.Println(arr)

	for d := 0; d < len(arr); d++ {
		if arr[d] == "cn" {
			continue
		}
		fmt.Println("arr - > %v", arr[d])
	}
}
