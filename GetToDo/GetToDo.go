package GetToDo

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

func GetToDoList() ([]ToDo, error) {
	filePath := os.Getenv("HOME") + `/.sizukuToDo.json`

	if !CheckExistence(filePath) {
		fp, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		fp.WriteString("[]")
		fp.Close()
	}

	fp, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

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

	return todolist, err
}

func CheckExistence(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}
