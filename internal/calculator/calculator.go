package calculator

import (
	"billme/internal/holidays"
	"time"
)

func CountWorkingDays(month, year int) int {
	return CountWorkingDaysWithHolidays(month, year, "", false)
}

func CountWorkingDaysWithHolidays(month, year int, country string, excludeHolidays bool) int {
	return CountWorkingDaysWithHolidaysAndVacation(month, year, country, excludeHolidays, 0)
}

func CountWorkingDaysWithHolidaysAndVacation(month, year int, country string, excludeHolidays bool, vacationDays int) int {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)

	var holidayList []holidays.Holiday
	if excludeHolidays && country != "" {
		provider := holidays.GetProvider(country)
		holidayList = provider.GetHolidays(year)
	}

	workingDays := 0

	for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
		weekday := day.Weekday()
		if weekday >= time.Monday && weekday <= time.Friday {
			if excludeHolidays && holidays.IsHoliday(day, holidayList) {
				continue
			}
			workingDays++
		}
	}

	// Subtract vacation days, but don't go below 0
	workingDays -= vacationDays
	if workingDays < 0 {
		workingDays = 0
	}

	return workingDays
}
