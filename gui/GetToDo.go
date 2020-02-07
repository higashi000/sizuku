package gui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ToDoList struct {
	ToDo []struct {
		Name      string `json:"name"`
		Details   string `json:"details"`
		StartTime string `json:"starttime"`
		EndTime   string `json:"endtime"`
	} `json:"todo"`
}

func GetToDo() (ToDoList, error) {
	homeEnv := os.Getenv("HOME")
	fp, err := os.Open(homeEnv + `/.sizukuToDo.json`)
	if err != nil {
		log.Fatal(err)
	}

	fp.Seek(0, 0)
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	var todolist ToDoList

	strData := string(data)

	err = json.Unmarshal([]byte(strData), &todolist)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(todolist.ToDo[0].Name)
	fmt.Println(todolist.ToDo[0].Details)

	fp.Close()
	return ToDoList{nil}, err
}
