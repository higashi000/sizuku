package gui

import (
	"github.com/gdamore/tcell"
	"github.com/higashi000/sizuku/ToDo"
	"github.com/rivo/tview"
)

type ToDoListPannel struct {
	*tview.Table
	todolist []ToDo.ToDoData
}

func NewToDoListPannel() *ToDoListPannel {
	pannel := &ToDoListPannel{
		Table: tview.NewTable(),
	}

	pannel.SetBorder(true).SetTitle("ToDoList").SetTitleAlign(tview.AlignLeft)
	pannel.SetFixed(1, 0).SetOffset(1, 1).SetSelectable(true, false)
	return pannel
}

func (t *ToDoListPannel) UpdateToDo(todolist []ToDo.ToDoData) {
	t.todolist = todolist

	table := t.Clear()

	headers := []string{
		"title",
		"details",
		"limit",
	}

	for i, e := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            e,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorLightBlue,
			BackgroundColor: tcell.ColorDefault,
		})
	}

	for i, todo := range todolist {
		table.SetCell(i+1, 0, tview.NewTableCell(todo.Name))
		table.SetCell(i+1, 1, tview.NewTableCell(todo.Details))
		table.SetCell(i+1, 2, tview.NewTableCell(todo.Limit))
	}
}
