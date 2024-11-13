package sub

import (
	"io"
	"regexp"

	"github.com/asticode/go-astisub"
)

var (
	// descriptiveRegex matches text within parentheses or square brackets
	descriptiveRegex = regexp.MustCompile(`\((.*?)\)|\[(.*?)\]`)
)

// RemoveDescriptiveSubtitles removes descriptive subtitles from the given file.
func RemoveDescriptiveSubtitles(inputFile string, output io.Writer) error {
	sub, err := astisub.OpenFile(inputFile)
	if err != nil {
		return err
	}

	cleanedSubtitles := []*astisub.Item{}

	for _, item := range sub.Items {
		subtitleText := multiLinesToOne(item.Lines)

		if !descriptiveRegex.MatchString(subtitleText) {
			cleanedSubtitles = append(cleanedSubtitles, item)
		}
	}

	sub.Items = cleanedSubtitles
	return sub.WriteToSRT(output)
}
