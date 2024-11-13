package sub

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RemoveUselessSubtitles removes useless subtitles using a provided repeat file.
func RemoveUselessSubtitles(inputFile, repeatFile string) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	repeats, err := os.Open(repeatFile)
	if err != nil {
		return err
	}
	defer repeats.Close()

	// Load repeats into a set
	repeatSet := make(map[string]struct{})
	scanner := bufio.NewScanner(repeats)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			repeatSet[line] = struct{}{}
		}
	}

	// Open output file to write cleaned subtitles
	outputFile, err := os.Create("cleaned_" + inputFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Scan through input file, writing non-useless subtitles to output
	scanner = bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if _, found := repeatSet[line]; !found && line != "" {
			fmt.Fprintln(outputFile, line)
		}
	}

	return scanner.Err()
}
