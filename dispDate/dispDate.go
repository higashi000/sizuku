package dispDate

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

type Day struct {
	Emphasis int
	Year     int
	Month    int
	Day      int
	Hour     int
	Minute   int
}

func DispDays(x, y int, date Day) {
	nowPos := 0

	yearRunes := []rune(strconv.Itoa(date.Year))
	monthRunes := []rune(strconv.Itoa(date.Month))
	dayRunes := []rune(strconv.Itoa(date.Day))
	hourRunes := []rune(strconv.Itoa(date.Hour))
	minuteRunes := []rune(strconv.Itoa(date.Minute))

	for i := 0; i < len(yearRunes); i++ {
		if date.Emphasis == 0 {
			termbox.SetCell(x+nowPos, y, yearRunes[i], termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(x+nowPos, y, yearRunes[i], termbox.ColorDefault, termbox.ColorDefault)
		}
		nowPos++
	}

	termbox.SetCell(x+nowPos, y, '/', termbox.ColorDefault, termbox.ColorDefault)
	nowPos++

	for i := 0; i < len(monthRunes); i++ {
		if date.Emphasis == 1 {
			termbox.SetCell(x+nowPos, y, monthRunes[i], termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(x+nowPos, y, monthRunes[i], termbox.ColorDefault, termbox.ColorDefault)
		}
		nowPos++
	}

	termbox.SetCell(x+nowPos, y, '/', termbox.ColorDefault, termbox.ColorDefault)
	nowPos++

	for i := 0; i < len(dayRunes); i++ {
		if date.Emphasis == 2 {
			termbox.SetCell(x+nowPos, y, dayRunes[i], termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(x+nowPos, y, dayRunes[i], termbox.ColorDefault, termbox.ColorDefault)
		}
		nowPos++
	}

	termbox.SetCell(x+nowPos, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	nowPos++

	for i := 0; i < len(hourRunes); i++ {
		if date.Emphasis == 3 {
			termbox.SetCell(x+nowPos, y, hourRunes[i], termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(x+nowPos, y, hourRunes[i], termbox.ColorDefault, termbox.ColorDefault)
		}
		nowPos++
	}

	termbox.SetCell(x+nowPos, y, ':', termbox.ColorDefault, termbox.ColorDefault)
	nowPos++

	for i := 0; i < len(minuteRunes); i++ {
		if date.Emphasis == 4 {
			termbox.SetCell(x+nowPos, y, minuteRunes[i], termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(x+nowPos, y, minuteRunes[i], termbox.ColorDefault, termbox.ColorDefault)
		}
		nowPos++
	}

	termbox.Flush()
}
