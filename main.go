package main

import (
	"log"

	"github.com/higashi000/sizuku/adjustmentDate"
	"github.com/higashi000/sizuku/dispDate"
	"github.com/nsf/termbox-go"
)

type ToDo struct {
	StartDay   string `json:"startday"`
	FinishDay  string `json:"finishday"`
	StartTime  string `json:"starttime"`
	FinishTime string `json:"finishtime"`
}

func main() {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	date := dispDate.Day{0, 2020, 1, 30, 0, 0}

MAINLOOP:
	for {
		dispDate.DispDays(0, 0, date)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP

			case termbox.KeyCtrlL:
				date.Emphasis = (date.Emphasis + 1) % 5

			case termbox.KeyCtrlK:
				switch date.Emphasis {
				case 0:
					adjustmentDate.YearSet(true, &date.Year)
				case 1:
					adjustmentDate.MonthSet(true, &date.Month)
				case 2:
					adjustmentDate.DaySet(true, &date.Day, date.Month, checkLeap(date.Year))
				case 3:
					adjustmentDate.HourSet(true, &date.Hour)
				case 4:
					adjustmentDate.MinuteSet(true, &date.Minute)
				}

			case termbox.KeyCtrlJ:
				switch date.Emphasis {
				case 0:
					adjustmentDate.YearSet(false, &date.Year)
				case 1:
					adjustmentDate.MonthSet(false, &date.Month)
				case 2:
					adjustmentDate.DaySet(false, &date.Day, date.Month, checkLeap(date.Year))
				case 3:
					adjustmentDate.HourSet(false, &date.Hour)
				case 4:
					adjustmentDate.MinuteSet(false, &date.Minute)
				}
			}

		default:
			dispDate.DispDays(0, 0, date)
		}
	}
}

func checkLeap(year int) bool {
	return true
}
