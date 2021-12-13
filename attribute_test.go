package htmlselector

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttributeParsert(t *testing.T) {
	attributes := parseAttributes("abc=test,abc,blabla=123")

	require.Len(t, attributes, 3)

	assert.Equal(t, "abc", attributes[0].key)
	assert.Equal(t, "test", attributes[0].value)
	assert.Equal(t, "abc", attributes[1].key)
	assert.Equal(t, "blabla", attributes[2].key)
	assert.Equal(t, "123", attributes[2].value)
}