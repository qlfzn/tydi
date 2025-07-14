package main

import (
	"fmt"
	"log"

	"github.com/qlfzn/tydi/internal/organiser"
)

func main() {
	f := organiser.File{
		CurrDir: "/Users/qlfzn/Downloads/test_run/",
	}

	files, err := f.GetAllFilesInDir(f.CurrDir)
	if err != nil {
		log.Fatal(err)
	}

	// unique := f.GetCountUniqueFileTypes(files)
	unique := f.GroupByPrefix(files)

	fmt.Println("Welcome, user! Let's organise with tydi")
	f.ShowExtCount(unique)
}
