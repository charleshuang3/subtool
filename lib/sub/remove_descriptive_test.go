package sub

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveDescriptiveSubtitles(t *testing.T) {
	sb := &strings.Builder{}
	err := RemoveDescriptiveSubtitles("../../test_files/descriptive.srt", sb)
	require.NoError(t, err)
	want := "\ufeff" + `1
00:00:25,000 --> 00:00:54,000
This is not descriptive
`
	assert.Equal(t, want, sb.String())
}
