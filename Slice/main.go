package main

import "fmt"

func main() {

	var sliceing = []int{1, 2, 3, 4, 5}
	fmt.Println("arr-> %T\n", sliceing)

	sl := append(sliceing, 7, 8)
	fmt.Println("arr-> %T\n", sl)

	sl = append(sl[:3])
	fmt.Println("arr-> %T\n", sl)

	// REMOVE A INDEX
	var courses = []string{"ai", "ml", "cn", "dbms", "os"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
