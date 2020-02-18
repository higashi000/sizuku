package DeleteToDo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/higashi000/sizuku/GetToDo"
	"github.com/higashi000/sizuku/ToDo"
)

func CheckTimeLimit(date string) bool {
	now := time.Now()

	parseDate := strings.Split(date, "/")

	limitYear, err := strconv.Atoi(parseDate[0])
	if err != nil {
		log.Fatal(err)
	}
	limitMonth, err := strconv.Atoi(parseDate[1])
	if err != nil {
		log.Fatal(err)
	}
	limitDay, err := strconv.Atoi(parseDate[2])
	if err != nil {
		log.Fatal(err)
	}

	timeLimit := time.Date(limitYear, time.Month(limitMonth), limitDay, 23, 59, 59, 0, time.UTC)

	if now.Before(timeLimit) {
		return false
	}

	return true
}

func DeleteOverLimit() {
	filePath := os.Getenv("HOME") + `/.sizukuToDo.json`

	if !GetToDo.CheckExistence(filePath) {
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

	var todolist []ToDo.ToDoData

	strData := string(data)

	err = json.Unmarshal([]byte(strData), &todolist)
	if err != nil {
		log.Fatal(err)
	}

	var newtodolist []ToDo.ToDoData

	for _, e := range todolist {
		if !CheckTimeLimit(e.Limit) {
			newtodolist = append(newtodolist, e)
		}
	}

	todolistJson, err := json.MarshalIndent(newtodolist, "", "  ")
	err = ioutil.WriteFile(filePath, todolistJson, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
