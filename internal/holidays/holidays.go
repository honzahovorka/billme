package holidays

import (
	"time"
)

type Holiday struct {
	Name string
	Date time.Time
}

type HolidayProvider interface {
	GetHolidays(year int) []Holiday
}

// calculateEaster computes the date of Easter Sunday for a given year
// using the Gregorian calendar algorithm (Anonymous Gregorian algorithm).
//
// Easter is defined as the first Sunday after the first full moon
// occurring on or after the spring equinox (March 21).
//
// This algorithm works for years 1583-4099 in the Gregorian calendar.
//
// Alternative simpler approach would be to use a lookup table for common years:
//
//	var easterDates = map[int]string{
//	    2024: "2024-03-31", 2025: "2025-04-20", 2026: "2026-04-05", ...
//	}
//
// But this algorithm is more flexible and works for any year in the valid range.
func calculateEaster(year int) time.Time {
	goldenNumber := year % 19
	century := year / 100
	yearInCentury := year % 100

	centuryLeapCorrection := century / 4
	centuryRemainder := century % 4

	moonOrbitCorrection := (century + 8) / 25
	moonCorrectionAdjustment := (century - moonOrbitCorrection + 1) / 3

	epact := (19*goldenNumber + century - centuryLeapCorrection - moonCorrectionAdjustment + 15) % 30

	yearLeapCorrection := yearInCentury / 4
	yearRemainder := yearInCentury % 4

	weekdayCorrection := (32 + 2*centuryRemainder + 2*yearLeapCorrection - epact - yearRemainder) % 7
	monthCorrection := (goldenNumber + 11*epact + 22*weekdayCorrection) / 451

	monthAndDaySum := epact + weekdayCorrection - 7*monthCorrection + 114
	month := monthAndDaySum / 31
	day := (monthAndDaySum % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

type CzechHolidayProvider struct{}

func (p *CzechHolidayProvider) GetHolidays(year int) []Holiday {
	holidays := []Holiday{
		{Name: "Nový rok", Date: time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "Svátek práce", Date: time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC)},
		{Name: "Den vítězství", Date: time.Date(year, 5, 8, 0, 0, 0, 0, time.UTC)},
		{Name: "Den slovanských věrozvěstů Cyrila a Metoděje", Date: time.Date(year, 7, 5, 0, 0, 0, 0, time.UTC)},
		{Name: "Den upálení mistra Jana Husa", Date: time.Date(year, 7, 6, 0, 0, 0, 0, time.UTC)},
		{Name: "Den české státnosti", Date: time.Date(year, 9, 28, 0, 0, 0, 0, time.UTC)},
		{Name: "Den vzniku samostatného československého státu", Date: time.Date(year, 10, 28, 0, 0, 0, 0, time.UTC)},
		{Name: "Den boje za svobodu a demokracii", Date: time.Date(year, 11, 17, 0, 0, 0, 0, time.UTC)},
		{Name: "Štědrý den", Date: time.Date(year, 12, 24, 0, 0, 0, 0, time.UTC)},
		{Name: "1. svátek vánoční", Date: time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC)},
		{Name: "2. svátek vánoční", Date: time.Date(year, 12, 26, 0, 0, 0, 0, time.UTC)},
	}

	holidays = append(holidays, p.getEasterMonday(year))

	return holidays
}

func (p *CzechHolidayProvider) getEasterMonday(year int) Holiday {
	easter := calculateEaster(year)
	easterMonday := easter.AddDate(0, 0, 1)
	return Holiday{Name: "Velikonoční pondělí", Date: easterMonday}
}

func GetProvider(country string) HolidayProvider {
	return &CzechHolidayProvider{}
}

func IsHoliday(date time.Time, holidays []Holiday) bool {
	for _, holiday := range holidays {
		if holiday.Date.Year() == date.Year() &&
			holiday.Date.Month() == date.Month() &&
			holiday.Date.Day() == date.Day() {
			return true
		}
	}
	return false
}
