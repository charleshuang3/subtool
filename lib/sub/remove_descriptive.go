package sub

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// RemoveDescriptiveSubtitles removes descriptive subtitles from the given file.
func RemoveDescriptiveSubtitles(inputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var cleanedSubtitles []string
	scanner := bufio.NewScanner(file)
	descriptiveRegex := regexp.MustCompile(`\((.*?)\)|\[(.*?)\]`)

	for scanner.Scan() {
		line := scanner.Text()
		// Remove descriptive parts
		cleanedLine := descriptiveRegex.ReplaceAllString(line, "")
		// Trim spaces and append if it's not completely empty
		if strings.TrimSpace(cleanedLine) != "" {
			cleanedSubtitles = append(cleanedSubtitles, cleanedLine)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	outputFile, err := os.Create("cleaned_" + inputFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, line := range cleanedSubtitles {
		fmt.Fprintln(outputFile, line)
	}

	return nil
}
