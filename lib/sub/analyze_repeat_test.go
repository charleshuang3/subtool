package sub

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeRepeatSubtitles(t *testing.T) {
	sb := &strings.Builder{}
	err := AnalyzeRepeatSubtitles("../../test_files/repeat.srt", sb)
	require.NoError(t, err)
	want := `This is repeated | 3
This is multiline and it's repeated too | 3
`
	assert.Equal(t, want, sb.String())
}
