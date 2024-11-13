package sub

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AnalyzeRepeatSubtitles analyzes the input file for repeated subtitles.
func AnalyzeRepeatSubtitles(inputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a map to count occurrences of each subtitle line
	subtitleCount := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			subtitleCount[trimmedLine]++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("Subtitle repetition analysis:")
	for subtitle, count := range subtitleCount {
		if count > 1 {
			fmt.Printf("Subtitle: '%s' repeated %d times\n", subtitle, count)
		}
	}

	return nil
}
