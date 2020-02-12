package AddToDo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/higashi000/sizuku/GetToDo"
	"github.com/higashi000/sizuku/ToDo"
)

func NewToDo() {
	var newTask ToDo.ToDoData
	var todolist []ToDo.ToDoData
	fmt.Print("Task Title >> ")
	fmt.Scanf("%s", &newTask.Name)
	fmt.Print("Task Details >> ")
	fmt.Scanf("%s", &newTask.Details)
	fmt.Print("Task Limit (YYYY/MM/DD) >> ")
	fmt.Scanf("%s", &newTask.Limit)

	if newTask.Limit == "" {
		newTask.Limit = "none"
	}

	filePath := os.Getenv("HOME") + "/.sizukuToDo.json"

	if GetToDo.CheckExistence(filePath) {
		fp, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0664)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		nowFileStr, err := ioutil.ReadAll(fp)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(nowFileStr), &todolist)
		if err != nil {
			log.Fatal(err)
		}
	}

	todolist = append(todolist, newTask)

	fmt.Println(todolist)

	todolistJson, err := json.MarshalIndent(todolist, "", "  ")
	err = ioutil.WriteFile(filePath, todolistJson, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
