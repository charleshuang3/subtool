package sub

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTranslateWithKeyFile(t *testing.T) {
	_, err := os.ReadFile("../../test_files/deepl.key")
	if err != nil {
		t.Skipf("no key file: %v", err)
	}
	sb := &strings.Builder{}
	err = TranslateSubtitles("../../test_files/translate.srt", "EN", "../../test_files/deepl.key", "free", sb)
	require.NoError(t, err)

	want := "\ufeff" + `1
00:00:00,000 --> 00:00:07,000
Hello. - Hello.

2
00:00:07,000 --> 00:00:16,000
I love Tokyo.
This is my first time here.
`
	assert.Equal(t, want, sb.String())
}

func TestTranslateWithEnv(t *testing.T) {
	if os.Getenv(deepLAuthnKeyKey) == "" {
		t.Skip("no env")
	}

	sb := &strings.Builder{}
	err := TranslateSubtitles("../../test_files/translate.srt", "EN", "", "free", sb)
	require.NoError(t, err)

	want := "\ufeff" + `1
00:00:00,000 --> 00:00:07,000
Hello. - Hello.

2
00:00:07,000 --> 00:00:16,000
I love Tokyo.
This is my first time here.
`
	assert.Equal(t, want, sb.String())
}
