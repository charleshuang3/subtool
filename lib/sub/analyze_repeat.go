package sub

import (
	"fmt"
	"io"

	"github.com/asticode/go-astisub"
)

const (
	RepeatThreshold = 3 // Minimum number of occurrences to consider a subtitle as repeated
)

// AnalyzeRepeatSubtitles analyzes the input file for repeated subtitles.
func AnalyzeRepeatSubtitles(inputFile string, output io.Writer) error {
	sub, err := astisub.OpenFile(inputFile)
	if err != nil {
		return err
	}

	// Create a map to count occurrences of each subtitle line
	subtitleCount := make(map[string]int)

	for _, item := range sub.Items {
		// Combine the lines into a single string
		subtitleText := joinLinesWithSpace(item.Lines)

		// Increment the count for this subtitle line
		subtitleCount[subtitleText]++
	}

	// Print the repeated subtitles and their counts
	for subtitle, count := range subtitleCount {
		if count >= RepeatThreshold {
			fmt.Fprintf(output, "%s | %d\n", subtitle, count)
		}
	}

	return nil
}
