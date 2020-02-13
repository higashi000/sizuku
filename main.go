package main

import (
	"flag"
	"log"

	"github.com/higashi000/sizuku/AddToDo"
	"github.com/higashi000/sizuku/gui"
)

var (
	viewToDo = flag.Bool("s", false, "view todolist")
	addToDo  = flag.Bool("a", false, "add new todo")
)

func main() {
	flag.Parse()

	if *addToDo {
		AddToDo.NewToDo()
	}

	if *viewToDo {
		err := gui.New().Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
