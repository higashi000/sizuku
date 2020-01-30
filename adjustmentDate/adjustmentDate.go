package adjustmentDate

import "github.com/nsf/termbox-go"

func YearSet(whichOpe bool, year *int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if whichOpe {
		*year++
	} else {
		*year--
	}
}

func MonthSet(whichOpe bool, month *int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if whichOpe {
		if *month == 12 {
			*month = 1
		} else {
			*month++
		}
	} else {
		if *month == 1 {
			*month = 12
		} else {
			*month--
		}
	}
}

func DaySet(whichOpe bool, day *int, month int, isLeap bool) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if whichOpe {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			if *day == 31 {
				*day = 1
			} else {
				*day++
			}
		case 4, 6, 9, 11:
			if *day == 30 {
				*day = 1
			} else {
				*day++
			}
		case 2:
			if isLeap {
				if *day == 29 {
					*day = 1
				} else {
					*day++
				}
			} else {
				if *day == 28 {
					*day = 1
				} else {
					*day++
				}
			}
		}
	} else {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			if *day == 1 {
				*day = 31
			} else {
				*day++
			}
		case 4, 6, 9, 11:
			if *day == 1 {
				*day = 30
			} else {
				*day++
			}
		case 2:
			if isLeap {
				if *day == 1 {
					*day = 29
				} else {
					*day++
				}
			} else {
				if *day == 1 {
					*day = 28
				} else {
					*day++
				}
			}
		}
	}
}

func HourSet(whichOpe bool, hour *int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if whichOpe {
		if *hour == 23 {
			*hour = 0
		} else {
			*hour++
		}
	} else {
		if *hour == 0 {
			*hour = 23
		} else {
			*hour--
		}
	}
}

func MinuteSet(whichOpe bool, minute *int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if whichOpe {
		if *minute == 59 {
			*minute = 0
		} else {
			*minute++
		}
	} else {
		if *minute == 0 {
			*minute = 59
		} else {
			*minute--
		}
	}
}
