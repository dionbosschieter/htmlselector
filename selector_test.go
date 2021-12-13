package htmlselector

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPoundSelector(t *testing.T) {
	selector, err := NewSelector("div#test")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, "div#test", selector.RawQuery)

	require.Equal(t, 1, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div#test", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"id": "test"}, row.Attributes)
}

func TestNewDoublePoundSelector(t *testing.T) {
	selector, err := NewSelector("div#test  div#test123")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, "div#test  div#test123", selector.RawQuery)

	require.Equal(t, 2, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div#test", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"id": "test"}, row.Attributes)

	row = selector.Rows[1]
	assert.Equal(t, false, row.IsFirst)
	assert.Equal(t, "div#test123", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"id": "test123"}, row.Attributes)
}

func TestNewFirstChildSelector(t *testing.T) {
	selector, err := NewSelector("div > div > p")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, "div > div > p", selector.RawQuery)

	require.Equal(t, 3, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div", row.RawQuery)
	assert.Equal(t, "div", row.Type)

	row = selector.Rows[1]
	assert.Equal(t, false, row.IsFirst)
	assert.Equal(t, 1, row.NthChild)
	assert.Equal(t, "div", row.RawQuery)
	assert.Equal(t, "div", row.Type)

	row = selector.Rows[2]
	assert.Equal(t, false, row.IsFirst)
	assert.Equal(t, 1, row.NthChild)
	assert.Equal(t, "p", row.RawQuery)
	assert.Equal(t, "p", row.Type)
}

func TestNewCombinedFirstChildSelector(t *testing.T) {
	selector, err := NewSelector("div#test > div.test > p a")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, "div#test > div.test > p a", selector.RawQuery)

	require.Equal(t, 4, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div#test", row.RawQuery)
	assert.Equal(t, "div", row.Type)

	row = selector.Rows[1]
	assert.Equal(t, false, row.IsFirst)
	assert.Equal(t, 1, row.NthChild)
	assert.Equal(t, "div.test", row.RawQuery)
	assert.Equal(t, "div", row.Type)

	row = selector.Rows[2]
	assert.Equal(t, false, row.IsFirst)
	assert.Equal(t, 1, row.NthChild)
	assert.Equal(t, "p", row.RawQuery)
	assert.Equal(t, "p", row.Type)

	row = selector.Rows[3]
	assert.Equal(t, false, row.IsFirst)
	assert.Equal(t, 0, row.NthChild)
	assert.Equal(t, "a", row.RawQuery)
	assert.Equal(t, "a", row.Type)
}

func TestNewClassSelector(t *testing.T) {
	selector, err := NewSelector("div.test")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, "div.test", selector.RawQuery)

	require.Equal(t, 1, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div.test", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"class": "test"}, row.Attributes)
}

func TestNewMultiClassSelector(t *testing.T) {
	selector, err := NewSelector("div.test.testA.testB")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, "div.test.testA.testB", selector.RawQuery)

	require.Equal(t, 1, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div.test.testA.testB", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"class": "test testA testB"}, row.Attributes)
}

func TestNewCombinedSelector(t *testing.T) {
	selector, err := NewSelector("div#test.test")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, selector.RawQuery, "div#test.test")

	require.Equal(t, 1, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div#test.test", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"id": "test", "class": "test"}, row.Attributes)
}

func TestNewCombinedSelectorAttributes(t *testing.T) {
	selector, err := NewSelector("div#test.test[abc=\"123\"]")

	require.NoError(t, err, "No error was thrown on a non error query")
	assert.NotNil(t, selector, "More nodes returned than expected")
	assert.Equal(t, selector.RawQuery, "div#test.test[abc=\"123\"]")

	require.Equal(t, 1, len(selector.Rows))
	row := selector.Rows[0]
	assert.Equal(t, true, row.IsFirst)
	assert.Equal(t, "div#test.test[abc=\"123\"]", row.RawQuery)
	assert.Equal(t, "div", row.Type)
	assert.Equal(t, map[string]string{"id": "test", "class": "test", "abc": "123"}, row.Attributes)
}
