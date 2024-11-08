package main

import (
	"Mongo/controller"
	model "Mongo/models"
	"fmt"
)

func main() {

	fmt.Println("Mongo DB Starting")

	sample := model.Netflix{
		Movie:   "inceprito",
		Watched: true,
	}

	controller.InsertMovie(sample)

}
