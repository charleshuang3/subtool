package sub

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/asticode/go-astisub"
	"github.com/michimani/deepl-sdk-go"
	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
)

// TranslateSubtitles translates the subtitles from the input file using Deepl.
// TODO: pass in target language https://github.com/michimani/deepl-sdk-go/blob/bdd76af53e59bafa3439b9e0934539543678a1fb/types/lang.go#L105
func TranslateSubtitles(inputFile, deeplKeyFile string, out io.Writer) error {
	sub, err := astisub.OpenFile(inputFile)
	if err != nil {
		return err
	}

	b, err := os.ReadFile(deeplKeyFile)
	if err != nil {
		return err
	}
	deeplKey := string(b)

	// TODO: arg key file is not a must if env var is set
	os.Setenv("DEEPL_API_AUTHN_KEY", deeplKey)
	// TODO: add optional arg DEEPL_API_PLAN
	os.Setenv("DEEPL_API_PLAN", "free")

	client, err := deepl.NewClient()
	if err != nil {
		return err
	}

	params := &params.TranslateTextParams{
		TargetLang: types.TargetLangEN,
		Text:       []string{},
	}

	for _, item := range sub.Items {
		subtitleText := joinLines(item.Lines, "\n", "\t")
		params.Text = append(params.Text, subtitleText)
	}

	res, errRes, err := client.TranslateText(context.Background(), params)

	if err != nil {
		return err
	}

	if errRes != nil {
		return fmt.Errorf("Deepl Err: %s %s", errRes.StatusCode.Description(), errRes.Message)
	}

	for i, translated := range res.Translations {
		lines := strings.Split(translated.Text, "\n")
		for j, line := range lines {
			items := strings.Split(line, "\t")
			for k, item := range items {
				sub.Items[i].Lines[j].Items[k].Text = item
			}
		}
	}

	return sub.WriteToSRT(out)
}
