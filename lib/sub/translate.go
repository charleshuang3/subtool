package sub

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/asticode/go-astisub"
	"github.com/michimani/deepl-sdk-go"
	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
)

const (
	deepLAuthnKeyKey = "DEEPL_API_AUTHN_KEY"
	deepLPlanKey     = "DEEPL_API_PLAN"
)

// TranslateSubtitles translates the subtitles from the input file using Deepl.
// TODO: pass in target language https://github.com/michimani/deepl-sdk-go/blob/bdd76af53e59bafa3439b9e0934539543678a1fb/types/lang.go#L105
func TranslateSubtitles(inputFile, targetLanguage, deeplKeyFile, deeplPlan string, out io.Writer) error {
	sub, err := astisub.OpenFile(inputFile)
	if err != nil {
		return err
	}

	if os.Getenv(deepLAuthnKeyKey) == "" && deeplKeyFile == "" {
		return fmt.Errorf("Deepl API key not found. Please set the %s environment variable or provide the key file path.", deepLAuthnKeyKey)
	}

	if os.Getenv(deepLAuthnKeyKey) == "" {
		b, err := os.ReadFile(deeplKeyFile)
		if err != nil {
			return err
		}
		deeplKey := string(b)

		os.Setenv(deepLAuthnKeyKey, deeplKey)
	}

	if os.Getenv(deepLPlanKey) == "" {
		os.Setenv(deepLPlanKey, deeplPlan)
	}

	client, err := deepl.NewClient()
	if err != nil {
		return err
	}

	params := &params.TranslateTextParams{
		TargetLang: types.TargetLangCode(targetLanguage),
		Text:       []string{},
	}

	params.Text = flattenedSubItem(sub)

	res, errRes, err := client.TranslateText(context.Background(), params)

	if err != nil {
		return err
	}

	if errRes != nil {
		return fmt.Errorf("Deepl Err: %s %s", errRes.StatusCode.Description(), errRes.Message)
	}

	translated := []string{}
	for _, it := range res.Translations {
		translated = append(translated, it.Text)
	}
	updateSubItem(sub, translated)

	return sub.WriteToSRT(out)
}
