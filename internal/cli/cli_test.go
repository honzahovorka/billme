package cli

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFormatOutput(t *testing.T) {
	tests := []struct {
		name        string
		workingDays int
		config      *Config
		expected    string
	}{
		{
			name:        "Invoice ready format",
			workingDays: 22,
			config:      &Config{InvoiceReady: true},
			expected:    "22",
		},
		{
			name:        "Ka-ching format",
			workingDays: 22,
			config:      &Config{KaChing: true},
			expected:    "22 days = CHA-CHING! 🤑",
		},
		{
			name:        "Verbose format",
			workingDays: 22,
			config:      &Config{Verbose: true, Month: 7, Year: 2024},
			expected:    "July 2024: 22 billable days 💸",
		},
		{
			name:        "Default format",
			workingDays: 22,
			config:      &Config{},
			expected:    "💰 22",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatOutput(tt.workingDays, tt.config)
			if result != tt.expected {
				t.Errorf("FormatOutput() = %q; want %q", result, tt.expected)
			}
		})
	}
}

func TestConfigDefaults(t *testing.T) {
	config := &Config{}

	if config.Verbose {
		t.Error("Expected Verbose to be false by default")
	}
	if config.KaChing {
		t.Error("Expected KaChing to be false by default")
	}
	if config.InvoiceReady {
		t.Error("Expected InvoiceReady to be false by default")
	}
	if config.Help {
		t.Error("Expected Help to be false by default")
	}
}

func TestFormatOutputPriority(t *testing.T) {
	config := &Config{
		InvoiceReady: true,
		KaChing:      true,
		Verbose:      true,
		Month:        7,
		Year:         2024,
	}

	result := FormatOutput(22, config)
	expected := "22"

	if result != expected {
		t.Errorf("FormatOutput() = %q; want %q (InvoiceReady should take priority)", result, expected)
	}
}

func TestParseArgsInvalidMonth(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"Invalid month string", []string{"abc"}},
		{"Month too low", []string{"0"}},
		{"Month too high", []string{"13"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = append([]string{"billme"}, tt.args...)
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			_, err := ParseArgs()
			if err == nil {
				t.Errorf("ParseArgs() should return error for invalid month: %v", tt.args)
			}
		})
	}
}

func TestParseArgsInvalidYear(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"billme", "7", "abc"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	_, err := ParseArgs()
	if err == nil {
		t.Error("ParseArgs() should return error for invalid year")
	}
}

func TestParseArgsTooManyArguments(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"billme", "7", "2024", "extra"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	_, err := ParseArgs()
	if err == nil {
		t.Error("ParseArgs() should return error for too many arguments")
	}
}

func TestParseArgsNoArguments(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"billme"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	config, err := ParseArgs()
	if err != nil {
		t.Errorf("ParseArgs() should not return error for no arguments: %v", err)
	}

	now := time.Now()
	if config.Month != int(now.Month()) {
		t.Errorf("Expected current month %d, got %d", int(now.Month()), config.Month)
	}
	if config.Year != now.Year() {
		t.Errorf("Expected current year %d, got %d", now.Year(), config.Year)
	}
}

func TestParseArgsOneArgument(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"billme", "7"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	config, err := ParseArgs()
	if err != nil {
		t.Errorf("ParseArgs() should not return error for one argument: %v", err)
	}

	if config.Month != 7 {
		t.Errorf("Expected month 7, got %d", config.Month)
	}
	if config.Year != time.Now().Year() {
		t.Errorf("Expected current year %d, got %d", time.Now().Year(), config.Year)
	}
}

func TestParseArgsTwoArguments(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"billme", "7", "2024"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	config, err := ParseArgs()
	if err != nil {
		t.Errorf("ParseArgs() should not return error for two arguments: %v", err)
	}

	if config.Month != 7 {
		t.Errorf("Expected month 7, got %d", config.Month)
	}
	if config.Year != 2024 {
		t.Errorf("Expected year 2024, got %d", config.Year)
	}
}

func TestParseArgsFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected Config
	}{
		{
			name:     "Verbose flag",
			args:     []string{"billme", "-v", "7", "2024"},
			expected: Config{Month: 7, Year: 2024, Verbose: true},
		},
		{
			name:     "Ka-ching flag",
			args:     []string{"billme", "-ka-ching", "7", "2024"},
			expected: Config{Month: 7, Year: 2024, KaChing: true},
		},
		{
			name:     "Invoice ready flag",
			args:     []string{"billme", "-invoice-ready", "7", "2024"},
			expected: Config{Month: 7, Year: 2024, InvoiceReady: true},
		},
		{
			name:     "Help flag",
			args:     []string{"billme", "-help"},
			expected: Config{Help: true},
		},
		{
			name:     "Multiple flags",
			args:     []string{"billme", "-v", "-ka-ching", "7", "2024"},
			expected: Config{Month: 7, Year: 2024, Verbose: true, KaChing: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = tt.args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			config, err := ParseArgs()
			if err != nil {
				t.Errorf("ParseArgs() returned error: %v", err)
				return
			}

			if config.Help != tt.expected.Help {
				t.Errorf("Expected Help %v, got %v", tt.expected.Help, config.Help)
			}
			if config.Verbose != tt.expected.Verbose {
				t.Errorf("Expected Verbose %v, got %v", tt.expected.Verbose, config.Verbose)
			}
			if config.KaChing != tt.expected.KaChing {
				t.Errorf("Expected KaChing %v, got %v", tt.expected.KaChing, config.KaChing)
			}
			if config.InvoiceReady != tt.expected.InvoiceReady {
				t.Errorf("Expected InvoiceReady %v, got %v", tt.expected.InvoiceReady, config.InvoiceReady)
			}
			if !config.Help {
				if config.Month != tt.expected.Month {
					t.Errorf("Expected Month %d, got %d", tt.expected.Month, config.Month)
				}
				if config.Year != tt.expected.Year {
					t.Errorf("Expected Year %d, got %d", tt.expected.Year, config.Year)
				}
			}
		})
	}
}

func TestParseArgsValidMonthRange(t *testing.T) {
	for month := 1; month <= 12; month++ {
		t.Run(fmt.Sprintf("Month_%d", month), func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = []string{"billme", fmt.Sprintf("%d", month), "2024"}
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			config, err := ParseArgs()
			if err != nil {
				t.Errorf("ParseArgs() should not return error for valid month %d: %v", month, err)
			}
			if config.Month != month {
				t.Errorf("Expected month %d, got %d", month, config.Month)
			}
		})
	}
}
