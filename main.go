package main

import (
	"fmt"
	"log"

	"github.com/qlfzn/tydi/internal/organiser"
)

func main() {
	f := organiser.File{}

	files, err := f.GetAllFilesInDir("/Users/qlfzn/Downloads/")
	if err != nil {
		log.Fatal(err)
	}

	unique := f.GetCountUniqueFileTypes(files)

	fmt.Println("Welcome, user! Let's organise with tydi")
	f.ShowExtCount(unique)
}
