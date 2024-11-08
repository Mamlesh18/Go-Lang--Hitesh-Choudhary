package main

import "fmt"

type user struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {

	mamlesh := user{"mamlesh", "mamleshgmail.com", 12, true}
	fmt.Println(mamlesh)
	fmt.Println("%+v", mamlesh)
	fmt.Println("%v", mamlesh.Name)
	mamlesh.Get()
}

func (u user) Get() {
	fmt.Println("is active", u.Status)
}
