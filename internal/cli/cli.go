package cli

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

type Config struct {
	Month           int
	Year            int
	Verbose         bool
	KaChing         bool
	InvoiceReady    bool
	Help            bool
	ExcludeHolidays bool
	VacationDays    int
}

func ParseArgs() (*Config, error) {
	config := &Config{}

	// Define shared variables for flags that have both short and long forms
	var verboseFlag bool
	var helpFlag bool
	var excludeHolidaysFlag bool
	var vacationDaysFlag int

	// Short flags
	flag.BoolVar(&verboseFlag, "v", false, "verbose output")
	flag.BoolVar(&helpFlag, "h", false, "show help")
	flag.BoolVar(&excludeHolidaysFlag, "x", false, "exclude Czech public holidays")
	flag.IntVar(&vacationDaysFlag, "d", 0, "vacation/time-off days to subtract")

	// Long flags (same variables)
	flag.BoolVar(&verboseFlag, "verbose", false, "verbose output")
	flag.BoolVar(&helpFlag, "help", false, "show help")
	flag.BoolVar(&excludeHolidaysFlag, "exclude-holidays", false, "exclude Czech public holidays from working days")
	flag.IntVar(&vacationDaysFlag, "vacation-days", 0, "number of vacation/time-off days to subtract")

	// Flags that only have long forms
	kaching := flag.Bool("ka-ching", false, "celebratory output")
	invoiceReady := flag.Bool("invoice-ready", false, "clean number only")

	flag.Parse()

	config.Verbose = verboseFlag
	config.KaChing = *kaching
	config.InvoiceReady = *invoiceReady
	config.Help = helpFlag
	config.ExcludeHolidays = excludeHolidaysFlag
	config.VacationDays = vacationDaysFlag

	if config.Help {
		return config, nil
	}

	args := flag.Args()
	now := time.Now()

	if len(args) == 0 {
		config.Month = int(now.Month())
		config.Year = now.Year()
	} else if len(args) == 1 {
		month, err := strconv.Atoi(args[0])
		if err != nil || month < 1 || month > 12 {
			return nil, fmt.Errorf("invalid month: %s", args[0])
		}
		config.Month = month
		config.Year = now.Year()
	} else if len(args) == 2 {
		month, err := strconv.Atoi(args[0])
		if err != nil || month < 1 || month > 12 {
			return nil, fmt.Errorf("invalid month: %s", args[0])
		}
		year, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, fmt.Errorf("invalid year: %s", args[1])
		}
		config.Month = month
		config.Year = year
	} else {
		return nil, fmt.Errorf("too many arguments")
	}

	return config, nil
}

func ShowHelp() {
	fmt.Println("💸 BILLME - Your billable days calculator! 💸")
	fmt.Println()
	fmt.Println("Usage: billme [month] [year] [options]")
	fmt.Println()
	fmt.Println("Stop counting on your fingers - let me bill you properly!")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  billme                    # Current month")
	fmt.Println("  billme 7                  # July this year")
	fmt.Println("  billme 7 2024             # July 2024")
	fmt.Println("  billme -v 7 2024          # Verbose output")
	fmt.Println("  billme -x -d 5 7          # Exclude holidays, 5 vacation days")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -v, --verbose             Verbose output")
	fmt.Println("  -h, --help                Show this help")
	fmt.Println("  -x, --exclude-holidays    Exclude Czech public holidays from working days")
	fmt.Println("  -d, --vacation-days <num> Number of vacation/time-off days to subtract")
	fmt.Println("  --ka-ching                Celebratory output")
	fmt.Println("  --invoice-ready           Clean number only (for piping)")
}

func ShowUsage() {
	fmt.Println("Usage: billme [month] [year] [options]")
	fmt.Println("Use -help for more information")
}

func FormatOutput(workingDays int, config *Config) string {
	if config.InvoiceReady {
		return fmt.Sprintf("%d", workingDays)
	} else if config.KaChing {
		return fmt.Sprintf("%d days = CHA-CHING! 🤑", workingDays)
	} else if config.Verbose {
		monthName := time.Month(config.Month).String()
		return fmt.Sprintf("%s %d: %d billable days 💸", monthName, config.Year, workingDays)
	} else {
		return fmt.Sprintf("💰 %d", workingDays)
	}
}
