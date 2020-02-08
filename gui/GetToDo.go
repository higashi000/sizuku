package gui

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ToDo struct {
	Title   string `json:"name"`
	Details string `json:"details"`
	Limit   string `json:"limit"`
}

func GetToDo() ([]ToDo, error) {
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

	var todolist []ToDo

	strData := string(data)

	err = json.Unmarshal([]byte(strData), &todolist)
	if err != nil {
		log.Fatal(err)
	}

	fp.Close()
	return todolist, err
}
