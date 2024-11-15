package sub

import (
	"strings"

	"github.com/asticode/go-astisub"
)

func joinLinesWithSpace(lines []astisub.Line) string {
	return joinLines(lines, " ", " ")
}

// joinLines joins multiline subtitle in 1 timeframe into 1 line with given splitter.
func joinLines(lines []astisub.Line, lineSpiltter, itemSpiltter string) string {
	isFirstLine := true
	sb := strings.Builder{}
	for _, line := range lines {
		if isFirstLine {
			isFirstLine = false
		} else {
			sb.WriteString(lineSpiltter)
		}
		isFirstItem := true
		for _, item := range line.Items {
			if isFirstItem {
				isFirstItem = false
			} else {
				sb.WriteString(itemSpiltter)
			}
			sb.WriteString(item.Text)
		}
	}
	return sb.String()
}
