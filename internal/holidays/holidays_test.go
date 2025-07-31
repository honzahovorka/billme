package holidays

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	holidays := []Holiday{
		{Name: "Test Holiday", Date: time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)},
	}

	testDate := time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)
	if !IsHoliday(testDate, holidays) {
		t.Error("Expected July 4, 2024 to be a holiday")
	}

	nonHolidayDate := time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC)
	if IsHoliday(nonHolidayDate, holidays) {
		t.Error("Expected July 5, 2024 to not be a holiday")
	}
}

func TestCzechHolidayProvider(t *testing.T) {
	provider := &CzechHolidayProvider{}
	holidays := provider.GetHolidays(2024)

	if len(holidays) != 12 {
		t.Errorf("Expected 12 Czech holidays, got %d", len(holidays))
	}

	expectedHolidays := map[string]string{
		"Nový rok":            "2024-01-01",
		"Svátek práce":        "2024-05-01",
		"Den vítězství":       "2024-05-08",
		"Štědrý den":          "2024-12-24",
		"1. svátek vánoční":   "2024-12-25",
		"2. svátek vánoční":   "2024-12-26",
		"Velikonoční pondělí": "2024-04-01", // Easter Monday 2024
	}

	for _, holiday := range holidays {
		if expected, exists := expectedHolidays[holiday.Name]; exists {
			expectedDate, _ := time.Parse("2006-01-02", expected)
			if !holiday.Date.Equal(expectedDate) {
				t.Errorf("Expected %s on %s, got %s", holiday.Name, expected, holiday.Date.Format("2006-01-02"))
			}
		}
	}
}

func TestGetProvider(t *testing.T) {
	tests := []string{"CZ", "US", "UK", "anything", ""}

	for _, country := range tests {
		provider := GetProvider(country)
		if provider == nil {
			t.Errorf("GetProvider(%s) returned nil", country)
		}

		// Should always return Czech provider now
		_, ok := provider.(*CzechHolidayProvider)
		if !ok {
			t.Errorf("GetProvider(%s) should return CzechHolidayProvider", country)
		}
	}
}

func TestEasterCalculation(t *testing.T) {
	tests := []struct {
		year     int
		expected string
	}{
		{2024, "2024-03-31"},
		{2025, "2025-04-20"},
		{2026, "2026-04-05"},
	}

	for _, tt := range tests {
		easter := calculateEaster(tt.year)
		expected, _ := time.Parse("2006-01-02", tt.expected)
		if !easter.Equal(expected) {
			t.Errorf("Easter %d: expected %s, got %s", tt.year, tt.expected, easter.Format("2006-01-02"))
		}
	}
}
