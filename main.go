package main

import (
	"billme/internal/calculator"
	"billme/internal/cli"
	"fmt"
	"os"
)

func main() {
	config, err := cli.ParseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		cli.ShowUsage()
		os.Exit(1)
	}

	if config.Help {
		cli.ShowHelp()
		return
	}

	workingDays := calculator.CountWorkingDaysWithHolidaysAndVacation(config.Month, config.Year, "CZ", config.ExcludeHolidays, config.VacationDays)
	output := cli.FormatOutput(workingDays, config)
	fmt.Println(output)
}
