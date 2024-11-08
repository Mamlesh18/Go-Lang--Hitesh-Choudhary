package main

import "fmt"

type user struct {
	Name  string
	Email string
	Age   int
}

func main() {

	mamlesh := user{"mamlesh", "mamleshgmail.com", 12}
	fmt.Println(mamlesh)
	fmt.Println("%+v", mamlesh)
	fmt.Println("%v", mamlesh.Name)

}
