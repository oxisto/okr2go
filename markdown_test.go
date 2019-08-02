package okr2go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMarkdown(t *testing.T) {
	objectives, err := ParseMarkdown("example.md")

	assert.Nil(t, err)

	t.Logf("%+v", objectives)
}
