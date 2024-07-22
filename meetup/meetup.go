package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

var Teenth WeekSchedule = 99
var First WeekSchedule = 1
var Second WeekSchedule = 2
var Third WeekSchedule = 3
var Fourth WeekSchedule = 4
var Last WeekSchedule = -1

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	
	if wSched == Last {
		date = time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1)
	}
	dayCount := 0

	for {
		if date.Weekday() == wDay {
			if wSched == Last {
				dayCount--
			} else {
				dayCount++
			}
		}

		if dayCount == int(wSched) {
			return date.Day()
		}

		if wSched == Teenth && date.Day() >= 13 && date.Day() <= 19 && date.Weekday() == wDay {
			return date.Day()
		}
		
		if wSched == Last {
			date = date.AddDate(0, 0, -1)
		} else {
			date = date.AddDate(0, 0, 1)
		}
	}
}
