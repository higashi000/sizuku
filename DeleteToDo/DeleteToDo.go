package DeleteToDo

import (
	"log"
	"strconv"
	"strings"
	"time"

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

func DeleteOverLimit(todolist []ToDo.ToDoData) []ToDo.ToDoData {
	var newtodolist []ToDo.ToDoData

	for _, e := range todolist {
		if !CheckTimeLimit(e.Limit) {
			newtodolist = append(newtodolist, e)
		}
	}

	return newtodolist
}
