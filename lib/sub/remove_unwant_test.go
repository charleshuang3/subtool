package sub

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveUnwantSubtitles(t *testing.T) {
	sb := &strings.Builder{}
	err := RemoveUnwantSubtitles("../../test_files/remove_unwant.srt", "../../test_files/unwant_list.txt", sb)
	require.NoError(t, err)
	want := "\ufeff" + `1
00:00:54,000 --> 00:01:05,000
This is wanted

2
00:01:05,000 --> 00:01:15,000
This is multiline
and it's wanted too
`
	assert.Equal(t, want, sb.String())

}
