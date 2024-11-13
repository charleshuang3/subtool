package sub

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// DeeplAPIResponse represents the response structure from Deepl API
type DeeplAPIResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

// TranslateSubtitles translates the subtitles from the input file using Deepl.
func TranslateSubtitles(inputFile, deeplKeyFile string) error {
	subtitles, err := readSubtitles(inputFile)
	if err != nil {
		return err
	}

	deeplKey, err := ioutil.ReadFile(deeplKeyFile)
	if err != nil {
		return err
	}

	translatedSubtitles, err := translateTextUsingDeepl(subtitles, string(deeplKey))
	if err != nil {
		return err
	}

	return writeTranslatedSubtitles("translated_"+inputFile, translatedSubtitles)
}

// readSubtitles reads the subtitle lines from the input file.
func readSubtitles(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var subtitles []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			subtitles = append(subtitles, line)
		}
	}
	return subtitles, scanner.Err()
}

// translateTextUsingDeepl calls the Deepl API to translate the text.
func translateTextUsingDeepl(text []string, deeplKey string) ([]string, error) {
	client := &http.Client{}
	reqBody, err := json.Marshal(map[string]interface{}{
		"text":        text,
		"target_lang": "EN", // Set to your desired language code
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", ioutil.NopCloser(strings.NewReader(string(reqBody))))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+deeplKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to call Deepl API")
	}

	var deeplResp DeeplAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&deeplResp)
	if err != nil {
		return nil, err
	}

	var translatedTexts []string
	for _, translation := range deeplResp.Translations {
		translatedTexts = append(translatedTexts, translation.Text)
	}

	return translatedTexts, nil
}

// writeTranslatedSubtitles writes the translated subtitles to the output file.
func writeTranslatedSubtitles(outputFile string, subtitles []string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range subtitles {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}
	return nil
}
