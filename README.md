# Subtool

Subtool is a command line utility for editing subtitle files (`.srt` only for now). It provides functionalities to clean up descriptive subtitles, analyze and remove unwant subtitles, and translate subtitles (via Deepl API only for now).

## Installation

Ensure that Go is installed on your system. Clone the repository and use `go build` to build the tool:

```bash
git clone https://github.com/charleshuang3/subtool
cd subtool
go build
```

## Usage

Subtool provides several commands with different functionalities:

### Remove Descriptive Subtitles

Removes descriptive subtitles (e.g., sound effects):

```bash
./subtool remove-descriptive-subtitle -i <input_file>
```

### Analyze Repeat

Analyzes subtitle files to identify repeated lines:

```bash
./subtool analyze-repeat -i <input_file>
```

### Remove Unwanted Subtitles

Removes subtitles considered unwant based on a file contains unwanted sub, you can use analyze-repeat's output, 1 per line:

```bash
./subtool remove-unwant -i <input_file> --unwant <unwant_file>
```

### Translate Subtitles via Deepl

Translates the subtitles using the Deepl API:

```bash
./subtool translate -i <input_file> --deepl-key <deepl_key_file>
```
