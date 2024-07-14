# Quake Log Parser

## Project Description

Quake Log Parser is a Go project designed to parse Quake game log files, extract relevant match data, and generate reports on player kills and deaths by different means. The project reads a log file, processes the data, and outputs structured reports in JSON format.

## Project Structure

```
go.mod
main.go
cmd/parser/quake_log_reporter.go
internal/domain/death_cause.go
internal/domain/match.go
internal/domain/player.go
internal/parser/log_parser.go
internal/parser/log_parser_test.go
internal/report/report.go
internal/service/match_service.go
internal/service/match_service_test.go
internal/utils/file_reader.go
internal/utils/file_reader_test.go
```

### Files Description

- `go.mod`: Go module definition file.
- `main.go`: Main entry point of the application.
- `cmd/parser/quake_log_reporter.go`: Handles the parsing and report generation process.
- `internal/domain/death_cause.go`: Defines death causes in the game.
- `internal/domain/match.go`: Defines match structure and serialization methods.
- `internal/domain/player.go`: Defines player structure.
- `internal/parser/log_parser.go`: Parses the log file to extract match data.
- `internal/parser/log_parser_test.go`: Unit tests for the log parser.
- `internal/report/report.go`: Generates JSON reports from parsed match data.
- `internal/service/match_service.go`: Provides services to group match data.
- `internal/service/match_service_test.go`: Unit tests for match services.
- `internal/utils/file_reader.go`: Utility to read log files.
- `internal/utils/file_reader_test.go`: Unit tests for file reading utility.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/quakelogparser.git
    cd quakelogparser
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```

## Running the Project

1. Ensure you have a log file named `qgames.log` in the `files` directory.
2. Run the project:
    ```sh
    go run main.go
    ```

## Running Tests

Execute the following command to run all tests:
```sh
go test ./...
```

## Example Usage

1. Place your Quake log file (must be named qgames.log) in the `files` directory.
2. Execute the project using the `go run main.go` command.
3. Check the `output` directory for the generated JSON reports:
    - `matches.json`: Contains detailed match data.
    - `matches_dbm.json`: Contains data on deaths by means.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Contact

For any questions or suggestions, please open an issue or contact the repository owner at your.email@example.com.
