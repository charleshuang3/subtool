package sub

import (
	"testing"

	"github.com/asticode/go-astisub"
	"github.com/stretchr/testify/require"
)

func TestJoinLines(t *testing.T) {
	sub, err := astisub.OpenFile("../../test_files/remove_unwant.srt")
	require.NoError(t, err)

	tests := []struct {
		name  string
		input []astisub.Line
		want  string
	}{
		{
			name:  "single_line",
			input: sub.Items[0].Lines,
			want:  "This is repeated unwant",
		},
		{
			name:  "multi_line",
			input: sub.Items[2].Lines,
			want:  "This is multiline***and it's repeated unwant too",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := joinLines(tt.input, "***", "+++")
			require.Equal(t, tt.want, got)
		})
	}
}
