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

// flattenedSubItem extrac all subtitle line items from subtitle.
func flattenedSubItem(sub *astisub.Subtitles) []string {
	var items []string
	for _, item := range sub.Items {
		for _, line := range item.Lines {
			for _, lineItem := range line.Items {
				items = append(items, lineItem.Text)
			}
		}
	}
	return items
}

func updateSubItem(sub *astisub.Subtitles, strs []string) {
	x := 0
	for i, item := range sub.Items {
		for j, line := range item.Lines {
			for k := range line.Items {
				sub.Items[i].Lines[j].Items[k].Text = strs[x]
				x++
			}
		}
	}
}
