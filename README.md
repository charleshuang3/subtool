# Subtool

Subtool is a command line utility for editing subtitle files (`.srt` only for now). It provides functionalities to clean up descriptive subtitles, analyze and remove useless subtitles, and translate subtitles (via Deepl API only for now).

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

### Remove Useless Subtitles

Removes subtitles considered useless based on a repeat file:

```bash
./subtool remove-useless-subtitle -i <input_file> --rm <repeat_file>
```

### Translate Subtitles via Deepl

Translates the subtitles using the Deepl API:

```bash
./subtool translate -i <input_file> --deepl-key <deepl_key_file>
```
