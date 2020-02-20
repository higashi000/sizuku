package AddToDo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/higashi000/sizuku/GetToDo"
	"github.com/higashi000/sizuku/ToDo"
)

func NewToDo() {
	var sc = bufio.NewScanner(os.Stdin)
	var newTask ToDo.ToDoData
	var todolist []ToDo.ToDoData

	fmt.Print("Task Title >> ")
	if sc.Scan() {
		newTask.Name = sc.Text()
	}

	fmt.Print("Task Details >> ")
	if sc.Scan() {
		newTask.Details = sc.Text()
	}

	for {
		fmt.Print("Task Limit YYYY/MM/DD >> ")
		fmt.Scanf("%s", &newTask.Limit)

		if CheckYYYYMMDD(newTask.Limit) || newTask.Limit == "" {
			break
		}
	}

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

	todolistJson, err := json.MarshalIndent(todolist, "", "  ")
	err = ioutil.WriteFile(filePath, todolistJson, 0664)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckYYYYMMDD(date string) bool {
	parseDate := strings.Split(date, "/")

	if (len(parseDate) > 3) || (1 > len(parseDate)) {
		return false
	} else {
		if (len(parseDate[0]) == 4) && (len(parseDate[1]) == 2) && (len(parseDate[2]) == 2) {
			return true
		} else {
			return false
		}
	}
}
