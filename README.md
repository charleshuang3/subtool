# Subtool

Subtool is a command line utility for editing subtitle files (`.srt` only for now). It provides functionalities to clean up descriptive subtitles, analyze and remove unwant subtitles, and translate subtitles (via Deepl API only for now).

<details>

<summary>Story</summary>

I am a big fan of JAV, unfortuantely, I does not know much about Japanese, so I have to rely on subtitles to understand the story. However, I found recent ASR (speech to text) AI is good enough to help me. I have tested Google Cloud speech to text, and OpenAI's Whisper, GCP is better than Whisper, but it more expensive, ~$1 per hours and I can run Whisper locally with Macbook or use cheap whiisper provinder (~$0.1 per hour). My workflow is pretty easy:

Firstly, extract audio from video with ffmpeg:

```bash
ffmpeg -i <video> -vn -y <name>.mp3
```

Then use ASR to generate srt file. For Whisper, generate SRT is good enough, but for GCP, need to `enableWordTimeOffsets`, `enableAutomaticPunctuation` and use [WordInfo](https://cloud.google.com/speech-to-text/v2/docs/reference/rest/v2/projects.locations.recognizers/recognize#wordinfo) to generate srt file.

then use subtool to clean up the subtitle file, in order:

1. remove descriptive subtitles: usually useless
2. analyze repeat: some lines are repeated, it is usually useless
3. translate to languge I want

</details>

## Installation

First, ensure that Go is installed on your system.

### Install binary

```bash
go install github.com/charleshuang3/subtool
```

### From Source

Clone the repository and use `go build` to build the tool:

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
./subtool translate -i <input_file> [--deepl-key <deepl_key_file>]
```
