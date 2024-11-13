package sub

import (
	"fmt"
	"strings"

	"github.com/asticode/go-astisub"
)

const (
	RepeatThreshold = 3 // Minimum number of occurrences to consider a subtitle as repeated
)

// AnalyzeRepeatSubtitles analyzes the input file for repeated subtitles.
func AnalyzeRepeatSubtitles(inputFile string) error {
	sub, err := astisub.OpenFile(inputFile)
	if err != nil {
		return err
	}

	// Create a map to count occurrences of each subtitle line
	subtitleCount := make(map[string]int)

	for _, item := range sub.Items {
		// Combine the lines into a single string
		subtitleText := multiLinesToOne(item.Lines)

		// Increment the count for this subtitle line
		subtitleCount[subtitleText]++
	}

	// Print the repeated subtitles and their counts
	fmt.Println("Repeated subtitles:")
	for subtitle, count := range subtitleCount {
		if count >= RepeatThreshold {
			fmt.Printf("%s | %d\n", subtitle, count)
		}
	}

	return nil
}

func multiLinesToOne(lines []astisub.Line) string {
	sb := strings.Builder{}
	for _, line := range lines {
		for _, item := range line.Items {
			sb.WriteString(item.Text)
			sb.WriteString(" ")
		}
	}
	return strings.TrimSpace(sb.String())
}
