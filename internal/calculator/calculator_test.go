package calculator

import (
	"fmt"
	"testing"
	"time"
)

func TestCountWorkingDays(t *testing.T) {
	tests := []struct {
		name     string
		month    int
		year     int
		expected int
	}{
		{
			name:     "January 2024",
			month:    1,
			year:     2024,
			expected: 23,
		},
		{
			name:     "February 2024 (leap year)",
			month:    2,
			year:     2024,
			expected: 21,
		},
		{
			name:     "February 2023 (non-leap year)",
			month:    2,
			year:     2023,
			expected: 20,
		},
		{
			name:     "July 2024",
			month:    7,
			year:     2024,
			expected: 23,
		},
		{
			name:     "December 2024",
			month:    12,
			year:     2024,
			expected: 22,
		},
		{
			name:     "April 2024 (30 days)",
			month:    4,
			year:     2024,
			expected: 22,
		},
		{
			name:     "May 2024 (31 days)",
			month:    5,
			year:     2024,
			expected: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWorkingDays(tt.month, tt.year)
			if result != tt.expected {
				t.Errorf("CountWorkingDays(%d, %d) = %d; want %d", tt.month, tt.year, result, tt.expected)
			}
		})
	}
}

func TestCountWorkingDaysEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		month int
		year  int
	}{
		{"Historical date", 1, 1900},
		{"Future date", 12, 2100},
		{"Minimum valid month", 1, 2024},
		{"Maximum valid month", 12, 2024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWorkingDays(tt.month, tt.year)
			if result < 0 {
				t.Errorf("CountWorkingDays should never return negative values, got %d", result)
			}
			if result > 31 {
				t.Errorf("CountWorkingDays should never return more than 31 days, got %d", result)
			}
		})
	}
}

func TestCountWorkingDaysConsistency(t *testing.T) {
	testCases := []struct {
		month int
		year  int
	}{
		{6, 2024},
		{1, 2023},
		{12, 2025},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Month_%d_Year_%d", tc.month, tc.year), func(t *testing.T) {
			firstDay := time.Date(tc.year, time.Month(tc.month), 1, 0, 0, 0, 0, time.UTC)
			lastDay := firstDay.AddDate(0, 1, -1)

			manualCount := 0
			for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
				weekday := day.Weekday()
				if weekday >= time.Monday && weekday <= time.Friday {
					manualCount++
				}
			}

			functionResult := CountWorkingDays(tc.month, tc.year)

			if functionResult != manualCount {
				t.Errorf("CountWorkingDays(%d, %d) = %d; manual count = %d", tc.month, tc.year, functionResult, manualCount)
			}
		})
	}
}

func TestCountWorkingDaysLeapYear(t *testing.T) {
	feb2024 := CountWorkingDays(2, 2024)
	feb2023 := CountWorkingDays(2, 2023)

	if feb2024 <= feb2023 {
		t.Errorf("February 2024 (leap year) should have more or equal working days than February 2023, got %d vs %d", feb2024, feb2023)
	}
}

func TestCountWorkingDaysAllMonths(t *testing.T) {
	year := 2024
	totalDays := 0

	for month := 1; month <= 12; month++ {
		days := CountWorkingDays(month, year)
		if days < 19 || days > 23 {
			t.Errorf("Month %d has unusual number of working days: %d", month, days)
		}
		totalDays += days
	}

	if totalDays < 250 || totalDays > 270 {
		t.Errorf("Total working days for year %d seems unusual: %d", year, totalDays)
	}
}

func TestCountWorkingDaysWithHolidays(t *testing.T) {
	tests := []struct {
		name            string
		month           int
		year            int
		excludeHolidays bool
		expectLess      bool
	}{
		{
			name:            "July 2024 CZ with holidays",
			month:           7,
			year:            2024,
			excludeHolidays: true,
			expectLess:      true,
		},
		{
			name:            "July 2024 CZ without holidays",
			month:           7,
			year:            2024,
			excludeHolidays: false,
			expectLess:      false,
		},
		{
			name:            "December 2024 CZ with holidays",
			month:           12,
			year:            2024,
			excludeHolidays: true,
			expectLess:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withHolidays := CountWorkingDaysWithHolidays(tt.month, tt.year, "CZ", tt.excludeHolidays)
			withoutHolidays := CountWorkingDays(tt.month, tt.year)

			if tt.expectLess && withHolidays >= withoutHolidays {
				t.Errorf("Expected working days with holidays (%d) to be less than without holidays (%d)", withHolidays, withoutHolidays)
			}

			if !tt.expectLess && withHolidays != withoutHolidays {
				t.Errorf("Expected working days to be the same when not excluding holidays: with=%d, without=%d", withHolidays, withoutHolidays)
			}
		})
	}
}
func TestCountWorkingDaysWithHolidaysSpecificCases(t *testing.T) {
	july2024WithHolidays := CountWorkingDaysWithHolidays(7, 2024, "CZ", true)
	july2024WithoutHolidays := CountWorkingDays(7, 2024)

	if july2024WithHolidays != july2024WithoutHolidays-1 {
		t.Errorf("July 2024 should have 1 less working day with Czech holidays (July 6th - Jan Hus Day), got %d vs %d", july2024WithHolidays, july2024WithoutHolidays)
	}
}

func TestCountWorkingDaysWithVacation(t *testing.T) {
	tests := []struct {
		name         string
		month        int
		year         int
		vacationDays int
		expected     int
	}{
		{
			name:         "July 2024 with 5 vacation days",
			month:        7,
			year:         2024,
			vacationDays: 5,
			expected:     18, // 23 working days - 5 vacation days
		},
		{
			name:         "July 2024 with 0 vacation days",
			month:        7,
			year:         2024,
			vacationDays: 0,
			expected:     23, // 23 working days - 0 vacation days
		},
		{
			name:         "July 2024 with excessive vacation days",
			month:        7,
			year:         2024,
			vacationDays: 30,
			expected:     0, // Should not go below 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWorkingDaysWithHolidaysAndVacation(tt.month, tt.year, "CZ", false, tt.vacationDays)
			if result != tt.expected {
				t.Errorf("Expected %d working days, got %d", tt.expected, result)
			}
		})
	}
}

func TestCountWorkingDaysWithHolidaysAndVacation(t *testing.T) {
	// July 2024: 23 working days, -1 for holiday (July 6), -3 for vacation = 19
	result := CountWorkingDaysWithHolidaysAndVacation(7, 2024, "CZ", true, 3)
	expected := 19

	if result != expected {
		t.Errorf("July 2024 with holidays and 3 vacation days: expected %d, got %d", expected, result)
	}
}
