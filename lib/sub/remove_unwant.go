package sub

import (
	"bufio"
	"io"
	"os"

	"github.com/asticode/go-astisub"
)

// RemoveUnwantSubtitles removes unwant subtitles using a provided repeat file.
func RemoveUnwantSubtitles(inputFile, unwantFile string, output io.Writer) error {
	sub, err := astisub.OpenFile(inputFile)
	if err != nil {
		return err
	}

	unwants, err := os.Open(unwantFile)
	if err != nil {
		return err
	}
	defer unwants.Close()

	unwantSet := make(map[string]bool)
	scanner := bufio.NewScanner(unwants)
	for scanner.Scan() {
		unwantSet[scanner.Text()] = true
	}

	cleanedSubtitles := []*astisub.Item{}
	for _, item := range sub.Items {
		subtitleText := joinLinesWithSpace(item.Lines)

		if _, ok := unwantSet[subtitleText]; !ok {
			cleanedSubtitles = append(cleanedSubtitles, item)
		}
	}

	sub.Items = cleanedSubtitles
	return sub.WriteToSRT(output)
}
