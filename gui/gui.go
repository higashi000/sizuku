package gui

import (
	"github.com/higashi000/sizuku/GetToDo"
	"github.com/rivo/tview"
)

type Gui struct {
	App       *tview.Application
	Pages     *tview.Pages
	ToDoPanel *ToDoListPannel
}

func New() *Gui {
	return &Gui{
		App:       tview.NewApplication(),
		Pages:     tview.NewPages(),
		ToDoPanel: NewToDoListPannel(),
	}
}

func (g *Gui) Run() error {
	todolist, err := GetToDo.GetToDoList()
	if err != nil {
		return err
	}

	g.ToDoPanel.UpdateToDo(todolist)

	grid := tview.NewGrid().SetRows(0).
		AddItem(g.ToDoPanel, 0, 0, 1, 1, 0, 0, true)

	g.Pages.AddAndSwitchToPage("main", grid, true)

	return g.App.SetRoot(g.Pages, true).Run()
}
