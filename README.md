# 💸 BILLME

Your billable days calculator! Stop counting on your fingers - let me bill you properly!

A simple command-line tool to calculate working days in a month, with support for Czech public holidays and vacation days. Perfect for freelancers, contractors, and anyone who needs to track billable time.

## Features

- 📅 Calculate working days (Monday-Friday) for any month/year
- 🇨🇿 Automatic Czech public holiday detection and exclusion
- 🏖️ Vacation/time-off day subtraction
- 🎯 Multiple output formats (default, verbose, invoice-ready, celebratory)
- ⚡ Fast and lightweight
- 🛠️ Unix-style CLI with short and long flags

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/honzahovorka/billme.git
cd billme

# Build the binary
go build

# Optional: Install to your PATH
go install
```

### Requirements

- Go 1.24.5 or later

## Usage

### Basic Usage

```bash
# Current month
billme

# Specific month (current year)
billme 7

# Specific month and year
billme 7 2024
```

### With Options

```bash
# Verbose output
billme -v 7 2024
billme --verbose 7 2024

# Exclude Czech public holidays
billme -x 7 2024
billme --exclude-holidays 7 2024

# Subtract vacation days
billme -d 5 7 2024
billme --vacation-days 5 7 2024

# Combine options
billme -v -x -d 3 7 2024    # Verbose, exclude holidays, 3 vacation days
```

### Output Formats

```bash
# Default format
billme 7 2024
# Output: 💰 23

# Verbose format
billme -v 7 2024
# Output: July 2024: 23 billable days 💸

# Invoice-ready format (clean number for piping)
billme --invoice-ready 7 2024
# Output: 23

# Celebratory format
billme --ka-ching 7 2024
# Output: 23 days = CHA-CHING! 🤑
```

## CLI Options

| Short | Long | Description |
|-------|------|-------------|
| `-v` | `--verbose` | Verbose output with month name |
| `-h` | `--help` | Show help message |
| `-x` | `--exclude-holidays` | Exclude Czech public holidays |
| `-d <num>` | `--vacation-days <num>` | Number of vacation days to subtract |
| | `--ka-ching` | Celebratory output format |
| | `--invoice-ready` | Clean number output (for piping) |

## Czech Public Holidays

The tool automatically recognizes these Czech public holidays when using `--exclude-holidays`:

- **Nový rok** (January 1) - New Year's Day
- **Velikonoční pondělí** (varies) - Easter Monday
- **Svátek práce** (May 1) - Labour Day
- **Den vítězství** (May 8) - Liberation Day
- **Den slovanských věrozvěstů Cyrila a Metoděje** (July 5) - St. Cyril and Methodius Day
- **Den upálení mistra Jana Husa** (July 6) - Jan Hus Day
- **Den české státnosti** (September 28) - Czech Statehood Day
- **Den vzniku samostatného československého státu** (October 28) - Independence Day
- **Den boje za svobodu a demokracii** (November 17) - Freedom Day
- **Štědrý den** (December 24) - Christmas Eve
- **1. svátek vánoční** (December 25) - Christmas Day
- **2. svátek vánoční** (December 26) - St. Stephen's Day

## Examples

```bash
# Basic calculation
billme 7 2024
# Output: 💰 23

# July 2024 with Czech holidays excluded (Jan Hus Day on July 6)
billme -x 7 2024
# Output: 💰 22

# July 2024 with holidays and 5 vacation days
billme -x -d 5 7 2024
# Output: 💰 17

# Verbose output for December with holidays (Christmas period)
billme -v -x 12 2024
# Output: December 2024: 19 billable days 💸

# Invoice-ready format for scripting
DAYS=$(billme --invoice-ready -x -d 2 7 2024)
echo "Billable days: $DAYS"
# Output: Billable days: 20
```

## Development

### Project Structure

```
billme/
├── main.go               # Main application entry point
├── internal/             # Private application code
│   ├── calculator/       # Business logic for day calculations
│   │   ├── calculator.go
│   │   └── calculator_test.go
│   ├── cli/              # Command-line interface handling
│   │   ├── cli.go
│   │   └── cli_test.go
│   └── holidays/         # Czech holiday definitions and logic
│       ├── holidays.go
│       └── holidays_test.go
├── go.mod
└── README.md
```

### Building

```bash
# Build for current platform
go build

# Build for different platforms
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=amd64 go build
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for specific package
go test ./internal/calculator
go test ./internal/holidays
go test ./internal/cli
```

### Code Organization

- **`main.go`** - Main application entry point and orchestration
- **`internal/calculator/`** - Core business logic for calculating working days
- **`internal/cli/`** - Command-line argument parsing and output formatting
- **`internal/holidays/`** - Czech holiday definitions and Easter calculation

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright (c) 2025 Honza Hovorka

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Setup

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/billme.git`
3. Create a feature branch: `git checkout -b feature-name`
4. Make your changes and add tests
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -am 'Add some feature'`
7. Push to the branch: `git push origin feature-name`
8. Submit a pull request

## Built With

This project was built with the assistance of [opencode](https://github.com/sst/opencode) - an AI-powered coding assistant that helped architect, implement, and document this tool.
